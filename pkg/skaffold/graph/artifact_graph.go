/*
Copyright 2021 The Skaffold Authors

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

package graph

import latest_v1 "github.com/GoogleContainerTools/skaffold/pkg/skaffold/schema/latest/v1"

// Artifact is the result corresponding to each successful build.
type Artifact struct {
	ImageName string `json:"imageName"`
	Tag       string `json:"tag"`
}

// ArtifactGraph is a map of [artifact image : artifact definition]
type ArtifactGraph map[string]*latest_v1.Artifact

// ToArtifactGraph creates an instance of `ArtifactGraph` from `[]*latest_v1.Artifact`
func ToArtifactGraph(artifacts []*latest_v1.Artifact) ArtifactGraph {
	m := make(map[string]*latest_v1.Artifact)
	for _, a := range artifacts {
		m[a.ImageName] = a
	}
	return m
}

// Dependencies returns the de-referenced slice of required artifacts for a given artifact
func (g ArtifactGraph) Dependencies(a *latest_v1.Artifact) []*latest_v1.Artifact {
	var sl []*latest_v1.Artifact
	for _, d := range a.Dependencies {
		sl = append(sl, g[d.ImageName])
	}
	return sl
}