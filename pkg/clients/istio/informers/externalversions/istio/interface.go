/*
Copyright The Flagger Authors.

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

// Code generated by informer-gen. DO NOT EDIT.

package networking

import (
	internalinterfaces "github.com/evry-ace/waf-util/pkg/clients/istio/informers/externalversions/internalinterfaces"
	v1alpha3 "github.com/evry-ace/waf-util/pkg/clients/istio/informers/externalversions/istio/v1alpha3"
)

// Interface provides access to each of this group's versions.
type Interface interface {
	// V1alpha3 provides access to shared informers for resources in V1alpha3.
	V1alpha3() v1alpha3.Interface
}

type group struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &group{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// V1alpha3 returns a new v1alpha3.Interface.
func (g *group) V1alpha3() v1alpha3.Interface {
	return v1alpha3.New(g.factory, g.namespace, g.tweakListOptions)
}
