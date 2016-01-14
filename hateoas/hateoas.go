package hateoas

// Error is an HATEOAS error
type Error struct {
	ID     int    `json:"id,omitempty"`
	Status int    `json:"status,omitempty"`
	Title  string `json:"title,omitempty"`
	Detail string `json:"detail,omitempty"`
}

// Errors is a slice of HATEOAS errors
type Errors []Error
