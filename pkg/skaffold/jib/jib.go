/*
Copyright 2018 The Skaffold Authors

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

package jib

import (
	"github.com/GoogleContainerTools/skaffold/pkg/skaffold/schema/v1alpha3"
	"github.com/pkg/errors"
)

func GetDependenciesMaven(_ /*workspace*/ string, _ /*a*/ *v1alpha3.JibMavenArtifact) ([]string, error) {
	return nil, errors.New("jib maven support is unimplemented")
}

func GetDependenciesGradle(_ /*workspace*/ string, _ /*a*/ *v1alpha3.JibGradleArtifact) ([]string, error) {
	return nil, errors.New("jib gradle support is unimplemented")
}