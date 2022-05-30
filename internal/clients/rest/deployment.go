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
	url2 "net/url"

	"github.com/crossplane/provider-vra/internal/clients/deployment"
)

// Create creates the web hook
func (c *Client) Create(ctx context.Context, depOptions deployment.CreateDeploymentOptions) (deployment.ResponseCreateDeploymentOptions, error) {
	marshalledPayload, err := json.Marshal(depOptions)
	if err != nil {
		return deployment.ResponseCreateDeploymentOptions{}, err
	}

	url := c.BaseURL + fmt.Sprintf("/catalog/api/items/%s/request",
		url2.PathEscape(depOptions.CatalogItemID))

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(marshalledPayload))
	if err != nil {
		return deployment.ResponseCreateDeploymentOptions{}, err
	}

	var response []deployment.ResponseCreateDeploymentOptions
	if err := c.sendRequest(req, &response); err != nil {
		return deployment.ResponseCreateDeploymentOptions{}, err
	}
	return response[0], nil
}

// Get gets the vra
func (c *Client) Get(ctx context.Context, id string) (deployment.GetDeploymentOptions, error) {
	url := c.BaseURL + fmt.Sprintf("/iaas/api/deployments/%s", id)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return deployment.GetDeploymentOptions{}, err
	}

	// The documentation says this is a paged API but it is not
	var payload deployment.GetDeploymentOptions
	if err := c.sendRequest(req, &payload); err != nil {
		fmt.Println("failed get ops umit")
		return deployment.GetDeploymentOptions{}, fmt.Errorf("GetDeployment(%s): %w", id, err)
	}

	return payload, nil
}

// Delete deletes the web hook
func (c *Client) Delete(ctx context.Context, depOptions deployment.DeleteDeploymentOptions, id string) error {
	marshalledPayload, err := json.Marshal(depOptions)
	if err != nil {
		return err
	}

	url := c.BaseURL + fmt.Sprintf("/deployment/api/deployments/%s/requests", id)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(marshalledPayload))
	if err != nil {
		return err
	}

	return c.sendRequest(req, nil)
}
