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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// GoddessMomentSpec defines the desired state of GoddessMoment
type GoddessMomentSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of GoddessMoment. Edit goddessmoment_types.go to remove/update
	FoodDemand []FoodDemand `json:"foodDemand,omitempty"`
}

type FoodDemand struct {
	Name string `json:"name,omitempty"`
}

// GoddessMomentStatus defines the observed state of GoddessMoment
type GoddessMomentStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	FoodDemand []FoodDemandStatus `json:"foodDemand,omitempty"`
}

type FoodDemandStatus struct {
	Name        string      `json:"name,omitempty"`
	Status      FoodStatus  `json:"status,omitempty"`
	ClaimTime   metav1.Time `json:"claimTime,omitempty"`
	ArrivalTime metav1.Time `json:"arrivalTime,omitempty"`
	ClaimBy     string      `json:"claimBy,omitempty"`
}

type FoodStatus string

const (
	//FoodStatusPending .
	FoodStatusPending = "Pending"
	//FoodStatusPendingArrival .
	FoodStatusPendingArrival = "PendingArrival"
	//FoodStatusArrived .
	FoodStatusArrived = "Arrived"
)

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// GoddessMoment is the Schema for the goddessmoments API
type GoddessMoment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GoddessMomentSpec   `json:"spec,omitempty"`
	Status GoddessMomentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GoddessMomentList contains a list of GoddessMoment
type GoddessMomentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GoddessMoment `json:"items"`
}

func init() {
	SchemeBuilder.Register(&GoddessMoment{}, &GoddessMomentList{})
}
