/*
Copyright 2021 Upbound Inc.
*/

package clients

import (
	"context"
	"encoding/json"

	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/upbound/upjet/pkg/terraform"

	"github.com/ankasoftco/upjet-provider-vra/apis/v1beta1"
)

const (
	// error messages
	errNoProviderConfig     = "no providerConfigRef provided"
	errGetProviderConfig    = "cannot get referenced ProviderConfig"
	errTrackUsage           = "cannot track ProviderConfig usage"
	errExtractCredentials   = "cannot extract credentials"
	errUnmarshalCredentials = "cannot unmarshal vRA credentials as JSON"
)

const (
	keyURL          = "url"
	keyRefreshToken = "refresh_token"

	envURL          = "VRA_URL"
	envRefreshToken = "VRA_REFRESH_TOKEN"
)

// TerraformSetupBuilder builds Terraform a terraform.SetupFn function which
// returns Terraform provider setup configuration
func TerraformSetupBuilder(version, providerSource, providerVersion string) terraform.SetupFn {
	return func(ctx context.Context, client client.Client, mg resource.Managed) (terraform.Setup, error) {
		ps := terraform.Setup{
			Version: version,
			Requirement: terraform.ProviderRequirement{
				Source:  providerSource,
				Version: providerVersion,
			},
		}

		configRef := mg.GetProviderConfigReference()
		if configRef == nil {
			return ps, errors.New(errNoProviderConfig)
		}
		pc := &v1beta1.ProviderConfig{}
		if err := client.Get(ctx, types.NamespacedName{Name: configRef.Name}, pc); err != nil {
			return ps, errors.Wrap(err, errGetProviderConfig)
		}

		t := resource.NewProviderConfigUsageTracker(client, &v1beta1.ProviderConfigUsage{})
		if err := t.Track(ctx, mg); err != nil {
			return ps, errors.Wrap(err, errTrackUsage)
		}

		data, err := resource.CommonCredentialExtractor(ctx, pc.Spec.Credentials.Source, client, pc.Spec.Credentials.CommonCredentialSelectors)
		if err != nil {
			return ps, errors.Wrap(err, errExtractCredentials)
		}

		// Configuration is a map of provider configuration values.
		vraCreds := map[string]string{}
		if err := json.Unmarshal(data, &vraCreds); err != nil {
			return ps, errors.Wrap(err, errUnmarshalCredentials)
		}

		ps.Configuration = map[string]interface{}{}
		if v, ok := vraCreds[keyURL]; ok {
			ps.Configuration[keyURL] = v
		}
		if v, ok := vraCreds[keyRefreshToken]; ok {
			ps.Configuration[keyRefreshToken] = v
		}
		
		// Set credentials in Terraform provider environment.
		/*ps.Env = []string{
			fmt.Sprintf("%s=%s", envURL, vraCreds[keyURL]),
			fmt.Sprintf("%s=%s", envRefreshToken, vraCreds[keyRefreshToken]),
		}*/

		// Set credentials in Terraform provider configuration.
		/*ps.Configuration = map[string]any{
			"username": creds["username"],
			"password": creds["password"],
		}*/
		return ps, nil
	}
}
