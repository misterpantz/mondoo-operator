//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022 Mondoo, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Image) DeepCopyInto(out *Image) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Image.
func (in *Image) DeepCopy() *Image {
	if in == nil {
		return nil
	}
	out := new(Image)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Metrics) DeepCopyInto(out *Metrics) {
	*out = *in
	if in.ResourceLabels != nil {
		in, out := &in.ResourceLabels, &out.ResourceLabels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Metrics.
func (in *Metrics) DeepCopy() *Metrics {
	if in == nil {
		return nil
	}
	out := new(Metrics)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MondooAuditConfig) DeepCopyInto(out *MondooAuditConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MondooAuditConfig.
func (in *MondooAuditConfig) DeepCopy() *MondooAuditConfig {
	if in == nil {
		return nil
	}
	out := new(MondooAuditConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MondooAuditConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MondooAuditConfigCondition) DeepCopyInto(out *MondooAuditConfigCondition) {
	*out = *in
	in.LastUpdateTime.DeepCopyInto(&out.LastUpdateTime)
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MondooAuditConfigCondition.
func (in *MondooAuditConfigCondition) DeepCopy() *MondooAuditConfigCondition {
	if in == nil {
		return nil
	}
	out := new(MondooAuditConfigCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MondooAuditConfigData) DeepCopyInto(out *MondooAuditConfigData) {
	*out = *in
	in.Nodes.DeepCopyInto(&out.Nodes)
	in.Workloads.DeepCopyInto(&out.Workloads)
	out.Webhooks = in.Webhooks
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MondooAuditConfigData.
func (in *MondooAuditConfigData) DeepCopy() *MondooAuditConfigData {
	if in == nil {
		return nil
	}
	out := new(MondooAuditConfigData)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MondooAuditConfigList) DeepCopyInto(out *MondooAuditConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MondooAuditConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MondooAuditConfigList.
func (in *MondooAuditConfigList) DeepCopy() *MondooAuditConfigList {
	if in == nil {
		return nil
	}
	out := new(MondooAuditConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MondooAuditConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MondooAuditConfigStatus) DeepCopyInto(out *MondooAuditConfigStatus) {
	*out = *in
	if in.Pods != nil {
		in, out := &in.Pods, &out.Pods
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]MondooAuditConfigCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MondooAuditConfigStatus.
func (in *MondooAuditConfigStatus) DeepCopy() *MondooAuditConfigStatus {
	if in == nil {
		return nil
	}
	out := new(MondooAuditConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MondooOperatorConfig) DeepCopyInto(out *MondooOperatorConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MondooOperatorConfig.
func (in *MondooOperatorConfig) DeepCopy() *MondooOperatorConfig {
	if in == nil {
		return nil
	}
	out := new(MondooOperatorConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MondooOperatorConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MondooOperatorConfigCondition) DeepCopyInto(out *MondooOperatorConfigCondition) {
	*out = *in
	in.LastUpdateTime.DeepCopyInto(&out.LastUpdateTime)
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MondooOperatorConfigCondition.
func (in *MondooOperatorConfigCondition) DeepCopy() *MondooOperatorConfigCondition {
	if in == nil {
		return nil
	}
	out := new(MondooOperatorConfigCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MondooOperatorConfigList) DeepCopyInto(out *MondooOperatorConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MondooOperatorConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MondooOperatorConfigList.
func (in *MondooOperatorConfigList) DeepCopy() *MondooOperatorConfigList {
	if in == nil {
		return nil
	}
	out := new(MondooOperatorConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MondooOperatorConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MondooOperatorConfigSpec) DeepCopyInto(out *MondooOperatorConfigSpec) {
	*out = *in
	in.Metrics.DeepCopyInto(&out.Metrics)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MondooOperatorConfigSpec.
func (in *MondooOperatorConfigSpec) DeepCopy() *MondooOperatorConfigSpec {
	if in == nil {
		return nil
	}
	out := new(MondooOperatorConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MondooOperatorConfigStatus) DeepCopyInto(out *MondooOperatorConfigStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]MondooOperatorConfigCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MondooOperatorConfigStatus.
func (in *MondooOperatorConfigStatus) DeepCopy() *MondooOperatorConfigStatus {
	if in == nil {
		return nil
	}
	out := new(MondooOperatorConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Nodes) DeepCopyInto(out *Nodes) {
	*out = *in
	out.Image = in.Image
	in.Resources.DeepCopyInto(&out.Resources)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Nodes.
func (in *Nodes) DeepCopy() *Nodes {
	if in == nil {
		return nil
	}
	out := new(Nodes)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebhookCertificateConfig) DeepCopyInto(out *WebhookCertificateConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebhookCertificateConfig.
func (in *WebhookCertificateConfig) DeepCopy() *WebhookCertificateConfig {
	if in == nil {
		return nil
	}
	out := new(WebhookCertificateConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Webhooks) DeepCopyInto(out *Webhooks) {
	*out = *in
	out.CertificateConfig = in.CertificateConfig
	out.Image = in.Image
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Webhooks.
func (in *Webhooks) DeepCopy() *Webhooks {
	if in == nil {
		return nil
	}
	out := new(Webhooks)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Workloads) DeepCopyInto(out *Workloads) {
	*out = *in
	in.Resources.DeepCopyInto(&out.Resources)
	out.Image = in.Image
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Workloads.
func (in *Workloads) DeepCopy() *Workloads {
	if in == nil {
		return nil
	}
	out := new(Workloads)
	in.DeepCopyInto(out)
	return out
}
