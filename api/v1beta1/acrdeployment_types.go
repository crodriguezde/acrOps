/*
Copyright 2024.

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
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ACRDeploymentSpec defines the desired state of ACRDeployment
type ACRDeploymentSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	Image      AzureContainerRegistryImage `json:"image,omitempty"`
	Deployment *appsv1.Deployment          `json:"deployment,omitempty"`
}

// ACRDeploymentStatus defines the observed state of ACRDeployment
type ACRDeploymentStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// ACRDeployment is the Schema for the acrdeployments API
type ACRDeployment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ACRDeploymentSpec   `json:"spec,omitempty"`
	Status ACRDeploymentStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ACRDeploymentList contains a list of ACRDeployment
type ACRDeploymentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ACRDeployment `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ACRDeployment{}, &ACRDeploymentList{})
}
