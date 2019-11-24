// +build !ignore_autogenerated

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
// Code generated by openapi-gen. DO NOT EDIT.

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	"github.com/go-openapi/spec"
	"k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/argoproj/argo-events/pkg/apis/gateway/v1alpha1.EventSourceRef":             schema_pkg_apis_gateway_v1alpha1_EventSourceRef(ref),
		"github.com/argoproj/argo-events/pkg/apis/gateway/v1alpha1.Gateway":                    schema_pkg_apis_gateway_v1alpha1_Gateway(ref),
		"github.com/argoproj/argo-events/pkg/apis/gateway/v1alpha1.GatewayList":                schema_pkg_apis_gateway_v1alpha1_GatewayList(ref),
		"github.com/argoproj/argo-events/pkg/apis/gateway/v1alpha1.GatewayNotificationWatcher": schema_pkg_apis_gateway_v1alpha1_GatewayNotificationWatcher(ref),
		"github.com/argoproj/argo-events/pkg/apis/gateway/v1alpha1.GatewaySpec":                schema_pkg_apis_gateway_v1alpha1_GatewaySpec(ref),
		"github.com/argoproj/argo-events/pkg/apis/gateway/v1alpha1.GatewayStatus":              schema_pkg_apis_gateway_v1alpha1_GatewayStatus(ref),
		"github.com/argoproj/argo-events/pkg/apis/gateway/v1alpha1.NodeStatus":                 schema_pkg_apis_gateway_v1alpha1_NodeStatus(ref),
		"github.com/argoproj/argo-events/pkg/apis/gateway/v1alpha1.NotificationWatchers":       schema_pkg_apis_gateway_v1alpha1_NotificationWatchers(ref),
		"github.com/argoproj/argo-events/pkg/apis/gateway/v1alpha1.SensorNotificationWatcher":  schema_pkg_apis_gateway_v1alpha1_SensorNotificationWatcher(ref),
	}
}

func schema_pkg_apis_gateway_v1alpha1_EventSourceRef(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "EventSourceRef holds information about the EventSourceRef custom resource",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"name": {
						SchemaProps: spec.SchemaProps{
							Description: "Name of the event source",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"namespace": {
						SchemaProps: spec.SchemaProps{
							Description: "Namespace of the event source Default value is the namespace where referencing gateway is deployed",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"name"},
			},
		},
	}
}

func schema_pkg_apis_gateway_v1alpha1_Gateway(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Gateway is the definition of a gateway resource",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/argoproj/argo-events/pkg/apis/gateway/v1alpha1.GatewayStatus"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/argoproj/argo-events/pkg/apis/gateway/v1alpha1.GatewaySpec"),
						},
					},
				},
				Required: []string{"metadata", "status", "spec"},
			},
		},
		Dependencies: []string{
			"github.com/argoproj/argo-events/pkg/apis/gateway/v1alpha1.GatewaySpec", "github.com/argoproj/argo-events/pkg/apis/gateway/v1alpha1.GatewayStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_gateway_v1alpha1_GatewayList(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "GatewayList is the list of Gateway resources",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ListMeta"),
						},
					},
					"items": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type": "items",
							},
						},
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/argoproj/argo-events/pkg/apis/gateway/v1alpha1.Gateway"),
									},
								},
							},
						},
					},
				},
				Required: []string{"metadata", "items"},
			},
		},
		Dependencies: []string{
			"github.com/argoproj/argo-events/pkg/apis/gateway/v1alpha1.Gateway", "k8s.io/apimachinery/pkg/apis/meta/v1.ListMeta"},
	}
}

