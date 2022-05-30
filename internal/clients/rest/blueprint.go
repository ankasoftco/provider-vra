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

package rest

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/crossplane/provider-vra/internal/clients/blueprint"
)

// CreateBlueprint creates the web hook
func (c *Client) CreateBlueprint(ctx context.Context, blueOptions blueprint.CreateBlueprintOptions) (blueprint.ResponseCreateBlueprintOptions, error) {
	marshalledPayload, err := json.Marshal(blueOptions)
	if err != nil {
		return blueprint.ResponseCreateBlueprintOptions{}, err
	}

	url := c.BaseURL + "/blueprint/api/blueprints"

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(marshalledPayload))
	if err != nil {
		return blueprint.ResponseCreateBlueprintOptions{}, err
	}

	var response blueprint.ResponseCreateBlueprintOptions
	if err := c.sendRequest(req, &response); err != nil {
		return blueprint.ResponseCreateBlueprintOptions{}, err
	}
	return response, nil
}

// GetBlueprint gets the vra
func (c *Client) GetBlueprint(ctx context.Context, id string) (blueprint.GetBlueprintOptions, error) {
	url := c.BaseURL + fmt.Sprintf("/blueprint/api/blueprints/%s", id)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return blueprint.GetBlueprintOptions{}, err
	}

	// The documentation says this is a paged API but it is not
	var payload blueprint.GetBlueprintOptions
	if err := c.sendRequest(req, &payload); err != nil {
		return blueprint.GetBlueprintOptions{}, fmt.Errorf("GetBlueprint(%s): %w", id, err)
	}

	return payload, nil
}
