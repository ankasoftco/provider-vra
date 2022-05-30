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
	token := "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6InNpZ25pbmdfMiJ9.eyJzdWIiOiJ2bXdhcmUuY29tOjIxOGI2YTRlLWQ0MTYtNDk2MS1hMjBlLWY0NGZmYmNiZjhhMCIsImlzcyI6Imh0dHBzOi8vZ2F6LXByZXZpZXcuY3NwLXZpZG0tcHJvZC5jb20iLCJjb250ZXh0X25hbWUiOiIzNTcyMDIzMS0zMmIyLTQ1ZTktYWFiMS0yOTljNmNjNDNkNzQiLCJfbm9uY2UiOiI0NTM0ZmExMC1jNjE3LTExZWMtODVjMS03YjAzOTkyNWVmMTYiLCJhenAiOiJwcm92aXNpb25pbmciLCJhdXRob3JpemF0aW9uX2RldGFpbHMiOltdLCJkb21haW4iOiJ2bXdhcmUuY29tIiwiY29udGV4dCI6IjU0YTYzNTUxLTg5YmItNDk1OS1iMDU2LTY0ZmVjYmZiNTM5NyIsInBlcm1zIjpbImNzcDpvcmdfb3duZXIiLCJleHRlcm5hbC9Zdy1IeUJlUXpqQ1hrTDJ3UVNlR3dhdUotbUFfL3JlZ2lvbjp1cyIsImV4dGVybmFsL1A2NXpNRGFZQTJaalRHVVJMWExhelFyODF0TV8vcmVnaW9uOnVzIiwiY3NwOnN1cHBvcnRfdXNlciIsImV4dGVybmFsL1l3LUh5QmVRempDWGtMMndRU2VHd2F1Si1tQV8vY2F0YWxvZzp1c2VyIiwiZXh0ZXJuYWwvdWx2cXRONDE0MWJlQ1Qyb09uYmotd2xrekdnXy9yZWdpb246dXMiLCJjc3A6b3JnX21lbWJlciIsImV4dGVybmFsL3VsdnF0TjQxNDFiZUNUMm9PbmJqLXdsa3pHZ18vQ29kZVN0cmVhbTp1c2VyIiwiZXh0ZXJuYWwvUDY1ek1EYVlBMlpqVEdVUkxYTGF6UXI4MXRNXy9hdXRvbWF0aW9uc2VydmljZTpjbG91ZF9hZG1pbiIsImV4dGVybmFsL3VsdnF0TjQxNDFiZUNUMm9PbmJqLXdsa3pHZ18vQ29kZVN0cmVhbTphZG1pbmlzdHJhdG9yIiwiZXh0ZXJuYWwvdWx2cXRONDE0MWJlQ1Qyb09uYmotd2xrekdnXy9Db2RlU3RyZWFtOmRldmVsb3BlciIsImV4dGVybmFsL1l3LUh5QmVRempDWGtMMndRU2VHd2F1Si1tQV8vY2F0YWxvZzphZG1pbiJdLCJleHAiOjE2NTM2NjE1MTcsImlhdCI6MTY1MzYzMjcxNywianRpIjoiNzgyMTUzMjctNmQyZi00NzllLTg4YzQtNDZkZTM4NTU0ZjVkIiwiYWNjdCI6InVkZW1pcnRhc0B2bXdhcmUuY29tIiwidXNlcm5hbWUiOiJ1ZGVtaXJ0YXMifQ.UJeaL82G9tIhi3v97HjLdBcfyFiFOLq6IPJ-Z6j8qZ17vwV3IC1KfWhMnHRe8Ffbv6fB-EsmJgW9akeEF55NDqlUFpbs51I4TYnUp3yoJsSzjaL3pWQzvFwfqRwIC9OsWw_AfjUz9QvAe70qAc_K4uiepjbVWq2MekMF_WV9YzrIJQLWOfy04pk2W137oR6CX570VC-BBMEMdqw8SmgpN9noYpuBnR-QYWaC0NXHZfrjEPNCR2ui5IYjN4dnRsYDfTnKTQDQdJTNlIpPPYehrMB0Bkby9F4gTCcjWZj4tbzbPZp3XxDCCwJ0n-QOyzUw_ZEi6gtAtklXw7naDMwtzA"

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
