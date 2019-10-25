// Code generated for package data by go-bindata DO NOT EDIT. (@generated)
// sources:
// pkg/apiresources/data/apiresources.yaml
// pkg/apiresources/data/apiresources_asset.go
package data

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)
type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _apiresourcesYaml = []byte(`"":
  bindings:
    namespaced: true
    verbs:
    - create
  componentstatuses:
    namespaced: false
    verbs:
    - get
    - list
  configmaps:
    namespaced: true
    verbs:
    - create
    - delete
    - deletecollection
    - get
    - list
    - patch
    - update
    - watch
  endpoints:
    namespaced: true
    verbs:
    - create
    - delete
    - deletecollection
    - get
    - list
    - patch
    - update
    - watch
  events:
    namespaced: true
    verbs:
    - create
    - delete
    - deletecollection
    - get
    - list
    - patch
    - update
    - watch
  limitranges:
    namespaced: true
    verbs:
    - create
    - delete
    - deletecollection
    - get
    - list
    - patch
    - update
    - watch
  namespaces:
    namespaced: false
    verbs:
    - create
    - delete
    - get
    - list
    - patch
    - update
    - watch
  namespaces/finalize:
    namespaced: false
    verbs:
    - update
  namespaces/status:
    namespaced: false
    verbs:
    - get
    - patch
    - update
  nodes:
    namespaced: false
    verbs:
    - create
    - delete
    - deletecollection
    - get
    - list
    - patch
    - update
    - watch
  nodes/proxy:
    namespaced: false
    verbs:
    - create
    - delete
    - get
    - patch
    - update
  nodes/status:
    namespaced: false
    verbs:
    - get
    - patch
    - update
  persistentvolumeclaims:
    namespaced: true
    verbs:
    - create
    - delete
    - deletecollection
    - get
    - list
    - patch
    - update
    - watch
  persistentvolumeclaims/status:
    namespaced: true
    verbs:
    - get
    - patch
    - update
  persistentvolumes:
    namespaced: false
    verbs:
    - create
    - delete
    - deletecollection
    - get
    - list
    - patch
    - update
    - watch
  persistentvolumes/status:
    namespaced: false
    verbs:
    - get
    - patch
    - update
  pods:
    namespaced: true
    verbs:
    - create
    - delete
    - deletecollection
    - get
    - list
    - patch
    - update
    - watch
  pods/attach:
    namespaced: true
    verbs:
    - create
    - get
  pods/binding:
    namespaced: true
    verbs:
    - create
  pods/eviction:
    namespaced: true
    verbs:
    - create
  pods/exec:
    namespaced: true
    verbs:
    - create
    - get
  pods/log:
    namespaced: true
    verbs:
    - get
  pods/portforward:
    namespaced: true
    verbs:
    - create
    - get
  pods/proxy:
    namespaced: true
    verbs:
    - create
    - delete
    - get
    - patch
    - update
  pods/status:
    namespaced: true
    verbs:
    - get
    - patch
    - update
  podtemplates:
    namespaced: true
    verbs:
    - create
    - delete
    - deletecollection
    - get
    - list
    - patch
    - update
    - watch
  replicationcontrollers:
    namespaced: true
    verbs:
    - create
    - delete
    - deletecollection
    - get
    - list
    - patch
    - update
    - watch
  replicationcontrollers/scale:
    namespaced: true
    verbs:
    - get
    - patch
    - update
  replicationcontrollers/status:
    namespaced: true
    verbs:
    - get
    - patch
    - update
  resourcequotas:
    namespaced: true
    verbs:
    - create
    - delete
    - deletecollection
    - get
    - list
    - patch
    - update
    - watch
  resourcequotas/status:
    namespaced: true
    verbs:
    - get
    - patch
    - update
  secrets:
    namespaced: true
    verbs:
    - create
    - delete
    - deletecollection
    - get
    - list
    - patch
    - update
    - watch
  serviceaccounts:
    namespaced: true
    verbs:
    - create
    - delete
    - deletecollection
    - get
    - list
    - patch
    - update
    - watch
  services:
    namespaced: true
    verbs:
    - create
    - delete
    - get
    - list
    - patch
    - update
    - watch
  services/proxy:
    namespaced: true
    verbs:
    - create
    - delete
    - get
    - patch
    - update
  services/status:
    namespaced: true
    verbs:
    - get
    - patch
    - update
admissionregistration.k8s.io:
  mutatingwebhookconfigurations:
    namespaced: false
    verbs:
    - create
    - delete
    - deletecollection
    - get
    - list
    - patch
    - update
    - watch
  validatingwebhookconfigurations:
    namespaced: false
    verbs:
    - create
    - delete
    - deletecollection
    - get
    - list
    - patch
    - update
    - watch
apiextensions.k8s.io:
  customresourcedefinitions:
    namespaced: false
    verbs:
    - create
    - delete
    - deletecollection
    - get
    - list
    - patch
    - update
    - watch
  customresourcedefinitions/status:
    namespaced: false
    verbs:
    - get
    - patch
    - update
apiregistration.k8s.io:
  apiservices:
    namespaced: false
    verbs:
    - create
    - delete
    - deletecollection
    - get
    - list
    - patch
    - update
    - watch
  apiservices/status:
    namespaced: false
    verbs:
    - get
    - patch
    - update
apps:
  controllerrevisions:
    namespaced: true
    verbs:
    - create
    - delete
    - deletecollection
    - get
    - list
    - patch
    - update
    - watch
  daemonsets:
    namespaced: true
    verbs:
    - create
    - delete
    - deletecollection
    - get
    - list
    - patch
    - update
    - watch
  daemonsets/status:
    namespaced: true
    verbs:
    - get
    - patch
    - update
  deployments:
    namespaced: true
    verbs:
    - create
    - delete
    - deletecollection
    - get
    - list
    - patch
    - update
    - watch
  deployments/scale:
    namespaced: true
    verbs:
    - get
    - patch
    - update
  deployments/status:
    namespaced: true
    verbs:
    - get
    - patch
    - update
  replicasets:
    namespaced: true
    verbs:
    - create
    - delete
    - deletecollection
    - get
    - list
    - patch
    - update
    - watch
  replicasets/scale:
    namespaced: true
    verbs:
    - get
    - patch
    - update
  replicasets/status:
    namespaced: true
    verbs:
    - get
    - patch
    - update
  statefulsets:
    namespaced: true
    verbs:
    - create
    - delete
    - deletecollection
    - get
    - list
    - patch
    - update
    - watch
  statefulsets/scale:
    namespaced: true
    verbs:
    - get
    - patch
    - update
  statefulsets/status:
    namespaced: true
    verbs:
    - get
    - patch
    - update
apps.openshift.io:
  deploymentconfigs:
    namespaced: true
    verbs:
    - create
    - delete
    - deletecollection
    - get
    - list
    - patch
    - update
    - watch
  deploymentconfigs/instantiate:
    namespaced: true
    verbs:
    - create
  deploymentconfigs/log:
    namespaced: true
    verbs:
    - get
  deploymentconfigs/rollback:
    namespaced: true
    verbs:
    - create
  deploymentconfigs/scale:
    namespaced: true
    verbs:
    - get
    - patch
    - update
  deploymentconfigs/status:
    namespaced: true
    verbs:
    - get
    - patch
    - update
authentication.k8s.io:
  tokenreviews:
    namespaced: false
    verbs:
    - create
authorization.k8s.io:
  localsubjectaccessreviews:
    namespaced: true
    verbs:
    - create
  selfsubjectaccessreviews:
    namespaced: false
    verbs:
    - create
  selfsubjectrulesreviews:
    namespaced: false
    verbs:
    - create
  subjectaccessreviews:
    namespaced: false
    verbs:
    - create
authorization.openshift.io:
  clusterrolebindings:
    namespaced: false
    verbs:
    - create
    - delete
    - get
    - list
    - patch
    - update
  clusterroles:
    namespaced: false
    verbs:
    - create
    - delete
    - get
    - list
    - patch
    - update
  localresourceaccessreviews:
    namespaced: true
    verbs:
    - create
  localsubjectaccessreviews:
    namespaced: true
    verbs:
    - create
  resourceaccessreviews:
    namespaced: false
    verbs:
    - create
  rolebindingrestrictions:
    namespaced: true
    verbs:
    - create
    - delete
    - deletecollection
    - get
    - list
    - patch
    - update
    - watch
  rolebindings:
    namespaced: true
    verbs:
    - create
    - delete
    - get
    - list
    - patch
    - update
  roles:
    namespaced: true
    verbs:
    - create
    - delete
    - get
    - list
    - patch
    - update
  selfsubjectrulesreviews:
    namespaced: true
    verbs:
    - create
  subjectaccessreviews:
    namespaced: false
    verbs:
    - create
  subjectrulesreviews:
    namespaced: true
    verbs:
    - create
autoscaling:
  horizontalpodautoscalers:
    namespaced: true
    verbs:
    - create
    - delete
    - deletecollection
    - get
    - list
    - patch
    - update
    - watch
  horizontalpodautoscalers/status:
    namespaced: true
    verbs:
    - get
    - patch
    - update
autoscaling.openshift.io:
  clusterautoscalers:
    namespaced: false
    verbs:
    - delete
    - deletecollection
    - get
    - list
    - patch
    - create
    - update
    - watch
  clusterautoscalers/status:
    namespaced: false
    verbs:
    - get
    - patch
    - update
batch:
  jobs:
    namespaced: true
    verbs:
    - create
    - delete
    - deletecollection
    - get
    - list
    - patch
    - update
    - watch
  jobs/status:
    namespaced: true
    verbs:
    - get
    - patch
    - update
build.openshift.io:
  buildconfigs:
    namespaced: true
    verbs:
    - create
    - delete
    - deletecollection
    - get
    - list
    - patch
    - update
    - watch
  buildconfigs/instantiate:
    namespaced: true
    verbs:
    - create
  buildconfigs/instantiatebinary:
    namespaced: true
    verbs:
    - create
  buildconfigs/webhooks:
    namespaced: true
    verbs:
    - create
  builds:
    namespaced: true
    verbs:
    - create
    - delete
    - deletecollection
    - get
    - list
    - patch
    - update
    - watch
  builds/clone:
    namespaced: true
    verbs:
    - create
  builds/details:
    namespaced: true
    verbs:
    - update
  builds/log:
    namespaced: true
    verbs:
    - get
certificates.k8s.io:
  certificatesigningrequests:
    namespaced: false
    verbs:
    - create
    - delete
    - deletecollection
    - get
    - list
    - patch
    - update
    - watch
  certificatesigningrequests/approval:
    namespaced: false
    verbs:
    - update
  certificatesigningrequests/status:
    namespaced: false
    verbs:
    - get
    - patch
    - update
cloudcredential.openshift.io:
  credentialsrequests:
    namespaced: true
    verbs:
    - delete
    - deletecollection
    - get
    - list
    - patch
    - create
    - update
    - watch
  credentialsrequests/status:
    namespaced: true
    verbs:
    - get
    - patch
    - update
config.openshift.io:
  apiservers:
    namespaced: false
    verbs:
    - delete
    - deletecollection
    - get
    - list
    - patch
    - create
    - update
    - watch
  authentications:
    namespaced: false
    verbs:
    - delete
    - deletecollection
    - get
    - list
    - patch
    - create
    - update
    - watch
  authentications/status:
    namespaced: false
    verbs:
    - get
    - patch
    - update
  builds:
    namespaced: false
    verbs:
    - delete
    - deletecollection
    - get
    - list
    - patch
    - create
    - update
    - watch
  clusteroperators:
    namespaced: false
    verbs:
    - delete
    - deletecollection
    - get
    - list
    - patch
    - create
    - update
    - watch
  clusteroperators/status:
    namespaced: false
    verbs:
    - get
    - patch
    - update
  clusterversions:
    namespaced: false
    verbs:
    - delete
    - deletecollection
    - get
    - list
    - patch
    - create
    - update
    - watch
  clusterversions/status:
    namespaced: false
    verbs:
    - get
    - patch
    - update
  consoles:
    namespaced: false
    verbs:
    - delete
    - deletecollection
    - get
    - list
    - patch
    - create
    - update
    - watch
  consoles/status:
    namespaced: false
    verbs:
    - get
    - patch
    - update
  dnses:
    namespaced: false
    verbs:
    - delete
    - deletecollection
    - get
    - list
    - patch
    - create
    - update
    - watch
  featuregates:
    namespaced: false
    verbs:
    - delete
    - deletecollection
    - get
    - list
    - patch
    - create
    - update
    - watch
  images:
    namespaced: false
    verbs:
    - delete
    - deletecollection
    - get
    - list
    - patch
    - create
    - update
    - watch
  images/status:
    namespaced: false
    verbs:
    - get
    - patch
    - update
  infrastructures:
    namespaced: false
    verbs:
    - delete
    - deletecollection
    - get
    - list
    - patch
    - create
    - update
    - watch
  ingresses:
    namespaced: false
    verbs:
    - delete
    - deletecollection
    - get
    - list
    - patch
    - create
    - update
    - watch
  networks:
    namespaced: false
    verbs:
    - delete
    - deletecollection
    - get
    - list
    - patch
    - create
    - update
    - watch
  oauths:
    namespaced: false
    verbs:
    - delete
    - deletecollection
    - get
    - list
    - patch
    - create
    - update
    - watch
  oauths/status:
    namespaced: false
    verbs:
    - get
    - patch
    - update
  projects:
    namespaced: false
    verbs:
    - delete
    - deletecollection
    - get
    - list
    - patch
    - create
    - update
    - watch
  schedulers:
    namespaced: false
    verbs:
    - delete
    - deletecollection
    - get
    - list
    - patch
    - create
    - update
    - watch
coordination.k8s.io:
  leases:
    namespaced: true
    verbs:
    - create
    - delete
    - deletecollection
    - get
    - list
    - patch
    - update
    - watch
core.kubefed.k8s.io:
  kubefedclusters:
    namespaced: true
    verbs:
    - delete
    - deletecollection
    - get
    - list
    - patch
    - create
    - update
    - watch
  kubefedclusters/status:
    namespaced: true
    verbs:
    - get
    - patch
    - update
events.k8s.io:
  events:
    namespaced: true
    verbs:
    - create
    - delete
    - deletecollection
    - get
    - list
    - patch
    - update
    - watch
extensions:
  daemonsets:
    namespaced: true
    verbs:
    - create
    - delete
    - deletecollection
    - get
    - list
    - patch
    - update
    - watch
  daemonsets/status:
    namespaced: true
    verbs:
    - get
    - patch
    - update
  deployments:
    namespaced: true
    verbs:
    - create
    - delete
    - deletecollection
    - get
    - list
    - patch
    - update
    - watch
  deployments/rollback:
    namespaced: true
    verbs:
    - create
  deployments/scale:
    namespaced: true
    verbs:
    - get
    - patch
    - update
  deployments/status:
    namespaced: true
    verbs:
    - get
    - patch
    - update
  ingresses:
    namespaced: true
    verbs:
    - create
    - delete
    - deletecollection
    - get
    - list
    - patch
    - update
    - watch
  ingresses/status:
    namespaced: true
    verbs:
    - get
    - patch
    - update
  networkpolicies:
    namespaced: true
    verbs:
    - create
    - delete
    - deletecollection
    - get
    - list
    - patch
    - update
    - watch
  podsecuritypolicies:
    namespaced: false
    verbs:
    - create
    - delete
    - deletecollection
    - get
    - list
    - patch
    - update
    - watch
  replicasets:
    namespaced: true
    verbs:
    - create
    - delete
    - deletecollection
    - get
    - list
    - patch
    - update
    - watch
  replicasets/scale:
    namespaced: true
    verbs:
    - get
    - patch
    - update
  replicasets/status:
    namespaced: true
    verbs:
    - get
    - patch
    - update
  replicationcontrollers:
    namespaced: true
    verbs: []
  replicationcontrollers/scale:
    namespaced: true
    verbs:
    - get
    - patch
    - update
healthchecking.openshift.io:
  machinehealthchecks:
    namespaced: true
    verbs:
    - delete
    - deletecollection
    - get
    - list
    - patch
    - create
    - update
    - watch
image.openshift.io:
  images:
    namespaced: false
    verbs:
    - create
    - delete
    - deletecollection
    - get
    - list
    - patch
    - update
    - watch
  imagesignatures:
    namespaced: false
    verbs:
    - create
    - delete
  imagestreamimages:
    namespaced: true
    verbs:
    - get
  imagestreamimports:
    namespaced: true
    verbs:
    - create
  imagestreammappings:
    namespaced: true
    verbs:
    - create
  imagestreams:
    namespaced: true
    verbs:
    - create
    - delete
    - deletecollection
    - get
    - list
    - patch
    - update
    - watch
  imagestreams/layers:
    namespaced: true
    verbs:
    - get
  imagestreams/secrets:
    namespaced: true
    verbs:
    - get
  imagestreams/status:
    namespaced: true
    verbs:
    - get
    - patch
    - update
  imagestreamtags:
    namespaced: true
    verbs:
    - create
    - delete
    - get
    - list
    - patch
    - update
imageregistry.operator.openshift.io:
  configs:
    namespaced: false
    verbs:
    - delete
    - deletecollection
    - get
    - list
    - patch
    - create
    - update
    - watch
  configs/status:
    namespaced: false
    verbs:
    - get
    - patch
    - update
k8s.cni.cncf.io:
  network-attachment-definitions:
    namespaced: true
    verbs:
    - delete
    - deletecollection
    - get
    - list
    - patch
    - create
    - update
    - watch
machine.openshift.io:
  machines:
    namespaced: true
    verbs:
    - delete
    - deletecollection
    - get
    - list
    - patch
    - create
    - update
    - watch
  machines/status:
    namespaced: true
    verbs:
    - get
    - patch
    - update
  machinesets:
    namespaced: true
    verbs:
    - delete
    - deletecollection
    - get
    - list
    - patch
    - create
    - update
    - watch
  machinesets/scale:
    namespaced: true
    verbs:
    - get
    - patch
    - update
  machinesets/status:
    namespaced: true
    verbs:
    - get
    - patch
    - update
machineconfiguration.openshift.io:
  containerruntimeconfigs:
    namespaced: false
    verbs:
    - delete
    - deletecollection
    - get
    - list
    - patch
    - create
    - update
    - watch
  containerruntimeconfigs/status:
    namespaced: false
    verbs:
    - get
    - patch
    - update
  controllerconfigs:
    namespaced: false
    verbs:
    - delete
    - deletecollection
    - get
    - list
    - patch
    - create
    - update
    - watch
  controllerconfigs/status:
    namespaced: false
    verbs:
    - get
    - patch
    - update
  kubeletconfigs:
    namespaced: false
    verbs:
    - delete
    - deletecollection
    - get
    - list
    - patch
    - create
    - update
    - watch
  kubeletconfigs/status:
    namespaced: false
    verbs:
    - get
    - patch
    - update
  machineconfigpools:
    namespaced: false
    verbs:
    - delete
    - deletecollection
    - get
    - list
    - patch
    - create
    - update
    - watch
  machineconfigpools/status:
    namespaced: false
    verbs:
    - get
    - patch
    - update
  machineconfigs:
    namespaced: false
    verbs:
    - delete
    - deletecollection
    - get
    - list
    - patch
    - create
    - update
    - watch
  mcoconfigs:
    namespaced: true
    verbs:
    - delete
    - deletecollection
    - get
    - list
    - patch
    - create
    - update
    - watch
metrics.k8s.io:
  nodes:
    namespaced: false
    verbs:
    - get
    - list
  pods:
    namespaced: true
    verbs:
    - get
    - list
monitoring.coreos.com:
  alertmanagers:
    namespaced: true
    verbs:
    - delete
    - deletecollection
    - get
    - list
    - patch
    - create
    - update
    - watch
  prometheuses:
    namespaced: true
    verbs:
    - delete
    - deletecollection
    - get
    - list
    - patch
    - create
    - update
    - watch
  prometheusrules:
    namespaced: true
    verbs:
    - delete
    - deletecollection
    - get
    - list
    - patch
    - create
    - update
    - watch
  servicemonitors:
    namespaced: true
    verbs:
    - delete
    - deletecollection
    - get
    - list
    - patch
    - create
    - update
    - watch
network.openshift.io:
  clusternetworks:
    namespaced: false
    verbs:
    - delete
    - deletecollection
    - get
    - list
    - patch
    - create
    - update
    - watch
  egressnetworkpolicies:
    namespaced: true
    verbs:
    - delete
    - deletecollection
    - get
    - list
    - patch
    - create
    - update
    - watch
  hostsubnets:
    namespaced: false
    verbs:
    - delete
    - deletecollection
    - get
    - list
    - patch
    - create
    - update
    - watch
  netnamespaces:
    namespaced: false
    verbs:
    - delete
    - deletecollection
    - get
    - list
    - patch
    - create
    - update
    - watch
networking.k8s.io:
  networkpolicies:
    namespaced: true
    verbs:
    - create
    - delete
    - deletecollection
    - get
    - list
    - patch
    - update
    - watch
oauth.openshift.io:
  oauthaccesstokens:
    namespaced: false
    verbs:
    - create
    - delete
    - deletecollection
    - get
    - list
    - patch
    - update
    - watch
  oauthauthorizetokens:
    namespaced: false
    verbs:
    - create
    - delete
    - deletecollection
    - get
    - list
    - patch
    - update
    - watch
  oauthclientauthorizations:
    namespaced: false
    verbs:
    - create
    - delete
    - deletecollection
    - get
    - list
    - patch
    - update
    - watch
  oauthclients:
    namespaced: false
    verbs:
    - create
    - delete
    - deletecollection
    - get
    - list
    - patch
    - update
    - watch
operator.openshift.io:
  authentications:
    namespaced: false
    verbs:
    - delete
    - deletecollection
    - get
    - list
    - patch
    - create
    - update
    - watch
  authentications/status:
    namespaced: false
    verbs:
    - get
    - patch
    - update
  consoles:
    namespaced: false
    verbs:
    - delete
    - deletecollection
    - get
    - list
    - patch
    - create
    - update
    - watch
  consoles/status:
    namespaced: false
    verbs:
    - get
    - patch
    - update
  dnses:
    namespaced: false
    verbs:
    - delete
    - deletecollection
    - get
    - list
    - patch
    - create
    - update
    - watch
  dnses/status:
    namespaced: false
    verbs:
    - get
    - patch
    - update
  ingresscontrollers:
    namespaced: true
    verbs:
    - delete
    - deletecollection
    - get
    - list
    - patch
    - create
    - update
    - watch
  ingresscontrollers/scale:
    namespaced: true
    verbs:
    - get
    - patch
    - update
  ingresscontrollers/status:
    namespaced: true
    verbs:
    - get
    - patch
    - update
  kubeapiservers:
    namespaced: false
    verbs:
    - delete
    - deletecollection
    - get
    - list
    - patch
    - create
    - update
    - watch
  kubeapiservers/status:
    namespaced: false
    verbs:
    - get
    - patch
    - update
  kubecontrollermanagers:
    namespaced: false
    verbs:
    - delete
    - deletecollection
    - get
    - list
    - patch
    - create
    - update
    - watch
  kubecontrollermanagers/status:
    namespaced: false
    verbs:
    - get
    - patch
    - update
  kubeschedulers:
    namespaced: false
    verbs:
    - delete
    - deletecollection
    - get
    - list
    - patch
    - create
    - update
    - watch
  kubeschedulers/status:
    namespaced: false
    verbs:
    - get
    - patch
    - update
  networks:
    namespaced: false
    verbs:
    - delete
    - deletecollection
    - get
    - list
    - patch
    - create
    - update
    - watch
  openshiftapiservers:
    namespaced: false
    verbs:
    - delete
    - deletecollection
    - get
    - list
    - patch
    - create
    - update
    - watch
  openshiftapiservers/status:
    namespaced: false
    verbs:
    - get
    - patch
    - update
  openshiftcontrollermanagers:
    namespaced: false
    verbs:
    - delete
    - deletecollection
    - get
    - list
    - patch
    - create
    - update
    - watch
  openshiftcontrollermanagers/status:
    namespaced: false
    verbs:
    - get
    - patch
    - update
  servicecas:
    namespaced: false
    verbs:
    - delete
    - deletecollection
    - get
    - list
    - patch
    - create
    - update
    - watch
  servicecas/status:
    namespaced: false
    verbs:
    - get
    - patch
    - update
  servicecatalogapiservers:
    namespaced: false
    verbs:
    - delete
    - deletecollection
    - get
    - list
    - patch
    - create
    - update
    - watch
  servicecatalogapiservers/status:
    namespaced: false
    verbs:
    - get
    - patch
    - update
  servicecatalogcontrollermanagers:
    namespaced: false
    verbs:
    - delete
    - deletecollection
    - get
    - list
    - patch
    - create
    - update
    - watch
  servicecatalogcontrollermanagers/status:
    namespaced: false
    verbs:
    - get
    - patch
    - update
operators.coreos.com:
  catalogsourceconfigs:
    namespaced: true
    verbs:
    - delete
    - deletecollection
    - get
    - list
    - patch
    - create
    - update
    - watch
  operatorgroups:
    namespaced: true
    verbs:
    - delete
    - deletecollection
    - get
    - list
    - patch
    - create
    - update
    - watch
  operatorgroups/status:
    namespaced: true
    verbs:
    - get
    - patch
    - update
  operatorsources:
    namespaced: true
    verbs:
    - delete
    - deletecollection
    - get
    - list
    - patch
    - create
    - update
    - watch
packages.operators.coreos.com:
  packagemanifests:
    namespaced: true
    verbs:
    - get
    - list
policy:
  poddisruptionbudgets:
    namespaced: true
    verbs:
    - create
    - delete
    - deletecollection
    - get
    - list
    - patch
    - update
    - watch
  poddisruptionbudgets/status:
    namespaced: true
    verbs:
    - get
    - patch
    - update
  podsecuritypolicies:
    namespaced: false
    verbs:
    - create
    - delete
    - deletecollection
    - get
    - list
    - patch
    - update
    - watch
project.openshift.io:
  projectrequests:
    namespaced: false
    verbs:
    - create
    - list
  projects:
    namespaced: false
    verbs:
    - create
    - delete
    - get
    - list
    - patch
    - update
    - watch
quota.openshift.io:
  appliedclusterresourcequotas:
    namespaced: true
    verbs:
    - get
    - list
  clusterresourcequotas:
    namespaced: false
    verbs:
    - create
    - delete
    - deletecollection
    - get
    - list
    - patch
    - update
    - watch
  clusterresourcequotas/status:
    namespaced: false
    verbs:
    - get
    - patch
    - update
rbac.authorization.k8s.io:
  clusterrolebindings:
    namespaced: false
    verbs:
    - create
    - delete
    - deletecollection
    - get
    - list
    - patch
    - update
    - watch
  clusterroles:
    namespaced: false
    verbs:
    - create
    - delete
    - deletecollection
    - get
    - list
    - patch
    - update
    - watch
  rolebindings:
    namespaced: true
    verbs:
    - create
    - delete
    - deletecollection
    - get
    - list
    - patch
    - update
    - watch
  roles:
    namespaced: true
    verbs:
    - create
    - delete
    - deletecollection
    - get
    - list
    - patch
    - update
    - watch
route.openshift.io:
  routes:
    namespaced: true
    verbs:
    - create
    - delete
    - deletecollection
    - get
    - list
    - patch
    - update
    - watch
  routes/status:
    namespaced: true
    verbs:
    - get
    - patch
    - update
samples.operator.openshift.io:
  configs:
    namespaced: false
    verbs:
    - delete
    - deletecollection
    - get
    - list
    - patch
    - create
    - update
    - watch
  configs/status:
    namespaced: false
    verbs:
    - get
    - patch
    - update
scheduling.k8s.io:
  priorityclasses:
    namespaced: false
    verbs:
    - create
    - delete
    - deletecollection
    - get
    - list
    - patch
    - update
    - watch
security.openshift.io:
  podsecuritypolicyreviews:
    namespaced: true
    verbs:
    - create
  podsecuritypolicyselfsubjectreviews:
    namespaced: true
    verbs:
    - create
  podsecuritypolicysubjectreviews:
    namespaced: true
    verbs:
    - create
  rangeallocations:
    namespaced: false
    verbs:
    - create
    - delete
    - deletecollection
    - get
    - list
    - patch
    - update
    - watch
  securitycontextconstraints:
    namespaced: false
    verbs:
    - create
    - delete
    - deletecollection
    - get
    - list
    - patch
    - update
    - watch
storage.k8s.io:
  storageclasses:
    namespaced: false
    verbs:
    - create
    - delete
    - deletecollection
    - get
    - list
    - patch
    - update
    - watch
  volumeattachments:
    namespaced: false
    verbs:
    - create
    - delete
    - deletecollection
    - get
    - list
    - patch
    - update
    - watch
  volumeattachments/status:
    namespaced: false
    verbs:
    - get
    - patch
    - update
template.openshift.io:
  brokertemplateinstances:
    namespaced: false
    verbs:
    - create
    - delete
    - deletecollection
    - get
    - list
    - patch
    - update
    - watch
  processedtemplates:
    namespaced: true
    verbs:
    - create
  templateinstances:
    namespaced: true
    verbs:
    - create
    - delete
    - deletecollection
    - get
    - list
    - patch
    - update
    - watch
  templateinstances/status:
    namespaced: true
    verbs:
    - get
    - patch
    - update
  templates:
    namespaced: true
    verbs:
    - create
    - delete
    - deletecollection
    - get
    - list
    - patch
    - update
    - watch
tuned.openshift.io:
  tuneds:
    namespaced: true
    verbs:
    - delete
    - deletecollection
    - get
    - list
    - patch
    - create
    - update
    - watch
user.openshift.io:
  groups:
    namespaced: false
    verbs:
    - create
    - delete
    - deletecollection
    - get
    - list
    - patch
    - update
    - watch
  identities:
    namespaced: false
    verbs:
    - create
    - delete
    - deletecollection
    - get
    - list
    - patch
    - update
    - watch
  useridentitymappings:
    namespaced: false
    verbs:
    - create
    - delete
    - get
    - patch
    - update
  users:
    namespaced: false
    verbs:
    - create
    - delete
    - deletecollection
    - get
    - list
    - patch
    - update
    - watch
`)

