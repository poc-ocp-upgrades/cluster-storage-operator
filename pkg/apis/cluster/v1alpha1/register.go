package v1alpha1

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/runtime/scheme"
)

var (
	SchemeGroupVersion	= schema.GroupVersion{Group: "cluster.storage.openshift.io", Version: "v1alpha1"}
	SchemeBuilder		= &scheme.Builder{GroupVersion: SchemeGroupVersion}
)
