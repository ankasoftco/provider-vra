package models 
type ProjectPrincipalsAssignment struct {
	modifyField []PrincipalRole

	removeField []PrincipalRole
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *ProjectPrincipalsAssignment) UnmarshalJSON(raw []byte) error {
	var data struct {
		Modify json.RawMessage `json:"modify"`

		Remove json.RawMessage `json:"remove"`
	}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m ProjectPrincipalsAssignment) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {
	}{})
	if err != nil {
		return nil, err
	}

