/*
Copyright 2017 The Kubernetes Authors.

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

package controller

import (
	"github.com/kubernetes-sigs/service-catalog/pkg/apis/servicecatalog/v1beta1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/klog"
)

func (c *controller) reconcileExtensionRequestKey(key string) error {
	klog.Infof("Extension reconcile %s", key)
	return nil
}

func (c *controller) extensionRequestAdd(obj interface{}) {
	c.enqueueExtensionRequest(obj)
}

func (c *controller) extensionRequestUpdate(oldObj, newObj interface{}) {
	c.enqueueExtensionRequest(newObj)
}

// extensionRequestDelete handles the ServiceExtensionRequest DELETED watch event
func (c *controller) extensionRequestDelete(obj interface{}) {
	extensionRequest, ok := obj.(*v1beta1.ExtensionRequest)
	if extensionRequest == nil || !ok {
		return
	}
}

// enqueueExtensionRequest adds the extensionRequest key to the work queue
func (c *controller) enqueueExtensionRequest(obj interface{}) {
	key, err := cache.DeletionHandlingMetaNamespaceKeyFunc(obj)
	if err != nil {
		return
	}

	c.extensionRequestQueue.Add(key)
}
