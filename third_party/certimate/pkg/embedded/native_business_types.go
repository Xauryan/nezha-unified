package embedded

type NativeCertificateDownloadOptions struct {
	FileFormat   string `json:"fileFormat"`
	PfxPassword  string `json:"pfxPassword,omitempty"`
	PfxEncoder   string `json:"pfxEncoder,omitempty"`
	JksAlias     string `json:"jksAlias,omitempty"`
	JksKeypass   string `json:"jksKeypass,omitempty"`
	JksStorepass string `json:"jksStorepass,omitempty"`
}
