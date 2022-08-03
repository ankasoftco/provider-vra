package models 
type CloudAccountAwsRegionEnumerationSpecification struct {

	// Aws Access key ID. Either provide accessKeyId or provide a cloudAccountId of an existing account.
	// Example: ACDC55DB4MFH6ADG75KK
	AccessKeyID string `json:"accessKeyId,omitempty"`

	// Existing cloud account id. Either provide existing cloud account id, or accessKeyId/secretAccessKey credentials pair.
	// Example: b8b7a918-342e-4a53-a3b0-b935da0fe601
	CloudAccountID string `json:"cloudAccountId,omitempty"`

	// Aws Secret Access Key. Either provide secretAccessKey or provide a cloudAccountId of an existing account.
	// Example: gfsScK345sGGaVdds222dasdfDDSSasdfdsa34fS
	SecretAccessKey string `json:"secretAccessKey,omitempty"`
}

