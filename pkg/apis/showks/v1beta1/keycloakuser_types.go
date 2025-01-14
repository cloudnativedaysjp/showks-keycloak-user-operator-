/*
Copyright 2019 TAKAISHI Ryo.

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

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// KeyCloakUserSpec defines the desired state of KeyCloakUser
type KeyCloakUserSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	UserName           string `json:"username,omitempty"`
	Realm              string `json:"realm"`
	PasswordSecretName string `json:"passwordSecretName,omitempty"`
}

// KeyCloakUserStatus defines the observed state of KeyCloakUser
type KeyCloakUserStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	ID string
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// KeyCloakUser is the Schema for the keycloakusers API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
type KeyCloakUser struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   KeyCloakUserSpec   `json:"spec,omitempty"`
	Status KeyCloakUserStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// KeyCloakUserList contains a list of KeyCloakUser
type KeyCloakUserList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KeyCloakUser `json:"items"`
}

func init() {
	SchemeBuilder.Register(&KeyCloakUser{}, &KeyCloakUserList{})
}
