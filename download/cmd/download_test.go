package cmd_test

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/codeready-toolchain/operator-role-linter/download/cmd"
	"github.com/codeready-toolchain/operator-role-linter/pkg/apiresources"
	"github.com/davecgh/go-spew/spew"

	"github.com/sergi/go-diff/diffmatchpatch"
	"github.com/stretchr/testify/require"
	gock "gopkg.in/h2non/gock.v1"
)

func TestDownload(t *testing.T) {

	t.Run("ok", func(t *testing.T) {
		// given
		clusterURL := "https://cluster:6443"
		token := "foo"
		defer gock.Off()
		client := &http.Client{Transport: &http.Transport{}}
		gock.InterceptClient(client)

		gock.New(clusterURL).
			Get("/api/v1").
			MatchHeader("Authorization", "Bearer "+token).
			Reply(200).
			BodyString(`groupVersion: v1
kind: APIResourceList
resources:
- kind: Binding
  name: bindings
  namespaced: true
  singularName: ""
  verbs:
  - create
- kind: ComponentStatus
  name: componentstatuses
  namespaced: false
  shortNames:
  - cs
  singularName: ""
  verbs:
  - get
  - list
- kind: Service
  name: services/status
  namespaced: true
  singularName: ""
  verbs:
  - get
  - patch
  - update`)

		gock.New(clusterURL).
			Get("/apis").
			MatchHeader("Authorization", "Bearer "+token).
			Reply(200).
			BodyString(`apiVersion: v1
groups:
- name: extensions 
  preferredVersion:
    groupVersion: extensions/v1beta1
    version: v1beta1
  versions:
  - groupVersion: extensions/v1beta1
  version: v1beta1
- name: apps
  preferredVersion:
    groupVersion: apps/v1
    version: v1
  versions:
  - groupVersion: apps/v1
    version: v1
  - groupVersion: apps/v1beta2
    version: v1beta2
  - groupVersion: apps/v1beta1
    version: v1beta1
- name: toolchain.dev.openshift.com
  preferredVersion:
    groupVersion: toolchain.dev.openshift.com/v1alpha1
    version: v1alpha1
  versions:
  - groupVersion: toolchain.dev.openshift.com/v1alpha1
    version: v1alpha1
- name: foo.openshift.com
  preferredVersion:
    groupVersion: foo.openshift.com/v1alpha1
    version: v1alpha1
  versions:
  - groupVersion: foo.openshift.com/v1alpha1
    version: v1alpha1`)

		gock.New(clusterURL).
			Get("/apis/extensions/v1beta1").
			MatchHeader("Authorization", "Bearer "+token).
			Reply(200).
			BodyString(`groupVersion: extensions/v1beta1
kind: APIResourceList
resources:
- kind: DaemonSet
  name: daemonsets
  namespaced: true
  shortNames:
  - ds
  singularName: ""
  verbs:
  - create
  - delete
  - deletecollection
  - get
  - list
  - patch
  - update
  - watch
- kind: DaemonSet
  name: daemonsets/status
  namespaced: true
  singularName: ""
  verbs:
  - get
  - patch
  - update
- kind: Deployment
  name: deployments
  namespaced: true
  shortNames:
  - deploy
  singularName: ""
  verbs:
  - create
  - delete
  - deletecollection
  - get
  - list
  - patch
  - update
  - watch
- group: extensions
  kind: Scale
  name: deployments/scale
  namespaced: true
  singularName: ""
  verbs:
  - get
  - patch
  - update
  version: v1beta1
- kind: Deployment
  name: deployments/status
  namespaced: true
  singularName: ""
  verbs:
  - get
  - patch
  - update`)

		gock.New(clusterURL).
			Get("/apis/apps/v1").
			MatchHeader("Authorization", "Bearer "+token).
			Reply(200).
			BodyString(`apiVersion: v1
groupVersion: apps/v1
kind: APIResourceList
resources:
- kind: ControllerRevision
  name: controllerrevisions
  namespaced: true
  singularName: ""
  verbs:
  - create
  - delete
  - deletecollection
  - get
  - list
  - patch
  - update
  - watch
- categories:
  - all
  kind: DaemonSet
  name: daemonsets
  namespaced: true
  shortNames:
  - ds
  singularName: ""
  verbs:
  - create
  - delete
  - deletecollection
  - get
  - list
  - patch
  - update
  - watch
- kind: DaemonSet
  name: daemonsets/status
  namespaced: true
  singularName: ""
  verbs:
  - get
  - patch
  - update
- categories:
  - all
  kind: Deployment
  name: deployments
  namespaced: true
  shortNames:
  - deploy
  singularName: ""
  verbs:
  - create
  - delete
  - deletecollection
  - get
  - list
  - patch
  - update
  - watch
- group: autoscaling
  kind: Scale
  name: deployments/scale
  namespaced: true
  singularName: ""
  verbs:
  - get
  - patch
  - update
  version: v1
- kind: Deployment
  name: deployments/status
  namespaced: true
  singularName: ""
  verbs:
  - get
  - patch
  - update`)

		// when
		result, err := cmd.Download(clusterURL, token, "toolchain.dev.openshift.com,foo.openshift.com", client)

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
				"daemonsets/status": {
					Namespaced: true,
					Verbs:      []string{"get", "patch", "update"},
				},
				"deployments": {
					Namespaced: true,
					Verbs:      []string{"create", "delete", "deletecollection", "get", "list", "patch", "update", "watch"},
				},
				"deployments/scale": {
					Namespaced: true,
					Verbs:      []string{"get", "patch", "update"},
				},
				"deployments/status": {
					Namespaced: true,
					Verbs:      []string{"get", "patch", "update"},
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
				"daemonsets/status": {
					Namespaced: true,
					Verbs:      []string{"get", "patch", "update"},
				},
				"deployments": {
					Namespaced: true,
					Verbs:      []string{"create", "delete", "deletecollection", "get", "list", "patch", "update", "watch"},
				},
				"deployments/scale": {
					Namespaced: true,
					Verbs:      []string{"get", "patch", "update"},
				},
				"deployments/status": {
					Namespaced: true,
					Verbs:      []string{"get", "patch", "update"},
				},
			},
		}
		compare(t, expected, result)
	})

	t.Run("failures", func(t *testing.T) {

		t.Run("unauthorized", func(t *testing.T) {
			// given
			clusterURL := "https://cluster:6443"
			token := "foo"
			defer gock.Off()
			client := &http.Client{Transport: &http.Transport{}}
			gock.InterceptClient(client)

			gock.New(clusterURL).
				Get("/api/v1").
				MatchHeader("Authorization", "Bearer "+token).
				Reply(401).
				BodyString(`Unauthorized`)

				// when
			_, err := cmd.Download(clusterURL, token, "toolchain.dev.openshift.com,foo.openshift.com", client)

			// then
			require.Error(t, err)

		})

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
		t.Logf("diffs:\n%s", dmp.DiffPrettyText(diffs))
	}

}
