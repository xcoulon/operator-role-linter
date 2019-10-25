package apiresources

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/codeready-toolchain/operator-role-linter/pkg/apiresources/data"
	"github.com/davecgh/go-spew/spew"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

// Verify verifies the file with the given `ruleFile`
// returns an array of linter errors, or a single error if something really wrong happened
func Verify(ruleFile, crdsDir, clusterVersion string) ([]error, error) {
	apiresources, err := loadAPIResources(clusterVersion, crdsDir)

	log.Debugf("now checking %s", ruleFile)
	// now check each rule
	rbacRulesContent, err := ioutil.ReadFile(ruleFile)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to verify %s", ruleFile)
	}
	actual := RBACRules{}
	err = yaml.Unmarshal(rbacRulesContent, &actual)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to verify %s", ruleFile)
	}
	// TODO:
	// 1. support resources '*'
	// 2. missing 'apps.deployments/finalizers' group/resource ?
	errors := []error{}
	for _, actualRule := range actual.Rules {
		for _, actualAPIGroup := range actualRule.APIGroups {
		apigroup_resource:
			for _, actualResource := range actualRule.Resources {
				log.Debugf("verifying group '%s' / resource '%s'", actualAPIGroup, actualResource)
				if group, ok := apiresources[actualAPIGroup]; ok {
					if actualResource == "*" {
						// any verb matches, so move on
						log.Debugf("  match found for resource '%s.%s'", actualAPIGroup, actualResource)
						continue apigroup_resource
					}
					if resource, ok := group[actualResource]; ok {
						if actual.Kind == "Role" && !resource.Namespaced {
							errors = append(errors, InvalidGroupResourceScopeError{
								Group:      actualAPIGroup,
								Resource:   actualResource,
								Namespaced: false,
							})
							continue apigroup_resource
						} else if actual.Kind == "ClusterRole" && resource.Namespaced {
							errors = append(errors, InvalidGroupResourceScopeError{
								Group:      actualAPIGroup,
								Resource:   actualResource,
								Namespaced: true,
							})
							continue apigroup_resource
						}
						for _, actualVerb := range actualRule.Verbs {
							log.Debugf(" checking %s group '%s' / resource '%s'", actualVerb, actualAPIGroup, actualResource)
							if actualVerb == "*" {
								// any verb matches, so move on
								log.Debugf("  match found for resource '%s %s.%s'", actualVerb, actualAPIGroup, actualResource)
								continue apigroup_resource
							}
							for _, v := range resource.Verbs {
								if v == actualVerb {
									// match found, move on
									log.Debugf("  match found for resource '%s %s.%s'", actualVerb, actualAPIGroup, actualResource)
									continue apigroup_resource
								}
							}
							errors = append(errors, UnknownGroupResourceVerbError{
								Group:    actualAPIGroup,
								Resource: actualResource,
								Verb:     actualVerb,
							})
							continue apigroup_resource
						}
					}
					errors = append(errors, UnknownGroupResourceError{
						Group:    actualAPIGroup,
						Resource: actualResource,
					})
				} else {
					errors = append(errors, UnknownGroupError{
						Group: actualAPIGroup,
					})
				}
			}
		}
	}
	return errors, nil
}

func loadAPIResources(clusterVersion, crdsPath string) (APIResources, error) {
	asset, err := data.Asset(clusterVersion + ".yaml")
	if err != nil {
		return APIResources{}, err
	}
	apiresources := APIResources{}
	err = yaml.Unmarshal(asset, &apiresources)
	if err != nil {
		return nil, err
	}
	log.Debugf("loaded %d API Resource groups", len(apiresources))
	crdResources, err := loadCRDs(crdsPath)
	if err != nil {
		return nil, err
	}
	for g, r := range crdResources {
		apiresources[g] = r
	}
	if log.IsLevelEnabled(log.DebugLevel) {
		log.Debug("API resources")
		spew.Dump(apiresources)
	}
	return apiresources, nil
}

func loadCRDs(crdsPath string) (APIResources, error) {
	log.Debugf("loading CRDs in %s", crdsPath)
	var crdFiles []string
	err := filepath.Walk(crdsPath, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			crdFiles = append(crdFiles, path)
		}
		return nil
	})
	if err != nil {
		return APIResources{}, err
	}
	apiresources := APIResources{}
	for _, crdFile := range crdFiles {
		crdContent, err := ioutil.ReadFile(crdFile)
		if err != nil {
			return APIResources{}, err
		}
		actual := CRD{}
		err = yaml.Unmarshal(crdContent, &actual)
		if err != nil {
			return APIResources{}, err
		}
		if _, exists := apiresources[actual.Spec.Group]; !exists {
			apiresources[actual.Spec.Group] = map[string]Resource{}
		}
		namespaced := actual.Spec.Scope == "Namespaced"
		apiresources[actual.Spec.Group][actual.Spec.Names.Plural] = Resource{
			Namespaced: namespaced,
		}
	}
	return apiresources, nil
}

type CRD struct {
	Spec CRDSpec `yaml:"spec"`
}

type CRDSpec struct {
	Group string       `yaml:"group"`
	Names CRDSpecNames `yaml:"names"`
	Scope string       `yaml:"scope"`
}

type CRDSpecNames struct {
	Plural string `yaml:"plural"`
}
