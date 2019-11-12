package v1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Hello HelloType is a top level type
type Hello struct {
	//
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +optional
	Status HelloStatus `json:"status,omitempty"`

	Spec HelloSpec `json:"spec,omitempty"`
}

// HelloStatus comment
type HelloStatus struct {
	Name string
}

// HelloSpec comment
type HelloSpec struct {
	Message string `json:"message,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime/Object

// HelloList comment
type HelloList struct {
	metav1.TypeMeta `json:",inline"`

	// +optional
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []Hello `json:"items"`
}
