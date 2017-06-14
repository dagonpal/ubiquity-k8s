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

// This file was automatically generated by lister-gen

package internalversion

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	admissionregistration "k8s.io/kubernetes/pkg/apis/admissionregistration"
)

// ExternalAdmissionHookConfigurationLister helps list ExternalAdmissionHookConfigurations.
type ExternalAdmissionHookConfigurationLister interface {
	// List lists all ExternalAdmissionHookConfigurations in the indexer.
	List(selector labels.Selector) (ret []*admissionregistration.ExternalAdmissionHookConfiguration, err error)
	// ExternalAdmissionHookConfigurations returns an object that can list and get ExternalAdmissionHookConfigurations.
	ExternalAdmissionHookConfigurations(namespace string) ExternalAdmissionHookConfigurationNamespaceLister
	ExternalAdmissionHookConfigurationListerExpansion
}

// externalAdmissionHookConfigurationLister implements the ExternalAdmissionHookConfigurationLister interface.
type externalAdmissionHookConfigurationLister struct {
	indexer cache.Indexer
}

// NewExternalAdmissionHookConfigurationLister returns a new ExternalAdmissionHookConfigurationLister.
func NewExternalAdmissionHookConfigurationLister(indexer cache.Indexer) ExternalAdmissionHookConfigurationLister {
	return &externalAdmissionHookConfigurationLister{indexer: indexer}
}

// List lists all ExternalAdmissionHookConfigurations in the indexer.
func (s *externalAdmissionHookConfigurationLister) List(selector labels.Selector) (ret []*admissionregistration.ExternalAdmissionHookConfiguration, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*admissionregistration.ExternalAdmissionHookConfiguration))
	})
	return ret, err
}

// ExternalAdmissionHookConfigurations returns an object that can list and get ExternalAdmissionHookConfigurations.
func (s *externalAdmissionHookConfigurationLister) ExternalAdmissionHookConfigurations(namespace string) ExternalAdmissionHookConfigurationNamespaceLister {
	return externalAdmissionHookConfigurationNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ExternalAdmissionHookConfigurationNamespaceLister helps list and get ExternalAdmissionHookConfigurations.
type ExternalAdmissionHookConfigurationNamespaceLister interface {
	// List lists all ExternalAdmissionHookConfigurations in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*admissionregistration.ExternalAdmissionHookConfiguration, err error)
	// Get retrieves the ExternalAdmissionHookConfiguration from the indexer for a given namespace and name.
	Get(name string) (*admissionregistration.ExternalAdmissionHookConfiguration, error)
	ExternalAdmissionHookConfigurationNamespaceListerExpansion
}

// externalAdmissionHookConfigurationNamespaceLister implements the ExternalAdmissionHookConfigurationNamespaceLister
// interface.
type externalAdmissionHookConfigurationNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ExternalAdmissionHookConfigurations in the indexer for a given namespace.
func (s externalAdmissionHookConfigurationNamespaceLister) List(selector labels.Selector) (ret []*admissionregistration.ExternalAdmissionHookConfiguration, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*admissionregistration.ExternalAdmissionHookConfiguration))
	})
	return ret, err
}

// Get retrieves the ExternalAdmissionHookConfiguration from the indexer for a given namespace and name.
func (s externalAdmissionHookConfigurationNamespaceLister) Get(name string) (*admissionregistration.ExternalAdmissionHookConfiguration, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(admissionregistration.Resource("externaladmissionhookconfiguration"), name)
	}
	return obj.(*admissionregistration.ExternalAdmissionHookConfiguration), nil
}
