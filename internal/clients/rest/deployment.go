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
	"github.com/crossplane/provider-vra/internal/clients/vra"
	"net/http"
)

// Get gets the vra
func (c *Client) Get(ctx context.Context, id int) (vra.Deployment, error) {
	url := c.BaseUrl + fmt.Sprintf("/rest/api/1.0/projects/%d", id)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return vra.Deployment{}, err
	}

	// The documentation says this is a paged API but it is not
	var payload vra.Deployment
	if err := c.sendRequest(req, &payload); err != nil {
		return vra.Deployment{}, fmt.Errorf("GetWebhook(%d): %w", id, err)
	}

	return payload, nil
}

// Create creates the web hook
func (c *Client) Create(ctx context.Context, hook vra.Deployment) (vra.Deployment, error) {
	marshalledPayload, err := json.Marshal(hook)
	if err != nil {
		return vra.Deployment{}, err
	}

	url := c.BaseUrl + fmt.Sprintf("/rest/api/1.0/projects/")

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(marshalledPayload))
	if err != nil {
		return vra.Deployment{}, err
	}

	var response vra.Deployment
	if err := c.sendRequest(req, &response); err != nil {
		return vra.Deployment{}, err
	}
	return response, nil
}
