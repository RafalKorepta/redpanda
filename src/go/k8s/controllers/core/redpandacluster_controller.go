// Copyright 2021 Vectorized, Inc.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.md
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0

// Package core is responsible for Reconcile the core.vectorized.io Custom Resource Definition
package core

import (
	"context"

	"github.com/go-logr/logr"
	corev1alpha1 "github.com/vectorizedio/kubernetes-operator/apis/core/v1alpha1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// RedpandaClusterReconciler reconciles a RedpandaCluster object
type RedpandaClusterReconciler struct {
	client.Client
	Log	logr.Logger
	Scheme	*runtime.Scheme
}

// +kubebuilder:rbac:groups=core.vectorized.io,resources=redpandaclusters,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=core.vectorized.io,resources=redpandaclusters/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=core.vectorized.io,resources=redpandaclusters/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the RedpandaCluster object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.7.0/pkg/reconcile
func (r *RedpandaClusterReconciler) Reconcile(
	ctx context.Context, req ctrl.Request,
) (ctrl.Result, error) {
	_ = r.Log.WithValues("redpandacluster", req.NamespacedName)

	// your logic here

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *RedpandaClusterReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&corev1alpha1.RedpandaCluster{}).
		Complete(r)
}
