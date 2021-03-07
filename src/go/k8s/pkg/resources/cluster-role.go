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
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	k8sclient "sigs.k8s.io/controller-runtime/pkg/client"
)

var _ Resource = &ClusterRoleResource{}

// ClusterRoleResource is part of the reconciliation of redpanda.vectorized.io CRD
// that gives init container ability to retrieve node external IP by RoleBinding.
type ClusterRoleResource struct {
	k8sclient.Client
	scheme       *runtime.Scheme
	pandaCluster *redpandav1alpha1.Cluster
	logger       logr.Logger
}

// NewClusterRole creates ClusterRoleResource
func NewClusterRole(
	client k8sclient.Client,
	pandaCluster *redpandav1alpha1.Cluster,
	scheme *runtime.Scheme,
	logger logr.Logger,
) *ClusterRoleResource {
	return &ClusterRoleResource{
		client,
		scheme,
		pandaCluster,
		logger.WithValues("Kind", clusterRoleKind()),
	}
}

// Ensure manages Role that is assigned to ServiceAccount used in initContainer
func (r *ClusterRoleResource) Ensure(ctx context.Context) error {
	_, err := ensure(ctx, r, &v1.ClusterRole{}, "ClusterRole", r.logger)
	return err
}

// Obj returns resource managed client.Object
func (r *ClusterRoleResource) Obj() (k8sclient.Object, error) {
	return &v1.ClusterRole{
		ObjectMeta: metav1.ObjectMeta{
			Name: r.Key().Name,
		},
		Rules: []v1.PolicyRule{
			{
				Verbs:     []string{"get"},
				APIGroups: []string{corev1.GroupName},
				Resources: []string{"nodes"},
			},
		},
	}, nil
}

// Key returns namespace/name object that is used to identify object.
// For reference please visit types.NamespacedName docs in k8s.io/apimachinery
func (r *ClusterRoleResource) Key() types.NamespacedName {
	return types.NamespacedName{Name: r.pandaCluster.Name}
}

// Kind returns v1.ClusterRole kind
func (r *ClusterRoleResource) Kind() string {
	return clusterRoleKind()
}

func clusterRoleKind() string {
	var r v1.ClusterRole
	return r.Kind
}