func schema_pkg_apis_gateway_v1alpha1_GatewayNotificationWatcher(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "GatewayNotificationWatcher is the gateway interested in listening to notifications from this gateway",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"name": {
						SchemaProps: spec.SchemaProps{
							Description: "Name is the gateway name",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"port": {
						SchemaProps: spec.SchemaProps{
							Description: "Port is http server port on which gateway is running",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"endpoint": {
						SchemaProps: spec.SchemaProps{
							Description: "Endpoint is REST API endpoint to post event to. Events are sent using HTTP POST method to this endpoint.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"namespace": {
						SchemaProps: spec.SchemaProps{
							Description: "Namespace of the gateway",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"name", "port", "endpoint"},
			},
		},
	}
}

func schema_pkg_apis_gateway_v1alpha1_GatewaySpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "GatewaySpec represents gateway specifications",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"template": {
						SchemaProps: spec.SchemaProps{
							Description: "Template is the pod specification for the gateway Refer https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.11/#pod-v1-core",
							Ref:         ref("k8s.io/api/core/v1.PodTemplateSpec"),
						},
					},
					"eventSourceRef": {
						SchemaProps: spec.SchemaProps{
							Description: "EventSourceRef refers to event-source that stores event source configurations for the gateway",
							Ref:         ref("github.com/argoproj/argo-events/pkg/apis/gateway/v1alpha1.EventSourceRef"),
						},
					},
					"type": {
						SchemaProps: spec.SchemaProps{
							Description: "Type is the type of gateway. Used as metadata.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"service": {
						SchemaProps: spec.SchemaProps{
							Description: "Service is the specifications of the service to expose the gateway Refer https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.11/#service-v1-core",
							Ref:         ref("k8s.io/api/core/v1.Service"),
						},
					},
					"watchers": {
						SchemaProps: spec.SchemaProps{
							Description: "Watchers are components which are interested listening to notifications from this gateway These only need to be specified when gateway dispatch mechanism is through HTTP POST notifications. In future, support for NATS, KAFKA will be added as a means to dispatch notifications in which case specifying watchers would be unnecessary.",
							Ref:         ref("github.com/argoproj/argo-events/pkg/apis/gateway/v1alpha1.NotificationWatchers"),
						},
					},
					"processorPort": {
						SchemaProps: spec.SchemaProps{
							Description: "Port on which the gateway event source processor is running on.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"eventProtocol": {
						SchemaProps: spec.SchemaProps{
							Description: "EventProtocol is the underlying protocol used to send events from gateway to watchers(components interested in listening to event from this gateway)",
							Ref:         ref("github.com/argoproj/argo-events/pkg/apis/common.EventProtocol"),
						},
					},
					"replica": {
						SchemaProps: spec.SchemaProps{
							Description: "Replica is the gateway deployment replicas",
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
				},
				Required: []string{"template", "type", "processorPort", "eventProtocol"},
			},
		},
		Dependencies: []string{
			"github.com/argoproj/argo-events/pkg/apis/common.EventProtocol", "github.com/argoproj/argo-events/pkg/apis/gateway/v1alpha1.EventSourceRef", "github.com/argoproj/argo-events/pkg/apis/gateway/v1alpha1.NotificationWatchers", "k8s.io/api/core/v1.PodTemplateSpec", "k8s.io/api/core/v1.Service"},
	}
}

func schema_pkg_apis_gateway_v1alpha1_GatewayStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "GatewayStatus contains information about the status of a gateway.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"phase": {
						SchemaProps: spec.SchemaProps{
							Description: "Phase is the high-level summary of the gateway",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"startedAt": {
						SchemaProps: spec.SchemaProps{
							Description: "StartedAt is the time at which this gateway was initiated",
							Ref:         ref("k8s.io/apimachinery/pkg/apis/meta/v1.Time"),
						},
					},
					"message": {
						SchemaProps: spec.SchemaProps{
							Description: "Message is a human readable string indicating details about a gateway in its phase",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"nodes": {
						SchemaProps: spec.SchemaProps{
							Description: "Nodes is a mapping between a node ID and the node's status it records the states for the configurations of gateway.",
							Type:        []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Allows: true,
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/argoproj/argo-events/pkg/apis/gateway/v1alpha1.NodeStatus"),
									},
								},
							},
						},
					},
				},
				Required: []string{"phase"},
			},
		},
		Dependencies: []string{
			"github.com/argoproj/argo-events/pkg/apis/gateway/v1alpha1.NodeStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.Time"},
	}
}

func schema_pkg_apis_gateway_v1alpha1_NodeStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "NodeStatus describes the status for an individual node in the gateway configurations. A single node can represent one configuration.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"id": {
						SchemaProps: spec.SchemaProps{
							Description: "ID is a unique identifier of a node within a sensor It is a hash of the node name",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"name": {
						SchemaProps: spec.SchemaProps{
							Description: "Name is a unique name in the node tree used to generate the node ID",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"displayName": {
						SchemaProps: spec.SchemaProps{
							Description: "DisplayName is the human readable representation of the node",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"phase": {
						SchemaProps: spec.SchemaProps{
							Description: "Phase of the node",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"message": {
						SchemaProps: spec.SchemaProps{
							Description: "Message store data or something to save for configuration",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"updateTime": {
						SchemaProps: spec.SchemaProps{
							Description: "UpdateTime is the time when node(gateway configuration) was updated",
							Ref:         ref("k8s.io/apimachinery/pkg/apis/meta/v1.MicroTime"),
						},
					},
				},
				Required: []string{"id", "name", "displayName", "phase"},
			},
		},
		Dependencies: []string{
			"k8s.io/apimachinery/pkg/apis/meta/v1.MicroTime"},
	}
}

func schema_pkg_apis_gateway_v1alpha1_NotificationWatchers(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "NotificationWatchers are components which are interested listening to notifications from this gateway",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"gateways": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type": "gateways",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "Gateways is the list of gateways interested in listening to notifications from this gateway",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/argoproj/argo-events/pkg/apis/gateway/v1alpha1.GatewayNotificationWatcher"),
									},
								},
							},
						},
					},
					"sensors": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type": "sensors",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "Sensors is the list of sensors interested in listening to notifications from this gateway",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/argoproj/argo-events/pkg/apis/gateway/v1alpha1.SensorNotificationWatcher"),
									},
								},
							},
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/argoproj/argo-events/pkg/apis/gateway/v1alpha1.GatewayNotificationWatcher", "github.com/argoproj/argo-events/pkg/apis/gateway/v1alpha1.SensorNotificationWatcher"},
	}
}

func schema_pkg_apis_gateway_v1alpha1_SensorNotificationWatcher(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "SensorNotificationWatcher is the sensor interested in listening to notifications from this gateway",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"name": {
						SchemaProps: spec.SchemaProps{
							Description: "Name is the name of the sensor",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"namespace": {
						SchemaProps: spec.SchemaProps{
							Description: "Namespace of the sensor",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"name"},
			},
		},
	}
}
