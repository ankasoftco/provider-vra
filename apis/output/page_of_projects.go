package models 
type PageOfProjects struct {
	contentField []Project

	// empty
	Empty bool `json:"empty,omitempty"`

	// first
	First bool `json:"first,omitempty"`

	// last
	Last bool `json:"last,omitempty"`

	// number
	Number int32 `json:"number,omitempty"`

	// number of elements
	NumberOfElements int32 `json:"numberOfElements,omitempty"`

	// pageable
	Pageable *Pageable `json:"pageable,omitempty"`

	// size
	Size int32 `json:"size,omitempty"`

	// sort
	Sort *Sort `json:"sort,omitempty"`

	// total elements
	TotalElements int64 `json:"totalElements,omitempty"`

	// total pages
	TotalPages int32 `json:"totalPages,omitempty"`
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *PageOfProjects) UnmarshalJSON(raw []byte) error {
	var data struct {
		Content json.RawMessage `json:"content"`

		Empty bool `json:"empty,omitempty"`

		First bool `json:"first,omitempty"`

		Last bool `json:"last,omitempty"`

		Number int32 `json:"number,omitempty"`

		NumberOfElements int32 `json:"numberOfElements,omitempty"`

		Pageable *Pageable `json:"pageable,omitempty"`

		Size int32 `json:"size,omitempty"`

		Sort *Sort `json:"sort,omitempty"`

		TotalElements int64 `json:"totalElements,omitempty"`

		TotalPages int32 `json:"totalPages,omitempty"`
	}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m PageOfProjects) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {
		Empty bool `json:"empty,omitempty"`

		First bool `json:"first,omitempty"`

		Last bool `json:"last,omitempty"`

		Number int32 `json:"number,omitempty"`

		NumberOfElements int32 `json:"numberOfElements,omitempty"`

		Pageable *Pageable `json:"pageable,omitempty"`

		Size int32 `json:"size,omitempty"`

		Sort *Sort `json:"sort,omitempty"`

		TotalElements int64 `json:"totalElements,omitempty"`

		TotalPages int32 `json:"totalPages,omitempty"`
	}{

		Empty: m.Empty,

		First: m.First,

		Last: m.Last,

		Number: m.Number,

		NumberOfElements: m.NumberOfElements,

		Pageable: m.Pageable,

		Size: m.Size,

		Sort: m.Sort,

		TotalElements: m.TotalElements,

		TotalPages: m.TotalPages,
	})
	if err != nil {
		return nil, err
	}

