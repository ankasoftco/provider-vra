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

package fake

import (
	"context"

	"github.com/crossplane/provider-vra/internal/clients/vra"
)

var _ vra.DeploymentClient = &MockDeploymentClient{}

// MockDeploymentClient is a fake implementation of MockDeploymentClient
type MockDeploymentClient struct {
	vra.DeploymentClient

	MockCreate func(ctx context.Context) (result vra.Deployment, err error)
	MockGet    func(ctx context.Context, id int) (result vra.Deployment, err error)
}

// CreateWebhook calls the mock
func (c *MockDeploymentClient) CreateWebhook(ctx context.Context) (result vra.Deployment, err error) {
	return c.MockCreate(ctx)
}

// GetWebhook calls the mock
func (c *MockDeploymentClient) GetWebhook(ctx context.Context, id int) (result vra.Deployment, err error) {
	return c.MockGet(ctx, id)
}
