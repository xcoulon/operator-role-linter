package apiresources_test

import (
	"testing"

	"github.com/codeready-toolchain/operator-role-linter/pkg/apiresources"
	_ "github.com/codeready-toolchain/operator-role-linter/pkg/log"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestVerify(t *testing.T) {
	// given
	ruleFile := "../../test/role.yaml"
	crdsDir := "../../test/crds"
	// when
	errors, err := apiresources.Verify(ruleFile, crdsDir, "openshift-4.2")
	// then
	require.NoError(t, err)
	require.Len(t, errors, 5)
	assert.Contains(t, errors, apiresources.InvalidGroupResourceScopeError{
		Group:      "",
		Resource:   "namespaces",
		Namespaced: false,
	})
	assert.Contains(t, errors, apiresources.InvalidGroupResourceScopeError{
		Group:      "",
		Resource:   "componentstatuses",
		Namespaced: false,
	})
	assert.Contains(t, errors, apiresources.InvalidGroupResourceScopeError{
		Group:      "",
		Resource:   "nodes",
		Namespaced: false,
	})
	assert.Contains(t, errors, apiresources.UnknownGroupResourceError{
		Group:    "",
		Resource: "processedtemplates",
	})
	assert.Contains(t, errors, apiresources.UnknownGroupResourceError{
		Group:    "apps",
		Resource: "deployments/finalizers",
	})

}
