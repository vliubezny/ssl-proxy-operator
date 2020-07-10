package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// SslProxySpec defines the desired state of SslProxy
type SslProxySpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html

	// Name of service to proxy TCP connection to
	// +kubebuilder:validation:MinLength=1
	ServiceName string `json:"serviceName"`

	// Port of service to proxy TCP connection to
	// +kubebuilder:validation:MinLength=1
	ServicePort string `json:"servicePort"`

	// GCP zone of KGE cluster
	// +kubebuilder:validation:MinLength=1
	Zone string `json:"zone"`

	// Name of network endpoint group used as SSL Proxy backend
	// +kubebuilder:validation:MinLength=1
	NetworkEndpointGroup string `json:"networkEndpointGroup"`
}

// SslProxyStatus defines the observed state of SslProxy
type SslProxyStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html

	// List of configured endpoints
	Endpoints []string `json:"endpoints"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// SslProxy is the Schema for the sslproxies API
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=sslproxies,scope=Namespaced
type SslProxy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SslProxySpec   `json:"spec,omitempty"`
	Status SslProxyStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// SslProxyList contains a list of SslProxy
type SslProxyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SslProxy `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SslProxy{}, &SslProxyList{})
}
