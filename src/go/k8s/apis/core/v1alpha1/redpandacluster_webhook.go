// Copyright 2021 Vectorized, Inc.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.md
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

// log is for logging in this package.
var redpandaclusterlog = logf.Log.WithName("redpandacluster-resource")

func (r *RedpandaCluster) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!

// +kubebuilder:webhook:path=/mutate-core-vectorized-io-v1alpha1-redpandacluster,mutating=true,failurePolicy=fail,sideEffects=None,groups=core.vectorized.io,resources=redpandaclusters,verbs=create;update,versions=v1alpha1,name=mredpandacluster.kb.io,admissionReviewVersions={v1,v1beta1}

var _ webhook.Defaulter = &RedpandaCluster{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (r *RedpandaCluster) Default() {
	redpandaclusterlog.Info("default", "name", r.Name)

	// TODO(user): fill in your defaulting logic.
}

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
// +kubebuilder:webhook:path=/validate-core-vectorized-io-v1alpha1-redpandacluster,mutating=false,failurePolicy=fail,sideEffects=None,groups=core.vectorized.io,resources=redpandaclusters,verbs=create;update,versions=v1alpha1,name=vredpandacluster.kb.io,admissionReviewVersions={v1,v1beta1}

var _ webhook.Validator = &RedpandaCluster{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *RedpandaCluster) ValidateCreate() error {
	redpandaclusterlog.Info("validate create", "name", r.Name)

	// TODO(user): fill in your validation logic upon object creation.
	return nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *RedpandaCluster) ValidateUpdate(old runtime.Object) error {
	redpandaclusterlog.Info("validate update", "name", r.Name)

	// TODO(user): fill in your validation logic upon object update.
	return nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *RedpandaCluster) ValidateDelete() error {
	redpandaclusterlog.Info("validate delete", "name", r.Name)

	// TODO(user): fill in your validation logic upon object deletion.
	return nil
}
