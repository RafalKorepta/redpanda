/*
Copyright 2021 vectorized.io.

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

// Package v1alpha1 represent Custom Resource definition of the vectorized.io core group
package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// RedpandaClusterSpec defines the desired state of RedpandaCluster
type RedpandaClusterSpec struct {
	// Important: Run "make" to regenerate code after modifying this file

	// Version of Redpanda to use from upstream.
	Version	string	`json:"version,omitempty"`
	// Replicas determine how big the cluster will be.
	// +kubebuilder:validation:Minimum=0
	Replicas	*int32	`json:"replicas,omitempty"`
	// Configuration represent redpanda specific configuration
	Configuration	RedpandaConfig	`json:"configuration,omitempty"`
}

// RedpandaClusterStatus defines the observed state of RedpandaCluster
type RedpandaClusterStatus struct {
	// Important: Run "make" to regenerate code after modifying this file

	// Replicas show how many nodes are working in the cluster
	// +optional
	Replicas	int32	`json:"replicas,omitempty"`
	// Nodes of the provision redpanda nodes
	// +optional
	Nodes	[]string	`json:"nodes,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// RedpandaCluster is the Schema for the redpandaclusters API
type RedpandaCluster struct {
	metav1.TypeMeta		`json:",inline"`
	metav1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	RedpandaClusterSpec	`json:"spec,omitempty"`
	Status	RedpandaClusterStatus	`json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RedpandaClusterList contains a list of RedpandaCluster
type RedpandaClusterList struct {
	metav1.TypeMeta	`json:",inline"`
	metav1.ListMeta	`json:"metadata,omitempty"`
	Items		[]RedpandaCluster	`json:"items"`
}

// RedpandaConfig is the definition of the main configuration
type RedpandaConfig struct {
	RPCServer		SocketAddress	`json:"rpcServer,omitempty"`
	AdvertisedRPCAPI	SocketAddress	`json:"advertisedRpcApi,omitempty"`
	KafkaAPI		SocketAddress	`json:"kafkaApi,omitempty"`
	AdvertisedKafkaAPI	SocketAddress	`json:"advertisedKafkaApi,omitempty"`
	KafkaAPITLS		ServerTLS	`json:"kafkaApiTLS,omitempty"`
	AdminAPI		SocketAddress	`json:"admin,omitempty"`
	DeveloperMode		bool		`json:"developerMode,omitempty"`
}

// SocketAddress provide the way to configure the port
type SocketAddress struct {
	Port int `json:"port,omitempty"`
}

// ServerTLS allows the redpanda to configure TLS connection
type ServerTLS struct {
	KeyFile		string	`json:"keyFile,omitempty"`
	CertFile	string	`json:"certFile,omitempty"`
	TruststoreFile	string	`json:"truststoreFile,omitempty"`
	Enabled		bool	`json:"enabled,omitempty"`
}

func init() {
	SchemeBuilder.Register(&RedpandaCluster{}, &RedpandaClusterList{})
}
