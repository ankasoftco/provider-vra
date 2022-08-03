package models 
type DiskSnapshotSpecification struct {

	// A human-friendly description.
	Description string `json:"description,omitempty"`

	// A human-friendly name used as an identifier in APIs that support this option.
	Name string `json:"name,omitempty"`

	// Cloud specific snapshot properties supplied in as name value pairs
	// Example: {\"incremental\": \"true\",
	SnapshotProperties map[string]string `json:"snapshotProperties,omitempty"`

	// A set of tag keys and optional values that have to be set on the snapshot in the cloud. Currently supported for Azure Snapshots
	// Example: [ { \"key\" : \"env\", \"value\": \"dev\" } ]
	Tags []*Tag `json:"tags"`
}

