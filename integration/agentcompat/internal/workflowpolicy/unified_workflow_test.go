package workflowpolicy_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/nezhahq/nezha/integration/agentcompat/internal/workflowpolicy"
	"github.com/stretchr/testify/require"
)

// Original quality-policy tests use preserved upstream workflow fixtures. This
// separately checks the active integration workflow using the same security policy.
func TestPolicy_UnifiedWorkflow(t *testing.T) {
	require.NoError(t, workflowpolicy.VerifyFile(filepath.Join(repositoryRoot(t), ".github", "workflows", "ci.yml"), workflowpolicy.RepositoryNezha))
}

func TestPolicy_NodeAutomaticCacheMustBeDisabled(t *testing.T) {
	data, err := os.ReadFile(filepath.Join(repositoryRoot(t), ".github", "workflows", "ci.yml"))
	require.NoError(t, err)
	unsafe := strings.Replace(string(data), "package-manager-cache: false", "package-manager-cache: true", 1)
	err = workflowpolicy.Verify([]byte(unsafe), workflowpolicy.RepositoryNezha)
	requireTypedPolicyError(t, err, workflowpolicy.RuleReusableExecutable)
}
