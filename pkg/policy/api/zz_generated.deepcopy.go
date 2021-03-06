// +build !ignore_autogenerated

// Copyright 2017 Authors of Cilium
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package api

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	reflect "reflect"
)

// GetGeneratedDeepCopyFuncs returns the generated funcs, since we aren't registering them.
//
// Deprecated: deepcopy registration will go away when static deepcopy is fully implemented.
func GetGeneratedDeepCopyFuncs() []conversion.GeneratedDeepCopyFunc {
	return []conversion.GeneratedDeepCopyFunc{
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*CIDRRule).DeepCopyInto(out.(*CIDRRule))
			return nil
		}, InType: reflect.TypeOf(&CIDRRule{})},
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*EgressRule).DeepCopyInto(out.(*EgressRule))
			return nil
		}, InType: reflect.TypeOf(&EgressRule{})},
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*EndpointSelector).DeepCopyInto(out.(*EndpointSelector))
			return nil
		}, InType: reflect.TypeOf(&EndpointSelector{})},
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*IngressRule).DeepCopyInto(out.(*IngressRule))
			return nil
		}, InType: reflect.TypeOf(&IngressRule{})},
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*K8sServiceNamespace).DeepCopyInto(out.(*K8sServiceNamespace))
			return nil
		}, InType: reflect.TypeOf(&K8sServiceNamespace{})},
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*K8sServiceSelectorNamespace).DeepCopyInto(out.(*K8sServiceSelectorNamespace))
			return nil
		}, InType: reflect.TypeOf(&K8sServiceSelectorNamespace{})},
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*L7Rules).DeepCopyInto(out.(*L7Rules))
			return nil
		}, InType: reflect.TypeOf(&L7Rules{})},
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*PortProtocol).DeepCopyInto(out.(*PortProtocol))
			return nil
		}, InType: reflect.TypeOf(&PortProtocol{})},
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*PortRule).DeepCopyInto(out.(*PortRule))
			return nil
		}, InType: reflect.TypeOf(&PortRule{})},
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*PortRuleHTTP).DeepCopyInto(out.(*PortRuleHTTP))
			return nil
		}, InType: reflect.TypeOf(&PortRuleHTTP{})},
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*PortRuleKafka).DeepCopyInto(out.(*PortRuleKafka))
			return nil
		}, InType: reflect.TypeOf(&PortRuleKafka{})},
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*Rule).DeepCopyInto(out.(*Rule))
			return nil
		}, InType: reflect.TypeOf(&Rule{})},
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*Service).DeepCopyInto(out.(*Service))
			return nil
		}, InType: reflect.TypeOf(&Service{})},
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ServiceSelector).DeepCopyInto(out.(*ServiceSelector))
			return nil
		}, InType: reflect.TypeOf(&ServiceSelector{})},
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CIDRRule) DeepCopyInto(out *CIDRRule) {
	*out = *in
	if in.ExceptCIDRs != nil {
		in, out := &in.ExceptCIDRs, &out.ExceptCIDRs
		*out = make([]CIDR, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CIDRRule.
func (in *CIDRRule) DeepCopy() *CIDRRule {
	if in == nil {
		return nil
	}
	out := new(CIDRRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EgressRule) DeepCopyInto(out *EgressRule) {
	*out = *in
	if in.ToPorts != nil {
		in, out := &in.ToPorts, &out.ToPorts
		*out = make([]PortRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ToCIDR != nil {
		in, out := &in.ToCIDR, &out.ToCIDR
		*out = make([]CIDR, len(*in))
		copy(*out, *in)
	}
	if in.ToCIDRSet != nil {
		in, out := &in.ToCIDRSet, &out.ToCIDRSet
		*out = make([]CIDRRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ToEntities != nil {
		in, out := &in.ToEntities, &out.ToEntities
		*out = make([]Entity, len(*in))
		copy(*out, *in)
	}
	if in.ToServices != nil {
		in, out := &in.ToServices, &out.ToServices
		*out = make([]Service, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EgressRule.
func (in *EgressRule) DeepCopy() *EgressRule {
	if in == nil {
		return nil
	}
	out := new(EgressRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EndpointSelector) DeepCopyInto(out *EndpointSelector) {
	*out = *in
	if in.LabelSelector != nil {
		in, out := &in.LabelSelector, &out.LabelSelector
		if *in == nil {
			*out = nil
		} else {
			*out = new(v1.LabelSelector)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EndpointSelector.
func (in *EndpointSelector) DeepCopy() *EndpointSelector {
	if in == nil {
		return nil
	}
	out := new(EndpointSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IngressRule) DeepCopyInto(out *IngressRule) {
	*out = *in
	if in.FromEndpoints != nil {
		in, out := &in.FromEndpoints, &out.FromEndpoints
		*out = make([]EndpointSelector, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.FromRequires != nil {
		in, out := &in.FromRequires, &out.FromRequires
		*out = make([]EndpointSelector, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ToPorts != nil {
		in, out := &in.ToPorts, &out.ToPorts
		*out = make([]PortRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.FromCIDR != nil {
		in, out := &in.FromCIDR, &out.FromCIDR
		*out = make([]CIDR, len(*in))
		copy(*out, *in)
	}
	if in.FromCIDRSet != nil {
		in, out := &in.FromCIDRSet, &out.FromCIDRSet
		*out = make([]CIDRRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.FromEntities != nil {
		in, out := &in.FromEntities, &out.FromEntities
		*out = make([]Entity, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IngressRule.
func (in *IngressRule) DeepCopy() *IngressRule {
	if in == nil {
		return nil
	}
	out := new(IngressRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *K8sServiceNamespace) DeepCopyInto(out *K8sServiceNamespace) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new K8sServiceNamespace.
func (in *K8sServiceNamespace) DeepCopy() *K8sServiceNamespace {
	if in == nil {
		return nil
	}
	out := new(K8sServiceNamespace)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *K8sServiceSelectorNamespace) DeepCopyInto(out *K8sServiceSelectorNamespace) {
	*out = *in
	in.Selector.DeepCopyInto(&out.Selector)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new K8sServiceSelectorNamespace.
func (in *K8sServiceSelectorNamespace) DeepCopy() *K8sServiceSelectorNamespace {
	if in == nil {
		return nil
	}
	out := new(K8sServiceSelectorNamespace)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *L7Rules) DeepCopyInto(out *L7Rules) {
	*out = *in
	if in.HTTP != nil {
		in, out := &in.HTTP, &out.HTTP
		*out = make([]PortRuleHTTP, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Kafka != nil {
		in, out := &in.Kafka, &out.Kafka
		*out = make([]PortRuleKafka, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new L7Rules.
func (in *L7Rules) DeepCopy() *L7Rules {
	if in == nil {
		return nil
	}
	out := new(L7Rules)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PortProtocol) DeepCopyInto(out *PortProtocol) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PortProtocol.
func (in *PortProtocol) DeepCopy() *PortProtocol {
	if in == nil {
		return nil
	}
	out := new(PortProtocol)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PortRule) DeepCopyInto(out *PortRule) {
	*out = *in
	if in.Ports != nil {
		in, out := &in.Ports, &out.Ports
		*out = make([]PortProtocol, len(*in))
		copy(*out, *in)
	}
	if in.Rules != nil {
		in, out := &in.Rules, &out.Rules
		if *in == nil {
			*out = nil
		} else {
			*out = new(L7Rules)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PortRule.
func (in *PortRule) DeepCopy() *PortRule {
	if in == nil {
		return nil
	}
	out := new(PortRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PortRuleHTTP) DeepCopyInto(out *PortRuleHTTP) {
	*out = *in
	if in.Headers != nil {
		in, out := &in.Headers, &out.Headers
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PortRuleHTTP.
func (in *PortRuleHTTP) DeepCopy() *PortRuleHTTP {
	if in == nil {
		return nil
	}
	out := new(PortRuleHTTP)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PortRuleKafka) DeepCopyInto(out *PortRuleKafka) {
	*out = *in
	if in.apiKeyInt != nil {
		in, out := &in.apiKeyInt, &out.apiKeyInt
		if *in == nil {
			*out = nil
		} else {
			*out = new(int16)
			**out = **in
		}
	}
	if in.apiVersionInt != nil {
		in, out := &in.apiVersionInt, &out.apiVersionInt
		if *in == nil {
			*out = nil
		} else {
			*out = new(int16)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PortRuleKafka.
func (in *PortRuleKafka) DeepCopy() *PortRuleKafka {
	if in == nil {
		return nil
	}
	out := new(PortRuleKafka)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Rule) DeepCopyInto(out *Rule) {
	*out = *in
	in.EndpointSelector.DeepCopyInto(&out.EndpointSelector)
	if in.Ingress != nil {
		in, out := &in.Ingress, &out.Ingress
		*out = make([]IngressRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Egress != nil {
		in, out := &in.Egress, &out.Egress
		*out = make([]EgressRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Labels != nil {
		out.Labels = in.Labels.DeepCopy()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Rule.
func (in *Rule) DeepCopy() *Rule {
	if in == nil {
		return nil
	}
	out := new(Rule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Service) DeepCopyInto(out *Service) {
	*out = *in
	if in.K8sServiceSelector != nil {
		in, out := &in.K8sServiceSelector, &out.K8sServiceSelector
		if *in == nil {
			*out = nil
		} else {
			*out = new(K8sServiceSelectorNamespace)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.K8sService != nil {
		in, out := &in.K8sService, &out.K8sService
		if *in == nil {
			*out = nil
		} else {
			*out = new(K8sServiceNamespace)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Service.
func (in *Service) DeepCopy() *Service {
	if in == nil {
		return nil
	}
	out := new(Service)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceSelector) DeepCopyInto(out *ServiceSelector) {
	*out = *in
	if in.LabelSelector != nil {
		in, out := &in.LabelSelector, &out.LabelSelector
		if *in == nil {
			*out = nil
		} else {
			*out = new(v1.LabelSelector)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceSelector.
func (in *ServiceSelector) DeepCopy() *ServiceSelector {
	if in == nil {
		return nil
	}
	out := new(ServiceSelector)
	in.DeepCopyInto(out)
	return out
}
