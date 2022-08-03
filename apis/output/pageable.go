package models 
type Pageable struct {

	// offset
	Offset int64 `json:"offset,omitempty"`

	// page number
	PageNumber int32 `json:"pageNumber,omitempty"`

	// page size
	PageSize int32 `json:"pageSize,omitempty"`

	// paged
	Paged bool `json:"paged,omitempty"`

	// sort
	Sort *Sort `json:"sort,omitempty"`

	// unpaged
	Unpaged bool `json:"unpaged,omitempty"`
}

