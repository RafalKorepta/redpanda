// Copyright 2021 Vectorized, Inc.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.md
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0

// Package resources contains reconciliation logic for redpanda.vectorized.io CRD
package resources

import (
	"context"
	"fmt"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// Resource decompose the reconciliation loop to specific kubernetes objects
type Resource interface {
	// Obj returns resource managed client.Object
	Obj() (client.Object, error)

	// Key returns namespace/name object that is used to identify object.
	// For reference please visit types.NamespacedName docs in k8s.io/apimachinery
	Key() types.NamespacedName

	// Kind returns the canonical name of the kubernetes managed resource
	Kind() string

	// Ensure reconcile only one resource available in Kubernetes API server
	Ensure(ctx context.Context) error
}

type internalResource interface {
	Resource
	client.Reader
	client.Writer
}

func ensure(
	ctx context.Context,
	r internalResource,
	checkObj client.Object,
	resourceName string,
	l logr.Logger,
) (bool, error) {
	err := r.Get(ctx, r.Key(), checkObj)
	if err != nil && !errors.IsNotFound(err) {
		return false, fmt.Errorf("error while fetching %s resource: %w", resourceName, err)
	}

	if errors.IsNotFound(err) {
		l.Info(fmt.Sprintf("%s %s does not exist, going to create one", resourceName, r.Key().Name))

		obj, err := r.Obj()
		if err != nil {
			return false, fmt.Errorf("unable to construct %s object: %w", resourceName, err)
		}

		if err := r.Create(ctx, obj); err != nil {
			return false, fmt.Errorf("unable to create %s resource: %w", resourceName, err)
		}

		return true, nil
	}

	return false, nil
}
