/*
Copyright 2019 The Kubernetes Authors.

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

package opentelemetry

import (
	"fmt"

	networking "k8s.io/api/networking/v1"
	"k8s.io/klog/v2"

	"k8s.io/ingress-nginx/internal/ingress/annotations/parser"
	"k8s.io/ingress-nginx/internal/ingress/resolver"
)

type opentelemetry struct {
	r resolver.Resolver
}

// Config contains the configuration to be used in the Ingress
type Config struct {
	Enabled             bool   `json:"enabled"`
	Set                 bool   `json:"set"`
	OpentelemetryConfig string `json:"config"`
	TrustEnabled        bool   `json:"trust-enabled"`
	TrustSet            bool   `json:"trust-set"`
	OperationName       string `json:"operation-name"`
}

// Equal tests for equality between two Config types
func (bd1 *Config) Equal(bd2 *Config) bool {

	if bd1.Enabled != bd2.Enabled {
		return false
	}

	if bd1.OpentelemetryConfig != bd2.OpentelemetryConfig {
		return false
	}

	if bd1.TrustEnabled != bd2.TrustEnabled {
		return false
	}

	return true
}

func NewParser(r resolver.Resolver) parser.IngressAnnotation {
	return opentelemetry{r}
}

// Parse parses the annotations to look for opentelemetry configurations
func (c opentelemetry) Parse(ing *networking.Ingress) (interface{}, error) {
	enabled, err := parser.GetBoolAnnotation("enable-opentelemetry", ing)
	if err != nil {
		klog.Errorf("#####1 - %v.", err)
		return &Config{}, nil
	}
	fmt.Println("######", enabled)
	if !enabled {
		klog.Errorf("#####1 - %v.", err)
		return &Config{Set: true, Enabled: false}, nil
	}
	conf, err := parser.GetStringAnnotation("opentelemetry-config", ing)
	fmt.Println("######", conf)
	if err != nil {
		klog.Errorf("OpenTelementry config file (opentelemetry_config) must be set - %v.", err)
		return &Config{OperationName: "new_op_name"}, nil
	}

	trustEnabled, err := parser.GetBoolAnnotation("opentelemetry-trust-incoming-span", ing)
	if err != nil {
		klog.Errorf("OpenTelementry config file (opentelemetry_config) must be set - %v.", err)
		return &Config{Set: true, Enabled: enabled, OpentelemetryConfig: conf, TrustEnabled: false}, nil
	}
	operationName, err := parser.GetStringAnnotation("opentelemetry-operation-name", ing)
	fmt.Println("######", operationName, err)

	return &Config{Set: true, Enabled: enabled, OpentelemetryConfig: conf, TrustEnabled: trustEnabled, TrustSet: true, OperationName: operationName}, nil
}
