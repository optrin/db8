package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// AuthenticationConfig configures the configuration of a database instance.
type AuthenticationConfig struct {
	// The name of the secret containing the `root` user's
	// credentials. If the Secret does not exist, it will
	// be created, where the name of the Secret will be
	// formatted as `db8-mariadb-<instancename>-creds`.
	SecretName string `json:"secretName,omitempty"`

	// Set to `true` to prevent operator from auto-generating
	// credentials if the supplied credentials are empty. Not
	// recommended for production instances.
	// +kubebuilder:default:=docker.io
	AllowEmptyPasswords string `json:"allowEmptyPasswords,omitempty"`
}

// ImageConfig contains all configuration related to the Pod's container image.
type ImageConfig struct {
	// +kubebuilder:default:=docker.io
	Registry string `json:"registry,omitempty"`

	// Make this field optional to make this type reusable.
	// The default will be set by the controller.
	// +kubebuilder:validation:Optional
	Repository string `json:"repository,omitempty"`

	// +kubebuilder:default:=latest
	Tag string `json:"tag,omitempty"`

	// +kubebuilder:default:=IfNotPresent
	PullPolicy corev1.PullPolicy `json:"pullPolicy,omitempty"`

	// +kubebuilder:validation:Optional
	PullSecretNames []string `json:"pullSecretNames,omitempty"`
}

// PersistenceConfig contains all configuration related to the Pod's persistent storage.
type PersistenceConfig struct {
	// If set to `false` EmptyDir will be used
	// as storage backend causing data NOT to
	// be persisted between Pod restarts.
	// +kubebuilder:default:=true
	Enabled bool `json:"enabled,omitempty"`

	// The name of the PersistentVolumeClaim used to
	// persist the data between Pod restarts. If the
	// Secret does not exist, it will be created, where
	// the name of the Secret will be formatted as
	// `db8-mariadb-<instancename>-primary-<ordinal>-data`.
	// +kubebuilder:validation:Optional
	PersistentVolumeClaimName string `json:"persistentVolumeClaimName,omitempty"`

	// The name of the StorageClass used to provision
	// the volume. If not specified, the default storage
	// class will be used.
	// +kubebuilder:validation:Optional
	StorageClassName string `json:"storageClassName,omitempty"`

	// Annotations to add to the persistent volume claim.
	// +kubebuilder:validation:Optional
	Annotations map[string]string `json:"annotations,omitempty"`

	// Access modes for the persistent volume.
	// +kubebuilder:validation:Optional
	AccessModes []corev1.PersistentVolumeAccessMode `json:"accessModes,omitempty"`

	// The size of the persistent volume.
	Size resource.Quantity `json:"size,omitempty"`

	// A selector to match an existing persistent volume.
	// +kubebuilder:validation:Optional
	Selector metav1.LabelSelector `json:"selector,omitempty"`
}
