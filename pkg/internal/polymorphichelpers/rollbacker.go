/*
Copyright 2018 The Kubernetes Authors.

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

package polymorphichelpers

import (
	kruiseclientsets "github.com/openkruise/kruise-api/client/clientset/versioned"
	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/client-go/kubernetes"
)

// Returns a Rollbacker for changing the rollback version of the specified RESTMapping type or an error
func rollbacker(restClientGetter genericclioptions.RESTClientGetter, mapping *meta.RESTMapping) (Rollbacker, error) {
	clientConfig, err := restClientGetter.ToRESTConfig()
	if err != nil {
		return nil, err
	}
	external, err := kubernetes.NewForConfig(clientConfig)
	if err != nil {
		return nil, err
	}
	kruiseExternal, err := kruiseclientsets.NewForConfig(clientConfig)
	if err != nil {
		return nil, err
	}

	return RollbackerFor(mapping.GroupVersionKind.GroupKind(), external, kruiseExternal)
}
