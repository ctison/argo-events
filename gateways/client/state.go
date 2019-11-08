/*
Copyright 2018 BlackRock, Inc.

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

package main

import (
	"time"

	"github.com/argoproj/argo-events/common"
	gtw "github.com/argoproj/argo-events/controllers/gateway"
	"github.com/argoproj/argo-events/pkg/apis/gateway"
	"github.com/argoproj/argo-events/pkg/apis/gateway/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EventSourceStatus encapsulates state of an event source
type EventSourceStatus struct {
	// Id of the event source
	Id string
	// Name of the event source
	Name string
	// Message
	Message string
	// Phase of the event source
	Phase v1alpha1.NodePhase
	// Gateway reference
	Gateway *v1alpha1.Gateway
}

// markGatewayNodePhase marks the node with a phase, returns the node
func (gatewayCfg *GatewayConfig) markGatewayNodePhase(nodeStatus *EventSourceStatus) *v1alpha1.NodeStatus {
	logger := gatewayCfg.logger.WithFields(
		map[string]interface{}{
			common.LabelNodeName: nodeStatus.Name,
			common.LabelPhase:    string(nodeStatus.Phase),
		},
	)

	logger.Infoln("marking node phase")

	node := gatewayCfg.getNodeByID(nodeStatus.Id)
	if node == nil {
		logger.Warnln("node is not initialized")
		return nil
	}
	if node.Phase != nodeStatus.Phase {
		logger.WithField("new-phase", string(nodeStatus.Phase)).Infoln("phase updated")
		node.Phase = nodeStatus.Phase
	}
	node.Message = nodeStatus.Message
	gatewayCfg.gateway.Status.Nodes[node.ID] = *node
	gatewayCfg.updated = true
	return node
}

// getNodeByName returns the node from this gateway for the nodeName
func (gatewayCfg *GatewayConfig) getNodeByID(nodeID string) *v1alpha1.NodeStatus {
	node, ok := gatewayCfg.gateway.Status.Nodes[nodeID]
	if !ok {
		return nil
	}
	return &node
}

// create a new node
func (gatewayCfg *GatewayConfig) initializeNode(nodeID string, nodeName string, messages string) v1alpha1.NodeStatus {
	if gatewayCfg.gateway.Status.Nodes == nil {
		gatewayCfg.gateway.Status.Nodes = make(map[string]v1alpha1.NodeStatus)
	}

	gatewayCfg.logger.WithField(common.LabelNodeName, nodeName).Infoln("node")

	node, ok := gatewayCfg.gateway.Status.Nodes[nodeID]
	if !ok {
		node = v1alpha1.NodeStatus{
			ID:          nodeID,
			Name:        nodeName,
			DisplayName: nodeName,
			StartedAt:   metav1.MicroTime{Time: time.Now().UTC()},
		}
	}
	node.Phase = v1alpha1.NodePhaseRunning
	node.Message = messages
	gatewayCfg.gateway.Status.Nodes[nodeID] = node

	gatewayCfg.logger.WithFields(
		map[string]interface{}{
			common.LabelNodeName: nodeName,
			"node-message":       node.Message,
		},
	).Infoln("node is running")

	gatewayCfg.updated = true
	return node
}

// UpdateGatewayResourceState updates gateway resource nodes state
func (gatewayCfg *GatewayConfig) UpdateGatewayResourceState(status *EventSourceStatus) {
	logger := gatewayCfg.logger
	if status.Phase != v1alpha1.NodePhaseResourceUpdate {
		logger = logger.WithField(common.LabelEventSource, status.Name).Logger
	}

	logger.Infoln("received a gateway state update notification")

	switch status.Phase {
	case v1alpha1.NodePhaseRunning:
		// init the node and mark it as running
		gatewayCfg.initializeNode(status.Id, status.Name, status.Message)

	case v1alpha1.NodePhaseCompleted, v1alpha1.NodePhaseError:
		gatewayCfg.markGatewayNodePhase(status)

	case v1alpha1.NodePhaseResourceUpdate:
		gatewayCfg.gateway = status.Gateway

	case v1alpha1.NodePhaseRemove:
		delete(gatewayCfg.gateway.Status.Nodes, status.Id)
		logger.Infoln("event source is removed")
		gatewayCfg.updated = true
	}

	if gatewayCfg.updated {
		// persist changes and create K8s event logging the change
		eventType := common.StateChangeEventType
		labels := map[string]string{
			common.LabelGatewayEventSourceName: status.Name,
			common.LabelGatewayName:            gatewayCfg.name,
			common.LabelGatewayEventSourceID:   status.Id,
			common.LabelOperation:              "persist_event_source_state",
		}
		updatedGw, err := gtw.PersistUpdates(gatewayCfg.gatewayClient, gatewayCfg.gateway, gatewayCfg.logger)
		if err != nil {
			logger.WithError(err).Errorln("failed to persist gateway resource updates, reverting to old state")
			eventType = common.EscalationEventType
		}

		// update gateway ref. in case of failure to persist updates, this is a deep copy of old gateway resource
		gatewayCfg.gateway = updatedGw
		labels[common.LabelEventType] = string(eventType)

		// generate a K8s event for persist event source state change
		if err := common.GenerateK8sEvent(gatewayCfg.k8sClient, status.Message, eventType, "event source state update", gatewayCfg.name, gatewayCfg.namespace, gatewayCfg.controllerInstanceID, gateway.Kind, labels); err != nil {
			logger.WithError(err).Errorln("failed to create K8s event to log event source state change")
		}
	}
	gatewayCfg.updated = false
}