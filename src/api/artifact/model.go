// Copyright Project Harbor Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package artifact

import (
	"fmt"

	"github.com/goharbor/harbor/src/api/tag"
	cmodels "github.com/goharbor/harbor/src/common/models"
	"github.com/goharbor/harbor/src/common/utils"
	"github.com/goharbor/harbor/src/pkg/artifact"
)

// Artifact is the overall view of artifact
type Artifact struct {
	artifact.Artifact
	Tags          []*tag.Tag               `json:"tags"`           // the list of tags that attached to the artifact
	AdditionLinks map[string]*AdditionLink `json:"addition_links"` // the resource link for build history(image), values.yaml(chart), dependency(chart), etc
	Labels        []*cmodels.Label         `json:"labels"`
}

// SetAdditionLink set a addition link
func (artifact *Artifact) SetAdditionLink(addition, version string) {
	if artifact.AdditionLinks == nil {
		artifact.AdditionLinks = make(map[string]*AdditionLink)
	}

	projectName, repo := utils.ParseRepository(artifact.RepositoryName)
	href := fmt.Sprintf("/api/%s/projects/%s/repositories/%s/artifacts/%s/additions/%s", version, projectName, repo, artifact.Digest, addition)

	artifact.AdditionLinks[addition] = &AdditionLink{HREF: href, Absolute: false}
}

// AdditionLink is a link via that the addition can be fetched
type AdditionLink struct {
	HREF     string `json:"href"`
	Absolute bool   `json:"absolute"` // specify the href is an absolute URL or not
}

// Option is used to specify the properties returned when listing/getting artifacts
type Option struct {
	WithTag   bool
	TagOption *tag.Option // only works when WithTag is set to true
	WithLabel bool
}
