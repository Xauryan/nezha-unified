// Package agentproto defines the optional, versioned unified Agent extension.
// Existing Task/TaskResult protobuf fields and task numbers are unchanged.
package agentproto

import "encoding/json"

const (
	Protocol                 = 1
	TaskCapabilities  uint64 = 0x4e5a0100
	TaskCall          uint64 = 0x4e5a0101
	TaskHostAddresses uint64 = 0x4e5a0102
	MaxPayload               = 256 << 10
)

type HostAddress struct {
	Interface  string `json:"interface"`
	Address    string `json:"address"`
	Temporary  bool   `json:"temporary"`
	Deprecated bool   `json:"deprecated"`
}
type HostAddresses struct {
	Protocol  int           `json:"protocol"`
	Addresses []HostAddress `json:"addresses"`
}

type Capabilities struct {
	Protocol          int  `json:"protocol"`
	CertificateDeploy bool `json:"certificate_deploy"`
	Online            bool `json:"online"`
}
type Request struct {
	Protocol    int             `json:"protocol"`
	OperationID string          `json:"operation_id"`
	Method      string          `json:"method"`
	Payload     json.RawMessage `json:"payload,omitempty"`
}
type Response struct {
	Protocol    int             `json:"protocol"`
	OperationID string          `json:"operation_id"`
	Phase       string          `json:"phase"`
	Final       bool            `json:"final"`
	Code        string          `json:"code,omitempty"`
	Payload     json.RawMessage `json:"payload,omitempty"`
}
type Target struct {
	ID            string   `json:"id"`
	Version       uint64   `json:"version"`
	BundlePath    string   `json:"bundle_path"`
	AllowedNames  []string `json:"allowed_names"`
	CheckCommand  []string `json:"check_command"`
	ReloadCommand []string `json:"reload_command"`
	HealthCommand []string `json:"health_command"`
	TLSAddress    string   `json:"tls_address,omitempty"`
	TLSServerName string   `json:"tls_server_name,omitempty"`
	FileMode      uint32   `json:"file_mode,omitempty"`
	OwnerUID      *int     `json:"owner_uid,omitempty"`
	OwnerGID      *int     `json:"owner_gid,omitempty"`
}
type Deployment struct {
	OperationID    string `json:"operation_id"`
	Sequence       uint64 `json:"sequence"`
	Target         Target `json:"target"`
	CertificatePEM string `json:"certificate_pem,omitempty"`
	PrivateKeyPEM  string `json:"private_key_pem,omitempty"`
	Fingerprint    string `json:"fingerprint"`
}
type Receipt struct {
	OperationID         string `json:"operation_id"`
	TargetID            string `json:"target_id"`
	TargetVersion       uint64 `json:"target_version"`
	Sequence            uint64 `json:"sequence"`
	Fingerprint         string `json:"fingerprint"`
	Phase               string `json:"phase"`
	Code                string `json:"code,omitempty"`
	Verified            bool   `json:"verified"`
	CertificateObserved bool   `json:"certificate_observed"`
	Inspected           bool   `json:"inspected"`
	CurrentOperationID  string `json:"current_operation_id,omitempty"`
	CurrentFingerprint  string `json:"current_fingerprint,omitempty"`
	CurrentPairValid    bool   `json:"current_pair_valid"`
	RolledBack          bool   `json:"rolled_back"`
}
