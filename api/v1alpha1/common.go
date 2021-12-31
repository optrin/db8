package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
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
	PersistentVolumeClaimName string `json:"persistentVolumeClaimName,omitempty"`

	// TODO: Continue implementation.
	// primary.persistence.storageClass	MariaDB primary persistent volume storage Class	""
	// primary.persistence.annotations	MariaDB primary persistent volume claim annotations	{}
	// primary.persistence.accessModes	MariaDB primary persistent volume access Modes	["ReadWriteOnce"]
	// primary.persistence.size	MariaDB primary persistent volume size	8Gi
	// primary.persistence.selector
}
