package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

func (in *ClusterStorage) DeepCopyInto(out *ClusterStorage) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}
func (in *ClusterStorage) DeepCopy() *ClusterStorage {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(ClusterStorage)
	in.DeepCopyInto(out)
	return out
}
func (in *ClusterStorage) DeepCopyObject() runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
func (in *ClusterStorageList) DeepCopyInto(out *ClusterStorageList) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterStorage, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}
func (in *ClusterStorageList) DeepCopy() *ClusterStorageList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(ClusterStorageList)
	in.DeepCopyInto(out)
	return out
}
func (in *ClusterStorageList) DeepCopyObject() runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
func (in *ClusterStorageSpec) DeepCopyInto(out *ClusterStorageSpec) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	return
}
func (in *ClusterStorageSpec) DeepCopy() *ClusterStorageSpec {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(ClusterStorageSpec)
	in.DeepCopyInto(out)
	return out
}
func (in *ClusterStorageStatus) DeepCopyInto(out *ClusterStorageStatus) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	return
}
func (in *ClusterStorageStatus) DeepCopy() *ClusterStorageStatus {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(ClusterStorageStatus)
	in.DeepCopyInto(out)
	return out
}
