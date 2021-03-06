// Copyright (c)  WSO2 Inc. (http://www.wso2.org) All Rights Reserved.
//
// WSO2 Inc. licenses this file to you under the Apache License,
// Version 2.0 (the "License"); you may not use this file except
// in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package v1alpha2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// APISpec defines the desired state of API
type APISpec struct {
	// Mode of the API. The mode from the swagger definition will be overridden by this value.
	// Supports "privateJet", "sidecar", "<empty>".
	// Default value "<empty>".
	// +optional
	Mode Mode `json:"mode,omitempty"`
	// Update API definition creating a new docker image. Make a rolling update to the existing API.
	// with prefixing the timestamp value.
	// Default value "<empty>".
	// +optional
	UpdateTimeStamp string `json:"updateTimeStamp,omitempty"`
	// Replica count of the API.
	// Default value "<empty>".
	// +optional
	Replicas int `json:"replicas,omitempty"`
	// Override the exiting API docker image.
	// Default value "false".
	// +optional
	Override bool `json:"override,omitempty"`
	// Version of the API. The version from the swagger definition will be overridden by this value.
	// Default value "<empty>".
	// +optional
	Version string `json:"version,omitempty"`
	// Environment variables to be added to the API deployment.
	// Default value "<empty>".
	// +optional
	EnvironmentVariables []string `json:"environmentVariables,omitempty"`
	// Docker image of the API to be deployed. If specified, ignores the values of `UpdateTimeStamp`, `Override`.
	// Uses the given image for the deployment.
	// Default value "<empty>".
	// +optional
	Image       string `json:"image,omitempty"`
	ApiEndPoint string `json:"apiEndPoint,omitempty"`
	// Ingress Hostname that the API is being exposed.
	// Default value "<empty>".
	// +optional
	IngressHostname string `json:"ingressHostname,omitempty"`
	//Config map name of which the project zip or swagger file is included
	SwaggerConfigMapName string `json:"swaggerConfigMapName"`
	// Config map name of the param values of the API project
	// Default value "<empty>".
	// +optional
	ParamsValues string `json:"paramsValues,omitempty"`
	// Config map name of the certs values of the API project
	// Default value "<empty>".
	// +optional
	CertsValues string `json:"certsValues,omitempty"`
}

// APIStatus defines the observed state of API
type APIStatus struct {
	// replicas field in the status sub-resource will define the initial replica count allocated to the API.This will be the minimum replica count for a single API
	// Default value "<empty>".
	// +optional
	Replicas int `json:"replicas,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// API is the Schema for the apis API
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Age",type=date,JSONPath=`.metadata.creationTimestamp`
type API struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   APISpec   `json:"spec,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// APIList contains a list of API
type APIList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []API `json:"items"`
}

type Mode string

const (
	PrivateJet Mode = "privateJet"
	Sidecar    Mode = "sidecar"
	Shared     Mode = "shared"
	Serverless Mode = "serverless"
)

func (c Mode) String() string {
	return string(c)
}

func init() {
	SchemeBuilder.Register(&API{}, &APIList{})
}
