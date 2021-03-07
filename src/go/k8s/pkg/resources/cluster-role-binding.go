// Copyright 2021 Vectorized, Inc.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.md
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0

package resources

import (
	"context"

	"github.com/go-logr/logr"
	redpandav1alpha1 "github.com/vectorizedio/redpanda/src/go/k8s/apis/redpanda/v1alpha1"
	v1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	k8sclient "sigs.k8s.io/controller-runtime/pkg/client"
)

var _ Resource = &ClusterRoleBindingResource{}

// ClusterRoleBindingResource is part of the reconciliation of redpanda.vectorized.io CRD
// that gives init container ability to retrieve node external IP by RoleBinding.
type ClusterRoleBindingResource struct {
	k8sclient.Client
	scheme       *runtime.Scheme
	pandaCluster *redpandav1alpha1.Cluster
	logger       logr.Logger
}

// NewClusterRoleBinding creates ClusterRoleBindingResource
func NewClusterRoleBinding(
	client k8sclient.Client,
	pandaCluster *redpandav1alpha1.Cluster,
	scheme *runtime.Scheme,
	logger logr.Logger,
) *ClusterRoleBindingResource {
	return &ClusterRoleBindingResource{
		client,
		scheme,
		pandaCluster,
		logger.WithValues("Kind", clusterRoleBindingKind()),
	}
}

// Ensure manages v1.ClusterRoleBinding that is assigned to v1.ServiceAccount used in initContainer
func (r *ClusterRoleBindingResource) Ensure(ctx context.Context) error {
	if !r.pandaCluster.Spec.ExternalConnectivity {
		return nil
	}

	_, err := ensure(ctx, r, &v1.ClusterRoleBinding{}, "ClusterRoleBinding", r.logger)
	return err
}

// Obj returns resource managed client.Object
// The cluster.redpanda.vectorized.io custom resource is namespaced resource, that's
// why v1.ClusterRoleBinding can not have assigned controller reference.
func (r *ClusterRoleBindingResource) Obj() (k8sclient.Object, error) {
	role := &ClusterRoleResource{pandaCluster: r.pandaCluster}
	sa := &ServiceAccountResource{pandaCluster: r.pandaCluster}

	return &v1.ClusterRoleBinding{
		// metav1.ObjectMeta can NOT have namespace set as
		// ClusterRoleBinding is the cluster wide resource.
		ObjectMeta: metav1.ObjectMeta{
			Name: r.Key().Name,
		},
		Subjects: []v1.Subject{
			{
				Kind:      "ServiceAccount",
				Name:      sa.Key().Name,
				Namespace: sa.Key().Namespace,
			},
		},
		RoleRef: v1.RoleRef{
			APIGroup: v1.GroupName,
			Kind:     "ClusterRole",
			Name:     role.Key().Name,
		},
	}, nil
}

// Key returns namespace/name object that is used to identify object.
// For reference please visit types.NamespacedName docs in k8s.io/apimachinery
func (r *ClusterRoleBindingResource) Key() types.NamespacedName {
	return types.NamespacedName{Name: r.pandaCluster.Name}
}

// Kind returns v1.ClusterRoleBinding kind
func (r *ClusterRoleBindingResource) Kind() string {
	return clusterRoleBindingKind()
}

func clusterRoleBindingKind() string {
	var r v1.ClusterRoleBinding
	return r.Kind
}
