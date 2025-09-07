package domain

import "context"

// UowCallback defines the signature for functions that can be executed by the UnitOfWork.
// It receives a context that may carry transactional information.
type UowCallback func(ctx context.Context) error

// UnitOfWork provides an interface for managing atomic operations that span
// across multiple repositories, ensuring that all changes are either committed
// or rolled back together.
type UnitOfWork interface {
	// Execute runs the given function within a single atomic transaction.
	// If the function returns an error, the transaction is rolled back.
	// Otherwise, the transaction is committed.
	Execute(ctx context.Context, fn UowCallback) error
}
