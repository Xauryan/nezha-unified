// Package policy contains shared authorization primitives for DNS and certificates.
package policy

import (
	"errors"
	"strings"

	"golang.org/x/net/idna"
)

var ErrName = errors.New("invalid DNS name")

// NormalizeName accepts DNS service labels, IDNA names and one leftmost wildcard.
// It deliberately rejects URLs, empty labels and ambiguous multi-root-dot inputs.
func NormalizeName(name string) (string, error) {
	if name == "" || name != strings.TrimSpace(name) {
		return "", ErrName
	}
	name = strings.TrimSuffix(name, ".")
	labels := strings.Split(name, ".")
	for i, label := range labels {
		if label == "*" && i == 0 && len(labels) > 1 {
			continue
		}
		if label == "" {
			return "", ErrName
		}
		prefix := ""
		// TXT verification and SRV records use underscore-prefixed labels.
		if strings.HasPrefix(label, "_") {
			prefix, label = "_", label[1:]
		}
		ascii, err := idna.Lookup.ToASCII(label)
		if err != nil || ascii == "" {
			return "", ErrName
		}
		ascii = strings.ToLower(ascii)
		if ascii[0] == '-' || ascii[len(ascii)-1] == '-' {
			return "", ErrName
		}
		for _, ch := range ascii {
			if !(ch >= 'a' && ch <= 'z' || ch >= '0' && ch <= '9' || ch == '-') {
				return "", ErrName
			}
		}
		labels[i] = prefix + ascii
		if len(labels[i]) > 63 {
			return "", ErrName
		}
	}
	result := strings.Join(labels, ".")
	if len(result) > 253 {
		return "", ErrName
	}
	return result, nil
}

type NameScope struct {
	Root          string `json:"root"`
	IncludeRoot   bool   `json:"include_root"`
	Descendants   bool   `json:"descendants"`
	AllowWildcard bool   `json:"allow_wildcard"`
}

func (s NameScope) Allows(name string) bool {
	root, err := NormalizeName(s.Root)
	if err != nil || strings.Contains(root, "*") {
		return false
	}
	n, err := NormalizeName(name)
	if err != nil {
		return false
	}
	if strings.HasPrefix(n, "*.") {
		if !s.AllowWildcard {
			return false
		}
		// A wildcard at the subtree root describes descendants, not the root.
		n = strings.TrimPrefix(n, "*.")
		return s.Descendants && (n == root || strings.HasSuffix(n, "."+root))
	}
	return n == root && s.IncludeRoot || s.Descendants && strings.HasSuffix(n, "."+root)
}

// AllowsAllCertificateNames is separate from DNS authorization. The caller must
// supply actual parsed SANs (or the final resolved ACME order names).
func AllowsAllCertificateNames(scopes []NameScope, names []string) bool {
	if len(names) == 0 {
		return false
	}
	for _, name := range names {
		if strings.Contains(name, "_") {
			return false
		}
		allowed := false
		for _, scope := range scopes {
			if scope.Allows(name) {
				allowed = true
				break
			}
		}
		if !allowed {
			return false
		}
	}
	return true
}

type DNSRecord struct {
	AccountID string `json:"account_id"`
	ZoneID    string `json:"zone_id"`
	ID        string `json:"id"`
	Name      string `json:"name"`
	Type      string `json:"type"`
}

type DNSGrant struct {
	AccountID string `json:"account_id"`
	ZoneID    string `json:"zone_id"`
	NameScope
	Types []string `json:"types"`
}

func (g DNSGrant) Allows(record DNSRecord) bool {
	if g.AccountID == "" || g.ZoneID == "" || g.AccountID != record.AccountID || g.ZoneID != record.ZoneID || !g.NameScope.Allows(record.Name) {
		return false
	}
	for _, typ := range g.Types {
		if typ == record.Type {
			return true
		}
	}
	return false
}

// AuthorizeUpdate requires a provider-fetched original record, not client
// supplied metadata. Both sides of a rename and the stable record identity matter.
func (g DNSGrant) AuthorizeUpdate(before, after DNSRecord) bool {
	return before.ID != "" && before.ID == after.ID && g.Allows(before) && g.Allows(after)
}
