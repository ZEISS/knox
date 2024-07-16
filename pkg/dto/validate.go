package dto

// Validator is the interface for validating a struct.
type Validator interface {
	// Validate validates the struct.
	Validate() error
}
