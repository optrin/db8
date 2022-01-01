package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// MariaDBArchitecture configures the database replication.
type MariaDBArchitecture string

const (
	// MariaDBStandalone defines an architecture where a single primary replica is deployed.
	MariaDBStandalone MariaDBArchitecture = "Standalone"
)

// IMPORTANT: Any new fields you add must have json tags for the fields to be serialized.

// MariaDBReplica defines the configuration of a replica.
type MariaDBReplica struct {
	// Persistence configures the settings for persistent data storage.
	Persistence PersistenceConfig `json:"persistence,omitempty"`

	// Resources defines the replica's resource requests and limits.
	Resources corev1.ResourceRequirements `json:"resources,omitempty"`
}

// MariaDBSpec defines the desired state of MariaDB.
type MariaDBSpec struct {
	// IMPORTANT: Run "make" to regenerate code after modifying this file.

	// Architecture configures the database replication architecture.
	// +kubebuilder:default:=Standalone
	Architecture MariaDBArchitecture `json:"architecture,omitempty"`

	// Authentication contains all authentication settings.
	Authentication AuthenticationConfig `json:"authentication,omitempty"`

	// Image contains all configuration related to the container image of the Pod.
	Image ImageConfig `json:"image,omitempty"`

	// Primary contains the configuration specific to the primary replica.
	Primary MariaDBReplica `json:"primary,omitempty"`
}

// MariaDBStatus defines the observed state of MariaDB.
type MariaDBStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// MariaDB is the Schema for the mariadbs API
type MariaDB struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MariaDBSpec   `json:"spec,omitempty"`
	Status MariaDBStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// MariaDBList contains a list of MariaDB
type MariaDBList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MariaDB `json:"items"`
}

func init() {
	SchemeBuilder.Register(&MariaDB{}, &MariaDBList{})
}
