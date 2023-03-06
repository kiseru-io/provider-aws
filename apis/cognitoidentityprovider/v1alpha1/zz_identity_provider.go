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

// IdentityProviderParameters defines the desired state of IdentityProvider
type IdentityProviderParameters struct {
	// Region is which region the IdentityProvider will be created.
	// +kubebuilder:validation:Required
	Region string `json:"region"`
	// A mapping of IdP attributes to standard and custom user pool attributes.
	AttributeMapping map[string]*string `json:"attributeMapping,omitempty"`
	// A list of IdP identifiers.
	IDpIdentifiers []*string `json:"idpIdentifiers,omitempty"`
	// The IdP type.
	// +kubebuilder:validation:Required
	ProviderType                     *string `json:"providerType"`
	CustomIdentityProviderParameters `json:",inline"`
}

// IdentityProviderSpec defines the desired state of IdentityProvider
type IdentityProviderSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       IdentityProviderParameters `json:"forProvider"`
}

// IdentityProviderObservation defines the observed state of IdentityProvider
type IdentityProviderObservation struct {
	// The date the IdP was created.
	CreationDate *metav1.Time `json:"creationDate,omitempty"`
	// The date the IdP was last modified.
	LastModifiedDate *metav1.Time `json:"lastModifiedDate,omitempty"`
	// The IdP name.
	ProviderName *string `json:"providerName,omitempty"`
	// The user pool ID.
	UserPoolID *string `json:"userPoolID,omitempty"`
}

// IdentityProviderStatus defines the observed state of IdentityProvider.
type IdentityProviderStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          IdentityProviderObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// IdentityProvider is the Schema for the IdentityProviders API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type IdentityProvider struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IdentityProviderSpec   `json:"spec"`
	Status            IdentityProviderStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IdentityProviderList contains a list of IdentityProviders
type IdentityProviderList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IdentityProvider `json:"items"`
}

// Repository type metadata.
var (
	IdentityProviderKind             = "IdentityProvider"
	IdentityProviderGroupKind        = schema.GroupKind{Group: CRDGroup, Kind: IdentityProviderKind}.String()
	IdentityProviderKindAPIVersion   = IdentityProviderKind + "." + GroupVersion.String()
	IdentityProviderGroupVersionKind = GroupVersion.WithKind(IdentityProviderKind)
)

func init() {
	SchemeBuilder.Register(&IdentityProvider{}, &IdentityProviderList{})
}
