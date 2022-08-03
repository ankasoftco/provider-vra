package models 
type PolicyStats struct {

	// conflict count
	ConflictCount int64 `json:"conflictCount,omitempty"`

	// enforced count
	EnforcedCount int64 `json:"enforcedCount,omitempty"`

	// not enforced count
	NotEnforcedCount int64 `json:"notEnforcedCount,omitempty"`
}

