// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	v1 "k8s.io/api/core/v1"
	corev1 "k8s.io/client-go/applyconfigurations/core/v1"
)

// HeadGroupSpecApplyConfiguration represents an declarative configuration of the HeadGroupSpec type for use
// with apply.
type HeadGroupSpecApplyConfiguration struct {
	ServiceType    *v1.ServiceType                           `json:"serviceType,omitempty"`
	HeadService    *v1.Service                               `json:"headService,omitempty"`
	EnableIngress  *bool                                     `json:"enableIngress,omitempty"`
	RayStartParams map[string]string                         `json:"rayStartParams,omitempty"`
	Template       *corev1.PodTemplateSpecApplyConfiguration `json:"template,omitempty"`
}

// HeadGroupSpecApplyConfiguration constructs an declarative configuration of the HeadGroupSpec type for use with
// apply.
func HeadGroupSpec() *HeadGroupSpecApplyConfiguration {
	return &HeadGroupSpecApplyConfiguration{}
}

// WithServiceType sets the ServiceType field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ServiceType field is set to the value of the last call.
func (b *HeadGroupSpecApplyConfiguration) WithServiceType(value v1.ServiceType) *HeadGroupSpecApplyConfiguration {
	b.ServiceType = &value
	return b
}

// WithHeadService sets the HeadService field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the HeadService field is set to the value of the last call.
func (b *HeadGroupSpecApplyConfiguration) WithHeadService(value v1.Service) *HeadGroupSpecApplyConfiguration {
	b.HeadService = &value
	return b
}

// WithEnableIngress sets the EnableIngress field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the EnableIngress field is set to the value of the last call.
func (b *HeadGroupSpecApplyConfiguration) WithEnableIngress(value bool) *HeadGroupSpecApplyConfiguration {
	b.EnableIngress = &value
	return b
}

// WithRayStartParams puts the entries into the RayStartParams field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the RayStartParams field,
// overwriting an existing map entries in RayStartParams field with the same key.
func (b *HeadGroupSpecApplyConfiguration) WithRayStartParams(entries map[string]string) *HeadGroupSpecApplyConfiguration {
	if b.RayStartParams == nil && len(entries) > 0 {
		b.RayStartParams = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.RayStartParams[k] = v
	}
	return b
}

// WithTemplate sets the Template field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Template field is set to the value of the last call.
func (b *HeadGroupSpecApplyConfiguration) WithTemplate(value *corev1.PodTemplateSpecApplyConfiguration) *HeadGroupSpecApplyConfiguration {
	b.Template = value
	return b
}