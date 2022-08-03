package models 
type DiskEncryptionSet struct {

	// id
	ID string `json:"id,omitempty"`

	// key
	Key string `json:"key,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// region Id
	RegionID string `json:"regionId,omitempty"`

	// vault
	Vault string `json:"vault,omitempty"`
}

