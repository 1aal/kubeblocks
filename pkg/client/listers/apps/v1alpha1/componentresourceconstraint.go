/*
Copyright (C) 2022-2023 ApeCloud Co., Ltd

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

package v1alpha1

import (
	v1alpha1 "github.com/1aal/kubeblocks/apis/apps/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ComponentResourceConstraintLister helps list ComponentResourceConstraints.
// All objects returned here must be treated as read-only.
type ComponentResourceConstraintLister interface {
	// List lists all ComponentResourceConstraints in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ComponentResourceConstraint, err error)
	// Get retrieves the ComponentResourceConstraint from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ComponentResourceConstraint, error)
	ComponentResourceConstraintListerExpansion
}

// componentResourceConstraintLister implements the ComponentResourceConstraintLister interface.
type componentResourceConstraintLister struct {
	indexer cache.Indexer
}

// NewComponentResourceConstraintLister returns a new ComponentResourceConstraintLister.
func NewComponentResourceConstraintLister(indexer cache.Indexer) ComponentResourceConstraintLister {
	return &componentResourceConstraintLister{indexer: indexer}
}

// List lists all ComponentResourceConstraints in the indexer.
func (s *componentResourceConstraintLister) List(selector labels.Selector) (ret []*v1alpha1.ComponentResourceConstraint, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ComponentResourceConstraint))
	})
	return ret, err
}

// Get retrieves the ComponentResourceConstraint from the index for a given name.
func (s *componentResourceConstraintLister) Get(name string) (*v1alpha1.ComponentResourceConstraint, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("componentresourceconstraint"), name)
	}
	return obj.(*v1alpha1.ComponentResourceConstraint), nil
}
