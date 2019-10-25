package apiresources

import "fmt"

// APIResources the API resources indexed by their group name and name
type APIResources map[string]map[string]Resource

// Resource an API resource, which indicates for a named resource:
// - if it is namespaced
// - the verbs that are allowed
type Resource struct {
	Namespaced bool
	Verbs      []string
}

// kind: Role
// apiVersion: rbac.authorization.k8s.io/v1
// metadata:
//   name: host-operator
// rules:
// - apiGroups:
//   - ""
//   resources:
//   - pods
//   - services
//   - endpoints
//   - persistentvolumeclaims
//   - events
//   - configmaps
//   - secrets
//   verbs:
//   - "*"

// RBACRules the whole Role or ClusterRole rules (depending on the Kind)
type RBACRules struct {
	Kind  string
	Rules []Rule
}

// Rule an RBAC rule
type Rule struct {
	APIGroups     []string `yaml:"apiGroups"`
	Resources     []string `yaml:"resources"`
	ResourceNames []string `yaml:"resourceNames"`
	Verbs         []string `yaml:"verbs"`
}

// UnknownGroupError linter error which is reported when the API group is unknown
type UnknownGroupError struct {
	Group string
}

func (e UnknownGroupError) Error() string {
	return fmt.Sprintf("unknown API Group: '%s'", e.Group)
}

// UnknownGroupResourceError linter error which is reported when the API resource is unknown
type UnknownGroupResourceError struct {
	Group    string
	Resource string
}

func (e UnknownGroupResourceError) Error() string {
	return fmt.Sprintf("unknown resource '%s' in API Group '%s'", e.Resource, e.Group)
}

// InvalidGroupResourceScopeError linter error which is reported when the API resource is namespaced (or not)
type InvalidGroupResourceScopeError struct {
	Group      string
	Resource   string
	Namespaced bool
}

func (e InvalidGroupResourceScopeError) Error() string {
	if e.Namespaced {
		return fmt.Sprintf("invalid API resource scope: '%s' in API Group '%s' is namespaced", e.Resource, e.Group)
	}
	return fmt.Sprintf("invalid API resource scope: '%s' in API Group '%s' is not namespaced", e.Resource, e.Group)
}

// UnknownGroupResourceVerbError linter error which is reported when verb on the API resource is unknown
type UnknownGroupResourceVerbError struct {
	Group    string
	Resource string
	Verb     string
}

func (e UnknownGroupResourceVerbError) Error() string {
	return fmt.Sprintf("unknown verb '%s' for resource '%s' in API Group '%s'", e.Verb, e.Resource, e.Group)
}
