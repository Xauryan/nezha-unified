package policy

import "testing"

func TestDNSSubtreeBoundary(t *testing.T) {
	s := NameScope{Root: "MiRaCle.Example.COM.", Descendants: true}
	for _, name := range []string{"www.miracle.example.com", "a.b.miracle.example.com", "_acme-challenge.miracle.example.com.", "WWW.MIRACLE.EXAMPLE.COM"} {
		if !s.Allows(name) {
			t.Errorf("denied %q", name)
		}
	}
	for _, name := range []string{"miracle.example.com", "evilmiracle.example.com", "other.example.com", "example.com", "*.miracle.example.com", "a..miracle.example.com", "x.miracle.example.com..", "https://x.miracle.example.com", " x.miracle.example.com"} {
		if s.Allows(name) {
			t.Errorf("allowed %q", name)
		}
	}
	s.IncludeRoot = true
	if !s.Allows("miracle.example.com") {
		t.Fatal("root toggle ignored")
	}
	s.AllowWildcard = true
	if !s.Allows("*.miracle.example.com") {
		t.Fatal("wildcard toggle ignored")
	}
	if s.Allows("*.*.miracle.example.com") {
		t.Fatal("nested wildcard")
	}
}

func TestIDNAAndCertificateSANs(t *testing.T) {
	s := NameScope{Root: "bücher.example", IncludeRoot: true, Descendants: true}
	if !s.Allows("xn--bcher-kva.example.") {
		t.Fatal("IDNA mismatch")
	}
	if !AllowsAllCertificateNames([]NameScope{s}, []string{"bücher.example", "www.bücher.example"}) {
		t.Fatal("valid SANs rejected")
	}
	for _, names := range [][]string{nil, {"bücher.example", "other.example"}, {"*.bücher.example"}, {"_acme-challenge.bücher.example"}} {
		if AllowsAllCertificateNames([]NameScope{s}, names) {
			t.Fatalf("unexpected SAN authorization: %v", names)
		}
	}
}

func TestDNSRenameAndProviderIdentity(t *testing.T) {
	g := DNSGrant{AccountID: "account-a", ZoneID: "zone-a", NameScope: NameScope{Root: "miracle.example.com", Descendants: true}, Types: []string{"A", "AAAA"}}
	r := DNSRecord{AccountID: "account-a", ZoneID: "zone-a", ID: "record-1", Name: "www.miracle.example.com", Type: "A"}
	if !g.AuthorizeUpdate(r, r) {
		t.Fatal("valid update rejected")
	}
	for _, mutate := range []func(*DNSRecord){func(r *DNSRecord) { r.Name = "www.other.example.com" }, func(r *DNSRecord) { r.ZoneID = "other" }, func(r *DNSRecord) { r.AccountID = "other" }, func(r *DNSRecord) { r.ID = "other" }, func(r *DNSRecord) { r.Type = "TXT" }} {
		changed := r
		mutate(&changed)
		if g.AuthorizeUpdate(r, changed) || g.AuthorizeUpdate(changed, r) {
			t.Fatal("update bypassed original or destination authorization")
		}
	}
}
