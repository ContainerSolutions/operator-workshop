package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ExampleWorkshopSpec defines the desired state of ExampleWorkshop
// +k8s:openapi-gen=true
type ExampleWorkshopSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
	ContainerSpec  ContainerSpec  `json:"containerSpec,required"`
}

type ContainerSpec struct {
	Name 	 		  string 	`json:"name,required"`
	Image 	 		  string 	`json:"image,required"`
	Command  		  []string  `json:"command,required"`
}

// ExampleWorkshopStatus defines the observed state of ExampleWorkshop
// +k8s:openapi-gen=true
type ExampleWorkshopStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ExampleWorkshop is the Schema for the exampleworkshops API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
type ExampleWorkshop struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ExampleWorkshopSpec   `json:"spec,omitempty"`
	Status ExampleWorkshopStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ExampleWorkshopList contains a list of ExampleWorkshop
type ExampleWorkshopList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ExampleWorkshop `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ExampleWorkshop{}, &ExampleWorkshopList{})
}
