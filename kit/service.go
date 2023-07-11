package kit

import "context"

// Service declares an interface each service must implement
type Service interface {
	// GetCode returns the service unique code
	GetCode() string
	// Init initializes the service
	Init(ctx context.Context) error
	// Start executes all background processes
	Start(ctx context.Context) error
	// Close closes the service
	Close(ctx context.Context)
}

// Adapter common interface for adapters
type Adapter interface {
	// Init initializes adapter
	Init(ctx context.Context, cfg interface{}) error
	// Close closes storage
	Close(ctx context.Context) error
}

// AdapterListener common interface for adapters with listeners
type AdapterListener interface {
	Adapter
	// ListenAsync runs async listening
	ListenAsync(ctx context.Context) error
}
