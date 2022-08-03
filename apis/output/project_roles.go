package models 
type ProjectRoles struct {
	administratorsField PrincipalRoleAssignment

	membersField PrincipalRoleAssignment

	supervisorsField PrincipalRoleAssignment

	viewersField PrincipalRoleAssignment
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *ProjectRoles) UnmarshalJSON(raw []byte) error {
	var data struct {
		Administrators json.RawMessage `json:"administrators,omitempty"`

		Members json.RawMessage `json:"members,omitempty"`

		Supervisors json.RawMessage `json:"supervisors,omitempty"`

		Viewers json.RawMessage `json:"viewers,omitempty"`
	}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m ProjectRoles) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {
	}{})
	if err != nil {
		return nil, err
	}

