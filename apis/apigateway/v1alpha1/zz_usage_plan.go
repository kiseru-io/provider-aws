/*
Copyright 2021 The Crossplane Authors.

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

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// UsagePlanParameters defines the desired state of UsagePlan
type UsagePlanParameters struct {
	// Region is which region the UsagePlan will be created.
	// +kubebuilder:validation:Required
	Region string `json:"region"`
	// The description of the usage plan.
	Description *string `json:"description,omitempty"`
	// The name of the usage plan.
	// +kubebuilder:validation:Required
	Name *string `json:"name"`
	// The quota of the usage plan.
	Quota *QuotaSettings `json:"quota,omitempty"`
	// The key-value map of strings. The valid character set is [a-zA-Z+-=._:/].
	// The tag key can be up to 128 characters and must not start with aws:. The
	// tag value can be up to 256 characters.
	Tags map[string]*string `json:"tags,omitempty"`
	// The throttling limits of the usage plan.
	Throttle                  *ThrottleSettings `json:"throttle,omitempty"`
	CustomUsagePlanParameters `json:",inline"`
}

// UsagePlanSpec defines the desired state of UsagePlan
type UsagePlanSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       UsagePlanParameters `json:"forProvider"`
}

// UsagePlanObservation defines the observed state of UsagePlan
type UsagePlanObservation struct {
	// The associated API stages of a usage plan.
	APIStages []*APIStage `json:"apiStages,omitempty"`
	// The identifier of a UsagePlan resource.
	ID *string `json:"id,omitempty"`
	// The AWS Markeplace product identifier to associate with the usage plan as
	// a SaaS product on AWS Marketplace.
	ProductCode *string `json:"productCode,omitempty"`
}

// UsagePlanStatus defines the observed state of UsagePlan.
type UsagePlanStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          UsagePlanObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// UsagePlan is the Schema for the UsagePlans API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type UsagePlan struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              UsagePlanSpec   `json:"spec"`
	Status            UsagePlanStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// UsagePlanList contains a list of UsagePlans
type UsagePlanList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []UsagePlan `json:"items"`
}

// Repository type metadata.
var (
	UsagePlanKind             = "UsagePlan"
	UsagePlanGroupKind        = schema.GroupKind{Group: CRDGroup, Kind: UsagePlanKind}.String()
	UsagePlanKindAPIVersion   = UsagePlanKind + "." + GroupVersion.String()
	UsagePlanGroupVersionKind = GroupVersion.WithKind(UsagePlanKind)
)

func init() {
	SchemeBuilder.Register(&UsagePlan{}, &UsagePlanList{})
}
