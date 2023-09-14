/*
Copyright 2022 The ANKA SOFTWARE Authors.
*/

package login

import (
	"github.com/crossplane/provider-vraprovider/internal/clients"

	"github.com/vmware/vra-sdk-go/pkg/client/login"
	"github.com/vmware/vra-sdk-go/pkg/models"
)

// NewLoginClient returns a new vRA Login service
func NewLoginClient(cfg clients.Config) login.ClientService {
	vra := clients.NewClient(cfg)

	return vra.Login
}

// GetBearerToken returns a bearer token for vRA Login service
func GetBearerToken(cfg clients.Config) (string, error) {
	apiClient := NewLoginClient(cfg)

	params := login.NewRetrieveAuthTokenParams().WithBody(
		&models.CspLoginSpecification{
			RefreshToken: &cfg.RefreshToken,
		},
	)

	authTokenResponse, err := apiClient.RetrieveAuthToken(params)
	if err != nil || *authTokenResponse.Payload.TokenType != "Bearer" {
		return "", err
	}

	return *authTokenResponse.Payload.Token, nil
}
