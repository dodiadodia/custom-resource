package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MyResource describes a MyResource resource
type MyResource struct {
	meta_v1.TypeMeta `json:",inline"`
	meta_v1.ObjectMeta `json:"metadata,omitempty"`

	Spec MyResourceSpec `json:"spec"`
}

type MyResourceSpec struct {
	Message	string `json:"message"`
	SomeValue *int32 `json:"someValue"`
}

type MyResourceList struct {
	meta_v1.TypeMeta `json:",inline"`
	meta_v1.ListMeta `json:"metadata"`
	
	Items []MyResource `json:"items"`
}

