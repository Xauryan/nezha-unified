//go:build !certimate_full

package embedded

import "context"

func (e *Engine) NativeCertificateDownload(context.Context, string, string, NativeCertificateDownloadOptions) ([]byte, error) {
	return nil, ErrRuntimeUnavailable
}

func (e *Engine) NativeRevokeCertificate(context.Context, string, string) error {
	return ErrRuntimeUnavailable
}

func (e *Engine) NativeTestNotification(context.Context, CredentialUse, string) error {
	return ErrRuntimeUnavailable
}
