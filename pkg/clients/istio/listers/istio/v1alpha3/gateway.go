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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha3

import (
	v1alpha3 "github.com/knative/pkg/apis/istio/v1alpha3"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// GatewayLister helps list Gateways.
type GatewayLister interface {
	// List lists all Gateways in the indexer.
	List(selector labels.Selector) (ret []*v1alpha3.Gateway, err error)
	// Gateways returns an object that can list and get Gateways.
	Gateways(namespace string) GatewayNamespaceLister
	GatewayListerExpansion
}

// gatewayLister implements the GatewayLister interface.
type gatewayLister struct {
	indexer cache.Indexer
}

// NewGatewayLister returns a new GatewayLister.
func NewGatewayLister(indexer cache.Indexer) GatewayLister {
	return &gatewayLister{indexer: indexer}
}

// List lists all Gateways in the indexer.
func (s *gatewayLister) List(selector labels.Selector) (ret []*v1alpha3.Gateway, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha3.Gateway))
	})
	return ret, err
}

// Gateways returns an object that can list and get Gateways.
func (s *gatewayLister) Gateways(namespace string) GatewayNamespaceLister {
	return gatewayNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// GatewayNamespaceLister helps list and get Gateways.
type GatewayNamespaceLister interface {
	// List lists all Gateways in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha3.Gateway, err error)
	// Get retrieves the Gateway from the indexer for a given namespace and name.
	Get(name string) (*v1alpha3.Gateway, error)
	GatewayNamespaceListerExpansion
}

// gatewayNamespaceLister implements the GatewayNamespaceLister
// interface.
type gatewayNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Gateways in the indexer for a given namespace.
func (s gatewayNamespaceLister) List(selector labels.Selector) (ret []*v1alpha3.Gateway, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha3.Gateway))
	})
	return ret, err
}

// Get retrieves the Gateway from the indexer for a given namespace and name.
func (s gatewayNamespaceLister) Get(name string) (*v1alpha3.Gateway, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha3.Resource("gateway"), name)
	}
	return obj.(*v1alpha3.Gateway), nil
}
