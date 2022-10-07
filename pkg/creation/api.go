/*
Copyright 2020 The Kruise Authors.

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

package creation

import "github.com/dong4325/kruise-tools-test/pkg/api"

type Control interface {
	Create(src api.ResourceRef, dst api.ResourceRef, opts Options) error
}

type Options struct {
	CopyReplicas bool
}
