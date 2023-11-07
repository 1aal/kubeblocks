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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	appsv1alpha1 "github.com/1aal/kubeblocks/apis/apps/v1alpha1"
	versioned "github.com/1aal/kubeblocks/pkg/client/clientset/versioned"
	internalinterfaces "github.com/1aal/kubeblocks/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/1aal/kubeblocks/pkg/client/listers/apps/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// BackupPolicyTemplateInformer provides access to a shared informer and lister for
// BackupPolicyTemplates.
type BackupPolicyTemplateInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.BackupPolicyTemplateLister
}

type backupPolicyTemplateInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewBackupPolicyTemplateInformer constructs a new informer for BackupPolicyTemplate type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewBackupPolicyTemplateInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredBackupPolicyTemplateInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredBackupPolicyTemplateInformer constructs a new informer for BackupPolicyTemplate type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredBackupPolicyTemplateInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AppsV1alpha1().BackupPolicyTemplates().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AppsV1alpha1().BackupPolicyTemplates().Watch(context.TODO(), options)
			},
		},
		&appsv1alpha1.BackupPolicyTemplate{},
		resyncPeriod,
		indexers,
	)
}

func (f *backupPolicyTemplateInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredBackupPolicyTemplateInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *backupPolicyTemplateInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&appsv1alpha1.BackupPolicyTemplate{}, f.defaultInformer)
}

func (f *backupPolicyTemplateInformer) Lister() v1alpha1.BackupPolicyTemplateLister {
	return v1alpha1.NewBackupPolicyTemplateLister(f.Informer().GetIndexer())
}
