/*
Copyright 2020 The Kubernetes Authors.

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

package v1beta1

import (
	v1beta1 "github.com/kubernetes-sigs/service-catalog/pkg/apis/servicecatalog/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ExtensionRequestLister helps list ExtensionRequests.
type ExtensionRequestLister interface {
	// List lists all ExtensionRequests in the indexer.
	List(selector labels.Selector) (ret []*v1beta1.ExtensionRequest, err error)
	// ExtensionRequests returns an object that can list and get ExtensionRequests.
	ExtensionRequests(namespace string) ExtensionRequestNamespaceLister
	ExtensionRequestListerExpansion
}

// extensionRequestLister implements the ExtensionRequestLister interface.
type extensionRequestLister struct {
	indexer cache.Indexer
}

// NewExtensionRequestLister returns a new ExtensionRequestLister.
func NewExtensionRequestLister(indexer cache.Indexer) ExtensionRequestLister {
	return &extensionRequestLister{indexer: indexer}
}

// List lists all ExtensionRequests in the indexer.
func (s *extensionRequestLister) List(selector labels.Selector) (ret []*v1beta1.ExtensionRequest, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.ExtensionRequest))
	})
	return ret, err
}

// ExtensionRequests returns an object that can list and get ExtensionRequests.
func (s *extensionRequestLister) ExtensionRequests(namespace string) ExtensionRequestNamespaceLister {
	return extensionRequestNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ExtensionRequestNamespaceLister helps list and get ExtensionRequests.
type ExtensionRequestNamespaceLister interface {
	// List lists all ExtensionRequests in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1beta1.ExtensionRequest, err error)
	// Get retrieves the ExtensionRequest from the indexer for a given namespace and name.
	Get(name string) (*v1beta1.ExtensionRequest, error)
	ExtensionRequestNamespaceListerExpansion
}

// extensionRequestNamespaceLister implements the ExtensionRequestNamespaceLister
// interface.
type extensionRequestNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ExtensionRequests in the indexer for a given namespace.
func (s extensionRequestNamespaceLister) List(selector labels.Selector) (ret []*v1beta1.ExtensionRequest, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.ExtensionRequest))
	})
	return ret, err
}

// Get retrieves the ExtensionRequest from the indexer for a given namespace and name.
func (s extensionRequestNamespaceLister) Get(name string) (*v1beta1.ExtensionRequest, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1beta1.Resource("extensionrequest"), name)
	}
	return obj.(*v1beta1.ExtensionRequest), nil
}
