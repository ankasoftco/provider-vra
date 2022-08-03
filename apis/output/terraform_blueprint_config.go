package models 
type TerraformBlueprintConfig struct {

	// Created time
	// Read Only: true
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"createdAt,omitempty"`

	// Created by
	// Read Only: true
	CreatedBy string `json:"createdBy,omitempty"`

	// Blueprint description.
	Description string `json:"description,omitempty"`

	// Object ID
	// Read Only: true
	ID string `json:"id,omitempty"`

	// Blueprint name.
	Name string `json:"name,omitempty"`

	// Org ID
	// Read Only: true
	OrgID string `json:"orgId,omitempty"`

	// Project ID
	ProjectID string `json:"projectId,omitempty"`

	// Project Name
	// Read Only: true
	ProjectName string `json:"projectName,omitempty"`

	// Terraform configuration to blueprint mapping.
	TerraformToBlueprintMapping *TerraformToBlueprintMapping `json:"terraformToBlueprintMapping,omitempty"`

	// Updated time
	// Read Only: true
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updatedAt,omitempty"`

	// Updated by
	// Read Only: true
	UpdatedBy string `json:"updatedBy,omitempty"`

	// Terraform runtime version.
	Version string `json:"version,omitempty"`
}

