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

package mgw

import (
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
	wso2v1alpha1 "github.com/wso2/k8s-api-operator/api-operator/pkg/apis/wso2/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func ExternalIP (client *client.Client, apiInstance *wso2v1alpha1.API, operatorMode string, svc *corev1.Service,
	ingressConfData map[string]string, openshiftConfData map[string]string) string {

	var log = logf.Log.WithName("endpoint value")
	var ip string
	if operatorMode == "default" {
		loadBalancerFound := svc.Status.LoadBalancer.Ingress
		ip = ""
		for _, elem := range loadBalancerFound {
			ip += elem.IP
		}
		apiInstance.Spec.ApiEndPoint = ip
		log.Info("IP value is :" + ip)
		log.Info("ENDPOINT value in default mode is ","apiEndpoint",apiInstance.Spec.ApiEndPoint)
	}
	if operatorMode == "ingress" {
		ingressHostConf := ingressConfData[ingressHostName]
		log.Info("Host Name is :" + ingressHostConf)
		apiInstance.Spec.ApiEndPoint = ingressHostConf
		log.Info("ENDPOINT value in ingress mode is","apiEndpoint",apiInstance.Spec.ApiEndPoint)
		ip = "<pending>"
	}
	if operatorMode == "route" {
		routeHostConf := openshiftConfData[routeHost]
		log.Info("Host Name is :" + routeHostConf)
		apiInstance.Spec.ApiEndPoint = routeHostConf
		log.Info("ENDPOINT value in route mode is","apiEndpoint",apiInstance.Spec.ApiEndPoint)
		ip = "<pending>"
	}
	if apiInstance.Spec.ApiEndPoint == "" {
		apiInstance.Spec.ApiEndPoint = "<pending>"
		log.Info("ENDPOINT value after updating is","apiEndpoint" ,apiInstance.Spec.ApiEndPoint)
	}

	return ip
}
