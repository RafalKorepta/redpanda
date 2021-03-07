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
	"errors"
	"fmt"
	"reflect"

	"github.com/go-logr/logr"
	redpandav1alpha1 "github.com/vectorizedio/redpanda/src/go/k8s/apis/redpanda/v1alpha1"
	"github.com/vectorizedio/redpanda/src/go/k8s/pkg/labels"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	k8sclient "sigs.k8s.io/controller-runtime/pkg/client"
)

var _ Resource = &ClusterResource{}

var errNodePortMissing = errors.New("the node port missing from the service")

// ClusterResource represents v1alpha1.Cluster custom resource
type ClusterResource struct {
	k8sclient.Client
	scheme       *runtime.Scheme
	pandaCluster *redpandav1alpha1.Cluster
	serviceFQDN  string
	nodePortName types.NamespacedName
	stsName      types.NamespacedName
	logger       logr.Logger
}

// NewClusterResource creates ClusterResource
func NewClusterResource(
	client k8sclient.Client,
	pandaCluster *redpandav1alpha1.Cluster,
	scheme *runtime.Scheme,
	serviceFQDN string,
	nodePortName types.NamespacedName,
	stsName types.NamespacedName,
	logger logr.Logger,
) *ClusterResource {
	return &ClusterResource{
		client,
		scheme,
		pandaCluster,
		serviceFQDN,
		nodePortName,
		stsName,
		logger.WithValues("Kind", clusterKind()),
	}
}

// Ensure will manage v1alpha1.Cluster custom resource
func (c *ClusterResource) Ensure(ctx context.Context) error {
	var observedPods corev1.PodList

	err := c.List(ctx, &observedPods, &k8sclient.ListOptions{
		LabelSelector: labels.ForCluster(c.pandaCluster).AsClientSelector(),
		Namespace:     c.pandaCluster.Namespace,
	})
	if err != nil {
		return fmt.Errorf("failed to retrieve pods redpanda pods: %w", err)
	}

	observedNodesInternal := make([]string, 0, len(observedPods.Items))
	for i := range observedPods.Items {
		observedNodesInternal = append(observedNodesInternal,
			fmt.Sprintf("%s.%s", observedPods.Items[i].Name, c.serviceFQDN))
	}

	observedNodesExternal, err := c.createExternalNodesList(ctx, observedPods.Items)
	if err != nil {
		return fmt.Errorf("failed to construct external node list: %w", err)
	}

	if !reflect.DeepEqual(observedNodesInternal, c.pandaCluster.Status.Nodes.Internal) ||
		!reflect.DeepEqual(observedNodesExternal, c.pandaCluster.Status.Nodes.External) {
		c.pandaCluster.Status.Nodes.Internal = observedNodesInternal
		c.pandaCluster.Status.Nodes.External = observedNodesExternal

		if err = c.Status().Update(ctx, c.pandaCluster); err != nil {
			return fmt.Errorf("failed to update cluster status nodes: %w", err)
		}
	}

	sts := appsv1.StatefulSet{}
	if err = c.Get(ctx, c.stsName, &sts); err != nil {
		return fmt.Errorf("failed to retrieve statefulset %s: %w", c.stsName, err)
	}

	if !reflect.DeepEqual(sts.Status.ReadyReplicas, c.pandaCluster.Status.Replicas) {
		c.pandaCluster.Status.Replicas = sts.Status.ReadyReplicas

		if err = c.Status().Update(ctx, c.pandaCluster); err != nil {
			return fmt.Errorf("unable to update cluster status replicas: %w", err)
		}
	}

	return nil
}

func (c *ClusterResource) createExternalNodesList(
	ctx context.Context, pods []corev1.Pod,
) ([]string, error) {
	if !c.pandaCluster.Spec.ExternalConnectivity {
		return []string{}, nil
	}

	var nodePortSvc corev1.Service
	if err := c.Get(ctx, c.nodePortName, &nodePortSvc); err != nil {
		return []string{}, fmt.Errorf("failed to retrieve node port service %s: %w", c.nodePortName, err)
	}

	if len(nodePortSvc.Spec.Ports) != 1 || nodePortSvc.Spec.Ports[0].NodePort == 0 {
		return []string{}, fmt.Errorf("node port service %s: %w", c.nodePortName, errNodePortMissing)
	}

	var node corev1.Node
	observedNodesExternal := make([]string, 0, len(pods))
	for i := range pods {
		if err := c.Get(ctx, types.NamespacedName{Name: pods[i].Spec.NodeName}, &node); err != nil {
			return []string{}, fmt.Errorf("failed to retrieve node %s: %w", pods[i].Spec.NodeName, err)
		}

		observedNodesExternal = append(observedNodesExternal,
			fmt.Sprintf("%s:%d",
				getExternalIP(&node),
				getNodePort(&nodePortSvc),
			))
	}
	return observedNodesExternal, nil
}

func getExternalIP(node *corev1.Node) string {
	if node == nil {
		return ""
	}
	for _, address := range node.Status.Addresses {
		if address.Type == corev1.NodeExternalIP {
			return address.Address
		}
	}
	return ""
}

func getNodePort(svc *corev1.Service) int32 {
	if svc == nil {
		return -1
	}
	for _, port := range svc.Spec.Ports {
		if port.NodePort != 0 {
			return port.NodePort
		}
	}
	return 0
}

// Obj can not be called
func (c *ClusterResource) Obj() (k8sclient.Object, error) {
	panic("should be never called")
}

// Key returns namespace/name object that is used to identify object.
// For reference please visit types.NamespacedName docs in k8s.io/apimachinery
func (c *ClusterResource) Key() types.NamespacedName {
	return types.NamespacedName{Name: c.pandaCluster.Name, Namespace: c.pandaCluster.Namespace}
}

// Kind returns v1alpha1.Cluster kind
func (c *ClusterResource) Kind() string {
	return clusterKind()
}

func clusterKind() string {
	var c redpandav1alpha1.Cluster
	return c.Kind
}
