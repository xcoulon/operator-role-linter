package download_test

import (
	"reflect"
	"testing"

	"github.com/codeready-toolchain/operator-role-linter/download"
	"github.com/codeready-toolchain/operator-role-linter/pkg/apiresources"
	"github.com/codeready-toolchain/operator-role-linter/tests"

	"github.com/davecgh/go-spew/spew"
	"github.com/sergi/go-diff/diffmatchpatch"
	"github.com/stretchr/testify/require"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestDownload(t *testing.T) {

	t.Run("ok", func(t *testing.T) {
		// given
		client := tests.NewDiscoveryInterfaceMock(t)
		client.ServerGroupsAndResourcesMock.Set(func() ([]*metav1.APIGroup, []*metav1.APIResourceList, error) {
			return []*metav1.APIGroup{
					{
						Name: "",
						PreferredVersion: metav1.GroupVersionForDiscovery{
							Version: "v1",
						},
					},
					{
						Name: "extensions",
						PreferredVersion: metav1.GroupVersionForDiscovery{
							Version: "v1beta1",
						},
					},
					{
						Name: "apps",
						PreferredVersion: metav1.GroupVersionForDiscovery{
							Version: "v1",
						},
					},
				},
				[]*metav1.APIResourceList{
					{
						GroupVersion: "v1",
						APIResources: []metav1.APIResource{
							{
								Name:       "bindings",
								Kind:       "Binding",
								Namespaced: true,
								Verbs:      []string{"create"},
							},
							{
								Kind:       "ComponentStatus",
								Name:       "componentstatuses",
								Namespaced: false,
								ShortNames: []string{"cs"},
								Verbs:      []string{"get", "list"},
							},
							{
								Kind:       "Service",
								Name:       "services/status",
								Namespaced: true,
								Verbs:      []string{"get", "patch", "update"},
							},
						},
					},
					{
						GroupVersion: "extensions/v1beta1",
						APIResources: []metav1.APIResource{
							{
								Kind:       "DaemonSet",
								Name:       "daemonsets",
								Namespaced: true,
								ShortNames: []string{"ds"},
								Verbs:      []string{"create", "delete", "deletecollection", "get", "list", "patch", "update", "watch"},
							},
							{
								Kind:       "Deployment",
								Name:       "deployments",
								Namespaced: true,
								ShortNames: []string{"deploy"},
								Verbs:      []string{"create", "delete", "deletecollection", "get", "list", "patch", "update", "watch"},
							},
						},
					},
					{
						GroupVersion: "apps/v1",
						APIResources: []metav1.APIResource{
							{
								Kind:       "ControllerRevision",
								Name:       "controllerrevisions",
								Namespaced: true,
								Verbs:      []string{"create", "delete", "deletecollection", "get", "list", "patch", "update", "watch"},
							},
							{
								Kind:       "DaemonSet",
								Name:       "daemonsets",
								Namespaced: true,
								ShortNames: []string{"ds"},
								Verbs:      []string{"create", "delete", "deletecollection", "get", "list", "patch", "update", "watch"},
							},
						},
					},
					{
						GroupVersion: "toolchain.dev.openshift.com/v1alpha1",
						APIResources: []metav1.APIResource{
							// whetever...
						},
					},
				}, nil
		})

		// when
		result, err := download.Download(client, "toolchain.dev.openshift.com,foo.openshift.com")

		// then
		require.NoError(t, err)
		expected := apiresources.APIResources{
			"": map[string]apiresources.Resource{
				"bindings": {
					Namespaced: true,
					Verbs:      []string{"create"},
				},
				"componentstatuses": {
					Namespaced: false,
					Verbs:      []string{"get", "list"},
				},
				"services/status": {
					Namespaced: true,
					Verbs:      []string{"get", "patch", "update"},
				},
			},
			"extensions": map[string]apiresources.Resource{
				"daemonsets": {
					Namespaced: true,
					Verbs:      []string{"create", "delete", "deletecollection", "get", "list", "patch", "update", "watch"},
				},
				"deployments": {
					Namespaced: true,
					Verbs:      []string{"create", "delete", "deletecollection", "get", "list", "patch", "update", "watch"},
				},
			},
			"apps": map[string]apiresources.Resource{
				"controllerrevisions": {
					Namespaced: true,
					Verbs:      []string{"create", "delete", "deletecollection", "get", "list", "patch", "update", "watch"},
				},
				"daemonsets": {
					Namespaced: true,
					Verbs:      []string{"create", "delete", "deletecollection", "get", "list", "patch", "update", "watch"},
				},
			},
		}
		compare(t, expected, result)
	})

}

// compare compares the 'actual' vs 'expected' values and outputs a 'diff' if needed
func compare(t *testing.T, expected interface{}, actual interface{}) {
	if !reflect.DeepEqual(actual, expected) {
		a := spew.Sdump(actual)
		e := spew.Sdump(expected)
		t.Logf("actual:\n%s", a)
		t.Logf("expected:\n%v", e)
		dmp := diffmatchpatch.New()
		diffs := dmp.DiffMain(a, e, true)
		t.Errorf("unmatched result: %s", dmp.DiffPrettyText(diffs))
	}

}