func apiresourcesYamlBytes() ([]byte, error) {
	return _apiresourcesYaml, nil
}

func apiresourcesYaml() (*asset, error) {
	bytes, err := apiresourcesYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "apiresources.yaml", size: 30585, mode: os.FileMode(420), modTime: time.Unix(1571988426, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _apiresources_assetGo = []byte(`// Code generated for package data by go-bindata DO NOT EDIT. (@generated)
// sources:
// pkg/apiresources/data/apiresources.yaml
// pkg/apiresources/data/apiresources_asset.go
package data

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)
type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _apiresourcesYaml = []byte(`+"`"+`"":
  bindings:
    namespaced: true
    verbs:
    - create
  componentstatuses:
    namespaced: false
    verbs:
    - get
    - list
  configmaps:
    namespaced: true
    verbs:
    - create
    - delete
    - deletecollection
    - get
    - list
    - patch
    - update
    - watch
  endpoints:
    namespaced: true
    verbs:
    - create
    - delete
    - deletecollection
    - get
    - list
    - patch
    - update
    - watch
  events:
    namespaced: true
    verbs:
    - create
    - delete
    - deletecollection
    - get
    - list
    - patch
    - update
    - watch
  limitranges:
    namespaced: true
    verbs:
    - create
    - delete
    - deletecollection
    - get
    - list
    - patch
    - update
    - watch
  namespaces:
    namespaced: false
    verbs:
    - create
    - delete
    - get
    - list
    - patch
    - update
    - watch
  namespaces/finalize:
    namespaced: false
    verbs:
    - update
  namespaces/status:
    namespaced: false
    verbs:
    - get
    - patch
    - update
  nodes:
    namespaced: false
    verbs:
    - create
    - delete
    - deletecollection
    - get
    - list
    - patch
    - update
    - watch
  nodes/proxy:
    namespaced: false
    verbs:
    - create
    - delete
    - get
    - patch
    - update
  nodes/status:
    namespaced: false
    verbs:
    - get
    - patch
    - update
  persistentvolumeclaims:
    namespaced: true
    verbs:
    - create
    - delete
    - deletecollection
    - get
    - list
    - patch
    - update
    - watch
  persistentvolumeclaims/status:
    namespaced: true
    verbs:
    - get
    - patch
    - update
  persistentvolumes:
    namespaced: false
    verbs:
    - create
    - delete
    - deletecollection
    - get
    - list
    - patch
    - update
    - watch
  persistentvolumes/status:
    namespaced: false
    verbs:
    - get
    - patch
    - update
  pods:
    namespaced: true
    verbs:
    - create
    - delete
    - deletecollection
    - get
    - list
    - patch
    - update
    - watch
  pods/attach:
    namespaced: true
    verbs:
    - create
    - get
  pods/binding:
    namespaced: true
    verbs:
    - create
  pods/eviction:
    namespaced: true
    verbs:
    - create
  pods/exec:
    namespaced: true
    verbs:
    - create
    - get
  pods/log:
    namespaced: true
    verbs:
    - get
  pods/portforward:
    namespaced: true
    verbs:
    - create
    - get
  pods/proxy:
    namespaced: true
    verbs:
    - create
    - delete
    - get
    - patch
    - update
  pods/status:
    namespaced: true
    verbs:
    - get
    - patch
    - update
  podtemplates:
    namespaced: true
    verbs:
    - create
    - delete
    - deletecollection
    - get
    - list
    - patch
    - update
    - watch
  replicationcontrollers:
    namespaced: true
    verbs:
    - create
    - delete
    - deletecollection
    - get
    - list
    - patch
    - update
    - watch
  replicationcontrollers/scale:
    namespaced: true
    verbs:
    - get
    - patch
    - update
  replicationcontrollers/status:
    namespaced: true
    verbs:
    - get
    - patch
    - update
  resourcequotas:
    namespaced: true
    verbs:
    - create
    - delete
    - deletecollection
    - get
    - list
    - patch
    - update
    - watch
  resourcequotas/status:
    namespaced: true
    verbs:
    - get
    - patch
    - update
  secrets:
    namespaced: true
    verbs:
    - create
    - delete
    - deletecollection
    - get
    - list
    - patch
    - update
    - watch
  serviceaccounts:
    namespaced: true
    verbs:
    - create
    - delete
    - deletecollection
    - get
    - list
    - patch
    - update
    - watch
  services:
    namespaced: true
    verbs:
    - create
    - delete
    - get
    - list
    - patch
    - update
    - watch
  services/proxy:
    namespaced: true
    verbs:
    - create
    - delete
    - get
    - patch
    - update
  services/status:
    namespaced: true
    verbs:
    - get
    - patch
    - update
admissionregistration.k8s.io:
  mutatingwebhookconfigurations:
    namespaced: false
    verbs:
    - create
    - delete
    - deletecollection
    - get
    - list
    - patch
    - update
    - watch
  validatingwebhookconfigurations:
    namespaced: false
    verbs:
    - create
    - delete
    - deletecollection
    - get
    - list
    - patch
    - update
    - watch
apiextensions.k8s.io:
  customresourcedefinitions:
    namespaced: false
    verbs:
    - create
    - delete
    - deletecollection
    - get
    - list
    - patch
    - update
    - watch
  customresourcedefinitions/status:
    namespaced: false
    verbs:
    - get
    - patch
    - update
apiregistration.k8s.io:
  apiservices:
    namespaced: false
    verbs:
    - create
    - delete
    - deletecollection
    - get
    - list
    - patch
    - update
    - watch
  apiservices/status:
    namespaced: false
    verbs:
    - get
    - patch
    - update
apps:
  controllerrevisions:
    namespaced: true
    verbs:
    - create
    - delete
    - deletecollection
    - get
    - list
    - patch
    - update
    - watch
  daemonsets:
    namespaced: true
    verbs:
    - create
    - delete
    - deletecollection
    - get
    - list
    - patch
    - update
    - watch
  daemonsets/status:
    namespaced: true
    verbs:
    - get
    - patch
    - update
  deployments:
    namespaced: true
    verbs:
    - create
    - delete
    - deletecollection
    - get
    - list
    - patch
    - update
    - watch
  deployments/scale:
    namespaced: true
    verbs:
    - get
    - patch
    - update
  deployments/status:
    namespaced: true
    verbs:
    - get
    - patch
    - update
  replicasets:
    namespaced: true
    verbs:
    - create
    - delete
    - deletecollection
    - get
    - list
    - patch
    - update
    - watch
  replicasets/scale:
    namespaced: true
    verbs:
    - get
    - patch
    - update
  replicasets/status:
    namespaced: true
    verbs:
    - get
    - patch
    - update
  statefulsets:
    namespaced: true
    verbs:
    - create
    - delete
    - deletecollection
    - get
    - list
    - patch
    - update
    - watch
  statefulsets/scale:
    namespaced: true
    verbs:
    - get
    - patch
    - update
  statefulsets/status:
    namespaced: true
    verbs:
    - get
    - patch
    - update
apps.openshift.io:
  deploymentconfigs:
    namespaced: true
    verbs:
    - create
    - delete
    - deletecollection
    - get
    - list
    - patch
    - update
    - watch
  deploymentconfigs/instantiate:
    namespaced: true
    verbs:
    - create
  deploymentconfigs/log:
    namespaced: true
    verbs:
    - get
  deploymentconfigs/rollback:
    namespaced: true
    verbs:
    - create
  deploymentconfigs/scale:
    namespaced: true
    verbs:
    - get
    - patch
    - update
  deploymentconfigs/status:
    namespaced: true
    verbs:
    - get
    - patch
    - update
authentication.k8s.io:
  tokenreviews:
    namespaced: false
    verbs:
    - create
authorization.k8s.io:
  localsubjectaccessreviews:
    namespaced: true
    verbs:
    - create
  selfsubjectaccessreviews:
    namespaced: false
    verbs:
    - create
  selfsubjectrulesreviews:
    namespaced: false
    verbs:
    - create
  subjectaccessreviews:
    namespaced: false
    verbs:
    - create
authorization.openshift.io:
  clusterrolebindings:
    namespaced: false
    verbs:
    - create
    - delete
    - get
    - list
    - patch
    - update
  clusterroles:
    namespaced: false
    verbs:
    - create
    - delete
    - get
    - list
    - patch
    - update
  localresourceaccessreviews:
    namespaced: true
    verbs:
    - create
  localsubjectaccessreviews:
    namespaced: true
    verbs:
    - create
  resourceaccessreviews:
    namespaced: false
    verbs:
    - create
  rolebindingrestrictions:
    namespaced: true
    verbs:
    - create
    - delete
    - deletecollection
    - get
    - list
    - patch
    - update
    - watch
  rolebindings:
    namespaced: true
    verbs:
    - create
    - delete
    - get
    - list
    - patch
    - update
  roles:
    namespaced: true
    verbs:
    - create
    - delete
    - get
    - list
    - patch
    - update
  selfsubjectrulesreviews:
    namespaced: true
    verbs:
    - create
  subjectaccessreviews:
    namespaced: false
    verbs:
    - create
  subjectrulesreviews:
    namespaced: true
    verbs:
    - create
autoscaling:
  horizontalpodautoscalers:
    namespaced: true
    verbs:
    - create
    - delete
    - deletecollection
    - get
    - list
    - patch
    - update
    - watch
  horizontalpodautoscalers/status:
    namespaced: true
    verbs:
    - get
    - patch
    - update
autoscaling.openshift.io:
  clusterautoscalers:
    namespaced: false
    verbs:
    - delete
    - deletecollection
    - get
    - list
    - patch
    - create
    - update
    - watch
  clusterautoscalers/status:
    namespaced: false
    verbs:
    - get
    - patch
    - update
batch:
  jobs:
    namespaced: true
    verbs:
    - create
    - delete
    - deletecollection
    - get
    - list
    - patch
    - update
    - watch
  jobs/status:
    namespaced: true
    verbs:
    - get
    - patch
    - update
build.openshift.io:
  buildconfigs:
    namespaced: true
    verbs:
    - create
    - delete
    - deletecollection
    - get
    - list
    - patch
    - update
    - watch
  buildconfigs/instantiate:
    namespaced: true
    verbs:
    - create
  buildconfigs/instantiatebinary:
    namespaced: true
    verbs:
    - create
  buildconfigs/webhooks:
    namespaced: true
    verbs:
    - create
  builds:
    namespaced: true
    verbs:
    - create
    - delete
    - deletecollection
    - get
    - list
    - patch
    - update
    - watch
  builds/clone:
    namespaced: true
    verbs:
    - create
  builds/details:
    namespaced: true
    verbs:
    - update
  builds/log:
    namespaced: true
    verbs:
    - get
certificates.k8s.io:
  certificatesigningrequests:
    namespaced: false
    verbs:
    - create
    - delete
    - deletecollection
    - get
    - list
    - patch
    - update
    - watch
  certificatesigningrequests/approval:
    namespaced: false
    verbs:
    - update
  certificatesigningrequests/status:
    namespaced: false
    verbs:
    - get
    - patch
    - update
cloudcredential.openshift.io:
  credentialsrequests:
    namespaced: true
    verbs:
    - delete
    - deletecollection
    - get
    - list
    - patch
    - create
    - update
    - watch
  credentialsrequests/status:
    namespaced: true
    verbs:
    - get
    - patch
    - update
config.openshift.io:
  apiservers:
    namespaced: false
    verbs:
    - delete
    - deletecollection
    - get
    - list
    - patch
    - create
    - update
    - watch
  authentications:
    namespaced: false
    verbs:
    - delete
    - deletecollection
    - get
    - list
    - patch
    - create
    - update
    - watch
  authentications/status:
    namespaced: false
    verbs:
    - get
    - patch
    - update
  builds:
    namespaced: false
    verbs:
    - delete
    - deletecollection
    - get
    - list
    - patch
    - create
    - update
    - watch
  clusteroperators:
    namespaced: false
    verbs:
    - delete
    - deletecollection
    - get
    - list
    - patch
    - create
    - update
    - watch
  clusteroperators/status:
    namespaced: false
    verbs:
    - get
    - patch
    - update
  clusterversions:
    namespaced: false
    verbs:
    - delete
    - deletecollection
    - get
    - list
    - patch
    - create
    - update
    - watch
  clusterversions/status:
    namespaced: false
    verbs:
    - get
    - patch
    - update
  consoles:
    namespaced: false
    verbs:
    - delete
    - deletecollection
    - get
    - list
    - patch
    - create
    - update
    - watch
  consoles/status:
    namespaced: false
    verbs:
    - get
    - patch
    - update
  dnses:
    namespaced: false
    verbs:
    - delete
    - deletecollection
    - get
    - list
    - patch
    - create
    - update
    - watch
  featuregates:
    namespaced: false
    verbs:
    - delete
    - deletecollection
    - get
    - list
    - patch
    - create
    - update
    - watch
  images:
    namespaced: false
    verbs:
    - delete
    - deletecollection
    - get
    - list
    - patch
    - create
    - update
    - watch
  images/status:
    namespaced: false
    verbs:
    - get
    - patch
    - update
  infrastructures:
    namespaced: false
    verbs:
    - delete
    - deletecollection
    - get
    - list
    - patch
    - create
    - update
    - watch
  ingresses:
    namespaced: false
    verbs:
    - delete
    - deletecollection
    - get
    - list
    - patch
    - create
    - update
    - watch
  networks:
    namespaced: false
    verbs:
    - delete
    - deletecollection
    - get
    - list
    - patch
    - create
    - update
    - watch
  oauths:
    namespaced: false
    verbs:
    - delete
    - deletecollection
    - get
    - list
    - patch
    - create
    - update
    - watch
  oauths/status:
    namespaced: false
    verbs:
    - get
    - patch
    - update
  projects:
    namespaced: false
    verbs:
    - delete
    - deletecollection
    - get
    - list
    - patch
    - create
    - update
    - watch
  schedulers:
    namespaced: false
    verbs:
    - delete
    - deletecollection
    - get
    - list
    - patch
    - create
    - update
    - watch
coordination.k8s.io:
  leases:
    namespaced: true
    verbs:
    - create
    - delete
    - deletecollection
    - get
    - list
    - patch
    - update
    - watch
core.kubefed.k8s.io:
  kubefedclusters:
    namespaced: true
    verbs:
    - delete
    - deletecollection
    - get
    - list
    - patch
    - create
    - update
    - watch
  kubefedclusters/status:
    namespaced: true
    verbs:
    - get
    - patch
    - update
events.k8s.io:
  events:
    namespaced: true
    verbs:
    - create
    - delete
    - deletecollection
    - get
    - list
    - patch
    - update
    - watch
extensions:
  daemonsets:
    namespaced: true
    verbs:
    - create
    - delete
    - deletecollection
    - get
    - list
    - patch
    - update
    - watch
  daemonsets/status:
    namespaced: true
    verbs:
    - get
    - patch
    - update
  deployments:
    namespaced: true
    verbs:
    - create
    - delete
    - deletecollection
    - get
    - list
    - patch
    - update
    - watch
  deployments/rollback:
    namespaced: true
    verbs:
    - create
  deployments/scale:
    namespaced: true
    verbs:
    - get
    - patch
    - update
  deployments/status:
    namespaced: true
    verbs:
    - get
    - patch
    - update
  ingresses:
    namespaced: true
    verbs:
    - create
    - delete
    - deletecollection
    - get
    - list
    - patch
    - update
    - watch
  ingresses/status:
    namespaced: true
    verbs:
    - get
    - patch
    - update
  networkpolicies:
    namespaced: true
    verbs:
    - create
    - delete
    - deletecollection
    - get
    - list
    - patch
    - update
    - watch
  podsecuritypolicies:
    namespaced: false
    verbs:
    - create
    - delete
    - deletecollection
    - get
    - list
    - patch
    - update
    - watch
  replicasets:
    namespaced: true
    verbs:
    - create
    - delete
    - deletecollection
    - get
    - list
    - patch
    - update
    - watch
  replicasets/scale:
    namespaced: true
    verbs:
    - get
    - patch
    - update
  replicasets/status:
    namespaced: true
    verbs:
    - get
    - patch
    - update
  replicationcontrollers:
    namespaced: true
    verbs: []
  replicationcontrollers/scale:
    namespaced: true
    verbs:
    - get
    - patch
    - update
healthchecking.openshift.io:
  machinehealthchecks:
    namespaced: true
    verbs:
    - delete
    - deletecollection
    - get
    - list
    - patch
    - create
    - update
    - watch
image.openshift.io:
  images:
    namespaced: false
    verbs:
    - create
    - delete
    - deletecollection
    - get
    - list
    - patch
    - update
    - watch
  imagesignatures:
    namespaced: false
    verbs:
    - create
    - delete
  imagestreamimages:
    namespaced: true
    verbs:
    - get
  imagestreamimports:
    namespaced: true
    verbs:
    - create
  imagestreammappings:
    namespaced: true
    verbs:
    - create
  imagestreams:
    namespaced: true
    verbs:
    - create
    - delete
    - deletecollection
    - get
    - list
    - patch
    - update
    - watch
  imagestreams/layers:
    namespaced: true
    verbs:
    - get
  imagestreams/secrets:
    namespaced: true
    verbs:
    - get
  imagestreams/status:
    namespaced: true
    verbs:
    - get
    - patch
    - update
  imagestreamtags:
    namespaced: true
    verbs:
    - create
    - delete
    - get
    - list
    - patch
    - update
imageregistry.operator.openshift.io:
  configs:
    namespaced: false
    verbs:
    - delete
    - deletecollection
    - get
    - list
    - patch
    - create
    - update
    - watch
  configs/status:
    namespaced: false
    verbs:
    - get
    - patch
    - update
k8s.cni.cncf.io:
  network-attachment-definitions:
    namespaced: true
    verbs:
    - delete
    - deletecollection
    - get
    - list
    - patch
    - create
    - update
    - watch
machine.openshift.io:
  machines:
    namespaced: true
    verbs:
    - delete
    - deletecollection
    - get
    - list
    - patch
    - create
    - update
    - watch
  machines/status:
    namespaced: true
    verbs:
    - get
    - patch
    - update
  machinesets:
    namespaced: true
    verbs:
    - delete
    - deletecollection
    - get
    - list
    - patch
    - create
    - update
    - watch
  machinesets/scale:
    namespaced: true
    verbs:
    - get
    - patch
    - update
  machinesets/status:
    namespaced: true
    verbs:
    - get
    - patch
    - update
machineconfiguration.openshift.io:
  containerruntimeconfigs:
    namespaced: false
    verbs:
    - delete
    - deletecollection
    - get
    - list
    - patch
    - create
    - update
    - watch
  containerruntimeconfigs/status:
    namespaced: false
    verbs:
    - get
    - patch
    - update
  controllerconfigs:
    namespaced: false
    verbs:
    - delete
    - deletecollection
    - get
    - list
    - patch
    - create
    - update
    - watch
  controllerconfigs/status:
    namespaced: false
    verbs:
    - get
    - patch
    - update
  kubeletconfigs:
    namespaced: false
    verbs:
    - delete
    - deletecollection
    - get
    - list
    - patch
    - create
    - update
    - watch
  kubeletconfigs/status:
    namespaced: false
    verbs:
    - get
    - patch
    - update
  machineconfigpools:
    namespaced: false
    verbs:
    - delete
    - deletecollection
    - get
    - list
    - patch
    - create
    - update
    - watch
  machineconfigpools/status:
    namespaced: false
    verbs:
    - get
    - patch
    - update
  machineconfigs:
    namespaced: false
    verbs:
    - delete
    - deletecollection
    - get
    - list
    - patch
    - create
    - update
    - watch
  mcoconfigs:
    namespaced: true
    verbs:
    - delete
    - deletecollection
    - get
    - list
    - patch
    - create
    - update
    - watch
metrics.k8s.io:
  nodes:
    namespaced: false
    verbs:
    - get
    - list
  pods:
    namespaced: true
    verbs:
    - get
    - list
monitoring.coreos.com:
  alertmanagers:
    namespaced: true
    verbs:
    - delete
    - deletecollection
    - get
    - list
    - patch
    - create
    - update
    - watch
  prometheuses:
    namespaced: true
    verbs:
    - delete
    - deletecollection
    - get
    - list
    - patch
    - create
    - update
    - watch
  prometheusrules:
    namespaced: true
    verbs:
    - delete
    - deletecollection
    - get
    - list
    - patch
    - create
    - update
    - watch
  servicemonitors:
    namespaced: true
    verbs:
    - delete
    - deletecollection
    - get
    - list
    - patch
    - create
    - update
    - watch
network.openshift.io:
  clusternetworks:
    namespaced: false
    verbs:
    - delete
    - deletecollection
    - get
    - list
    - patch
    - create
    - update
    - watch
  egressnetworkpolicies:
    namespaced: true
    verbs:
    - delete
    - deletecollection
    - get
    - list
    - patch
    - create
    - update
    - watch
  hostsubnets:
    namespaced: false
    verbs:
    - delete
    - deletecollection
    - get
    - list
    - patch
    - create
    - update
    - watch
  netnamespaces:
    namespaced: false
    verbs:
    - delete
    - deletecollection
    - get
    - list
    - patch
    - create
    - update
    - watch
networking.k8s.io:
  networkpolicies:
    namespaced: true
    verbs:
    - create
    - delete
    - deletecollection
    - get
    - list
    - patch
    - update
    - watch
oauth.openshift.io:
  oauthaccesstokens:
    namespaced: false
    verbs:
    - create
    - delete
    - deletecollection
    - get
    - list
    - patch
    - update
    - watch
  oauthauthorizetokens:
    namespaced: false
    verbs:
    - create
    - delete
    - deletecollection
    - get
    - list
    - patch
    - update
    - watch
  oauthclientauthorizations:
    namespaced: false
    verbs:
    - create
    - delete
    - deletecollection
    - get
    - list
    - patch
    - update
    - watch
  oauthclients:
    namespaced: false
    verbs:
    - create
    - delete
    - deletecollection
    - get
    - list
    - patch
    - update
    - watch
operator.openshift.io:
  authentications:
    namespaced: false
    verbs:
    - delete
    - deletecollection
    - get
    - list
    - patch
    - create
    - update
    - watch
  authentications/status:
    namespaced: false
    verbs:
    - get
    - patch
    - update
  consoles:
    namespaced: false
    verbs:
    - delete
    - deletecollection
    - get
    - list
    - patch
    - create
    - update
    - watch
  consoles/status:
    namespaced: false
    verbs:
    - get
    - patch
    - update
  dnses:
    namespaced: false
    verbs:
    - delete
    - deletecollection
    - get
    - list
    - patch
    - create
    - update
    - watch
  dnses/status:
    namespaced: false
    verbs:
    - get
    - patch
    - update
  ingresscontrollers:
    namespaced: true
    verbs:
    - delete
    - deletecollection
    - get
    - list
    - patch
    - create
    - update
    - watch
  ingresscontrollers/scale:
    namespaced: true
    verbs:
    - get
    - patch
    - update
  ingresscontrollers/status:
    namespaced: true
    verbs:
    - get
    - patch
    - update
  kubeapiservers:
    namespaced: false
    verbs:
    - delete
    - deletecollection
    - get
    - list
    - patch
    - create
    - update
    - watch
  kubeapiservers/status:
    namespaced: false
    verbs:
    - get
    - patch
    - update
  kubecontrollermanagers:
    namespaced: false
    verbs:
    - delete
    - deletecollection
    - get
    - list
    - patch
    - create
    - update
    - watch
  kubecontrollermanagers/status:
    namespaced: false
    verbs:
    - get
    - patch
    - update
  kubeschedulers:
    namespaced: false
    verbs:
    - delete
    - deletecollection
    - get
    - list
    - patch
    - create
    - update
    - watch
  kubeschedulers/status:
    namespaced: false
    verbs:
    - get
    - patch
    - update
  networks:
    namespaced: false
    verbs:
    - delete
    - deletecollection
    - get
    - list
    - patch
    - create
    - update
    - watch
  openshiftapiservers:
    namespaced: false
    verbs:
    - delete
    - deletecollection
    - get
    - list
    - patch
    - create
    - update
    - watch
  openshiftapiservers/status:
    namespaced: false
    verbs:
    - get
    - patch
    - update
  openshiftcontrollermanagers:
    namespaced: false
    verbs:
    - delete
    - deletecollection
    - get
    - list
    - patch
    - create
    - update
    - watch
  openshiftcontrollermanagers/status:
    namespaced: false
    verbs:
    - get
    - patch
    - update
  servicecas:
    namespaced: false
    verbs:
    - delete
    - deletecollection
    - get
    - list
    - patch
    - create
    - update
    - watch
  servicecas/status:
    namespaced: false
    verbs:
    - get
    - patch
    - update
  servicecatalogapiservers:
    namespaced: false
    verbs:
    - delete
    - deletecollection
    - get
    - list
    - patch
    - create
    - update
    - watch
  servicecatalogapiservers/status:
    namespaced: false
    verbs:
    - get
    - patch
    - update
  servicecatalogcontrollermanagers:
    namespaced: false
    verbs:
    - delete
    - deletecollection
    - get
    - list
    - patch
    - create
    - update
    - watch
  servicecatalogcontrollermanagers/status:
    namespaced: false
    verbs:
    - get
    - patch
    - update
operators.coreos.com:
  catalogsourceconfigs:
    namespaced: true
    verbs:
    - delete
    - deletecollection
    - get
    - list
    - patch
    - create
    - update
    - watch
  operatorgroups:
    namespaced: true
    verbs:
    - delete
    - deletecollection
    - get
    - list
    - patch
    - create
    - update
    - watch
  operatorgroups/status:
    namespaced: true
    verbs:
    - get
    - patch
    - update
  operatorsources:
    namespaced: true
    verbs:
    - delete
    - deletecollection
    - get
    - list
    - patch
    - create
    - update
    - watch
packages.operators.coreos.com:
  packagemanifests:
    namespaced: true
    verbs:
    - get
    - list
policy:
  poddisruptionbudgets:
    namespaced: true
    verbs:
    - create
    - delete
    - deletecollection
    - get
    - list
    - patch
    - update
    - watch
  poddisruptionbudgets/status:
    namespaced: true
    verbs:
    - get
    - patch
    - update
  podsecuritypolicies:
    namespaced: false
    verbs:
    - create
    - delete
    - deletecollection
    - get
    - list
    - patch
    - update
    - watch
project.openshift.io:
  projectrequests:
    namespaced: false
    verbs:
    - create
    - list
  projects:
    namespaced: false
    verbs:
    - create
    - delete
    - get
    - list
    - patch
    - update
    - watch
quota.openshift.io:
  appliedclusterresourcequotas:
    namespaced: true
    verbs:
    - get
    - list
  clusterresourcequotas:
    namespaced: false
    verbs:
    - create
    - delete
    - deletecollection
    - get
    - list
    - patch
    - update
    - watch
  clusterresourcequotas/status:
    namespaced: false
    verbs:
    - get
    - patch
    - update
rbac.authorization.k8s.io:
  clusterrolebindings:
    namespaced: false
    verbs:
    - create
    - delete
    - deletecollection
    - get
    - list
    - patch
    - update
    - watch
  clusterroles:
    namespaced: false
    verbs:
    - create
    - delete
    - deletecollection
    - get
    - list
    - patch
    - update
    - watch
  rolebindings:
    namespaced: true
    verbs:
    - create
    - delete
    - deletecollection
    - get
    - list
    - patch
    - update
    - watch
  roles:
    namespaced: true
    verbs:
    - create
    - delete
    - deletecollection
    - get
    - list
    - patch
    - update
    - watch
route.openshift.io:
  routes:
    namespaced: true
    verbs:
    - create
    - delete
    - deletecollection
    - get
    - list
    - patch
    - update
    - watch
  routes/status:
    namespaced: true
    verbs:
    - get
    - patch
    - update
samples.operator.openshift.io:
  configs:
    namespaced: false
    verbs:
    - delete
    - deletecollection
    - get
    - list
    - patch
    - create
    - update
    - watch
  configs/status:
    namespaced: false
    verbs:
    - get
    - patch
    - update
scheduling.k8s.io:
  priorityclasses:
    namespaced: false
    verbs:
    - create
    - delete
    - deletecollection
    - get
    - list
    - patch
    - update
    - watch
security.openshift.io:
  podsecuritypolicyreviews:
    namespaced: true
    verbs:
    - create
  podsecuritypolicyselfsubjectreviews:
    namespaced: true
    verbs:
    - create
  podsecuritypolicysubjectreviews:
    namespaced: true
    verbs:
    - create
  rangeallocations:
    namespaced: false
    verbs:
    - create
    - delete
    - deletecollection
    - get
    - list
    - patch
    - update
    - watch
  securitycontextconstraints:
    namespaced: false
    verbs:
    - create
    - delete
    - deletecollection
    - get
    - list
    - patch
    - update
    - watch
storage.k8s.io:
  storageclasses:
    namespaced: false
    verbs:
    - create
    - delete
    - deletecollection
    - get
    - list
    - patch
    - update
    - watch
  volumeattachments:
    namespaced: false
    verbs:
    - create
    - delete
    - deletecollection
    - get
    - list
    - patch
    - update
    - watch
  volumeattachments/status:
    namespaced: false
    verbs:
    - get
    - patch
    - update
template.openshift.io:
  brokertemplateinstances:
    namespaced: false
    verbs:
    - create
    - delete
    - deletecollection
    - get
    - list
    - patch
    - update
    - watch
  processedtemplates:
    namespaced: true
    verbs:
    - create
  templateinstances:
    namespaced: true
    verbs:
    - create
    - delete
    - deletecollection
    - get
    - list
    - patch
    - update
    - watch
  templateinstances/status:
    namespaced: true
    verbs:
    - get
    - patch
    - update
  templates:
    namespaced: true
    verbs:
    - create
    - delete
    - deletecollection
    - get
    - list
    - patch
    - update
    - watch
tuned.openshift.io:
  tuneds:
    namespaced: true
    verbs:
    - delete
    - deletecollection
    - get
    - list
    - patch
    - create
    - update
    - watch
user.openshift.io:
  groups:
    namespaced: false
    verbs:
    - create
    - delete
    - deletecollection
    - get
    - list
    - patch
    - update
    - watch
  identities:
    namespaced: false
    verbs:
    - create
    - delete
    - deletecollection
    - get
    - list
    - patch
    - update
    - watch
  useridentitymappings:
    namespaced: false
    verbs:
    - create
    - delete
    - get
    - patch
    - update
  users:
    namespaced: false
    verbs:
    - create
    - delete
    - deletecollection
    - get
    - list
    - patch
    - update
    - watch
`+"`"+``)

func apiresources_assetGoBytes() ([]byte, error) {
	return _apiresources_assetGo, nil
}

func apiresources_assetGo() (*asset, error) {
	bytes, err := apiresources_assetGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "apiresources_asset.go", size: 63664, mode: os.FileMode(420), modTime: time.Unix(1571988577, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"apiresources.yaml":     apiresourcesYaml,
	"apiresources_asset.go": apiresources_assetGo,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"apiresources.yaml":     &bintree{apiresourcesYaml, map[string]*bintree{}},
	"apiresources_asset.go": &bintree{apiresources_assetGo, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
