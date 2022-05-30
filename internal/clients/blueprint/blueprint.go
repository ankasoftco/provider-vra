/*
Copyright 2021 The Crossplane Authors.

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

package blueprint

import (
	"context"

	"github.com/crossplane/provider-vra/apis/blueprint/v1alpha1"
)

// Client is the API for creating/listing/deleting/getting deployments
type Client interface {
	CreateBlueprint(ctx context.Context, blueprint CreateBlueprintOptions) (result ResponseCreateBlueprintOptions, err error)
	GetBlueprint(ctx context.Context, id string) (result GetBlueprintOptions, err error)
}

// CreateBlueprintOptions request payload
type CreateBlueprintOptions struct {
	Content     string `json:"content,omitempty"`
	Description string `json:"description,omitempty"`
	Name        string `json:"name,omitempty"`
	ProjectID   string `json:"projectId,omitempty"`
}

// GetBlueprintOptions request payload
type GetBlueprintOptions struct {
	Content string `json:"projectId,omitempty"`
	ID      string `json:"id,omitempty"`
}

// ResponseCreateBlueprintOptions response of the create a deployment with catalog item
type ResponseCreateBlueprintOptions struct {
	UpdatedAt string `json:"updatedAt,omitempty"`
	ID        string `json:"id,omitempty"`
}

// GenerateCreateBlueprintOptions returns the payload for REST API Client
func GenerateCreateBlueprintOptions(p *v1alpha1.BlueprintParameters) CreateBlueprintOptions {
	blueprint := CreateBlueprintOptions{
		Content:     p.Content,
		Description: p.Description,
		ProjectID:   p.ProjectID,
		Name:        p.Name,
	}

	return blueprint
}
