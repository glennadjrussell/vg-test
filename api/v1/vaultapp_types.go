/*


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

// VaultappSpec defines the desired state of Vaultapp
type VaultappSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of Vaultapp. Edit Vaultapp_types.go to remove/update

	// Name is a human readable name for the app
	Name string `json:"name,omitempty"`

	// AuthBackend is the Vault backend
	AuthBackend AuthBackend `json:"auth"`
}

// VaultappStatus defines the observed state of Vaultapp
type VaultappStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true

// Vaultapp is the Schema for the vaultapps API
type Vaultapp struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VaultappSpec   `json:"spec,omitempty"`
	Status VaultappStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VaultappList contains a list of Vaultapp
type VaultappList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Vaultapp `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Vaultapp{}, &VaultappList{})
}
