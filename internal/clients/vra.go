/*
Copyright 2020 The Crossplane Authors.

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

package clients

import (
	"context"
	"net/http"

	"github.com/pkg/errors"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	"github.com/crossplane/crossplane-runtime/pkg/resource"

	"github.com/crossplane/provider-vra/apis/v1alpha1"
	"github.com/crossplane/provider-vra/internal/clients/deployment"
	"github.com/crossplane/provider-vra/internal/clients/rest"
)

// Config for vRA Client authentication struct
type Config struct {
	BaseURL      string
	RefreshToken string
}

// NewClient creates new vRA Client with provided base URL and credentials
func NewClient(c Config) *rest.Client {
	httpClient := http.Client{
		Transport: &http.Transport{},
	}

	// Get a bearer token

	token := "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6InNpZ25pbmdfMiJ9.eyJzdWIiOiJ2bXdhcmUuY29tOjRiYzYwODliLWZlYWEtNGMyMS05MjMyLTZjY2YzMTczYTAzYyIsImlzcyI6Imh0dHBzOi8vZ2F6LmNzcC12aWRtLXByb2QuY29tIiwiY29udGV4dF9uYW1lIjoiYTE0OGNkNDMtYmE0OC00YTdlLWJjMTUtMzQzNDI5NGQxNmY0IiwiX25vbmNlIjoiNmMzNDJiNjAtMDMyZS0xMWVjLTk5MTUtMTVlMzJjOGU5ZGRmIiwiYXpwIjoicHJvdmlzaW9uaW5nIiwiYXV0aG9yaXphdGlvbl9kZXRhaWxzIjpbXSwiZG9tYWluIjoidm13YXJlLmNvbSIsImNvbnRleHQiOiI3NGRmYmU0Ni0xZTAzLTRjYmYtYmVmZi00NmNlYTc5ZWQ0MjgiLCJwZXJtcyI6WyJleHRlcm5hbC9Zdy1IeUJlUXpqQ1hrTDJ3UVNlR3dhdUotbUFfL2NhdGFsb2c6dXNlciIsImNzcDpvcmdfbWVtYmVyIiwiZXh0ZXJuYWwvWnk5MjRtRTNkd24yQVN5VlpSME5uN2x1cGVBXy9hdXRvbWF0aW9uc2VydmljZTp1c2VyIiwiZXh0ZXJuYWwvWnk5MjRtRTNkd24yQVN5VlpSME5uN2x1cGVBXy9hdXRvbWF0aW9uc2VydmljZTpjbG91ZF9hZG1pbiIsImV4dGVybmFsL1l3LUh5QmVRempDWGtMMndRU2VHd2F1Si1tQV8vY2F0YWxvZzphZG1pbiJdLCJleHAiOjE2NTM2MTE3MTksImlhdCI6MTY1MzU4MjkxOSwianRpIjoiMDhiMWVhNjAtZWM1MS00NjEyLTk3ZjktNmZmODE2MzJiMDg2IiwiYWNjdCI6InVkZW1pcnRhc0B2bXdhcmUuY29tIiwidXNlcm5hbWUiOiJ1ZGVtaXJ0YXMifQ.MrkZ221qvjOJGaUlcOkfCoW1Yecbpr3JMJ4MSi1WHM7uAkYtD5urFhUDUfeEzQCLRFK785h4Q5yGuJiDFKHKPYeHujMczrEZX_CZbNzpI83BtfWRXu2p4AC_zeLVfOhz3tGZ2Ecw6hqPqqu4EbBXiaIY7ApT44sBoHqDHT3KODDiDOvei1yYGryEyv6jiHwyRydXcGtJEFW9E5gD9Me9kxOh7ZUbnJDfc3L5F3_pG5HfVNDmTlI2_6V6N8hGpgg59V9AQCKjAeIGcSGoNaJm57xLnClTOUWjLgqFioBUCNWdakxCzqrwDeBlaTkrhZkoC5V2wPg85O7Z7A9iHfCt7g"
	return &rest.Client{
		BearerToken: token,
		BaseURL:     c.BaseURL,
		HTTPClient:  &httpClient,
	}
}

// NewDeploymentClient creates a new client for the deployment api
func NewDeploymentClient(c Config) deployment.Client {
	return NewClient(c)
}

// GetConfig constructs a Config that can be used to authenticate to vRA
func GetConfig(ctx context.Context, c client.Client, mg resource.Managed) (*Config, error) {
	switch {
	case mg.GetProviderConfigReference() != nil:
		return UseProviderConfig(ctx, c, mg)
	default:
		return nil, errors.New("providerConfigRef is not given")
	}
}

// UseProviderConfig to produce a config that can be used to authenticate to AWS.
func UseProviderConfig(ctx context.Context, c client.Client, mg resource.Managed) (*Config, error) {
	pc := &v1alpha1.ProviderConfig{}
	if err := c.Get(ctx, types.NamespacedName{Name: mg.GetProviderConfigReference().Name}, pc); err != nil {
		return nil, errors.Wrap(err, "cannot get referenced Provider")
	}

	t := resource.NewProviderConfigUsageTracker(c, &v1alpha1.ProviderConfigUsage{})
	if err := t.Track(ctx, mg); err != nil {
		return nil, errors.Wrap(err, "cannot track ProviderConfig usage")
	}

	switch s := pc.Spec.Credentials.Source; s { //nolint:exhaustive
	case xpv1.CredentialsSourceSecret:
		csr := pc.Spec.Credentials.SecretRef
		if csr == nil {
			return nil, errors.New("no credentials secret referenced")
		}
		s := &corev1.Secret{}
		if err := c.Get(ctx, types.NamespacedName{Namespace: csr.Namespace, Name: csr.Name}, s); err != nil {
			return nil, errors.Wrap(err, "cannot get credentials secret")
		}
		return &Config{BaseURL: pc.Spec.BaseURL, RefreshToken: string(s.Data[csr.Key])}, nil
	default:
		return nil, errors.Errorf("credentials source %s is not currently supported", s)
	}
}
