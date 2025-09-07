package application

import (
	"context"
	"github.com/kevinsoras/employee-management/shared/domain"
)

// UseCase defines a generic interface for any use case following the command pattern.
// TRequest is the input data type (e.g., a request DTO).
// TResponse is the output data type (e.g., a response DTO).
type UseCase[TRequest any, TResponse any] interface {
	Execute(ctx context.Context, req TRequest) (TResponse, error)
}

// TransactionalDecorator is a generic decorator that wraps a UseCase to run it within a transaction.
type TransactionalDecorator[TRequest any, TResponse any] struct {
	useCase UseCase[TRequest, TResponse]
	uow     domain.UnitOfWork
}

// NewTransactionalDecorator creates a new transactional decorator.
func NewTransactionalDecorator[TRequest any, TResponse any](useCase UseCase[TRequest, TResponse], uow domain.UnitOfWork) UseCase[TRequest, TResponse] {
	return &TransactionalDecorator[TRequest, TResponse]{
		useCase: useCase,
		uow:     uow,
	}
}

// Execute wraps the execution of the decorated use case within a database transaction.
func (d *TransactionalDecorator[TRequest, TResponse]) Execute(ctx context.Context, req TRequest) (TResponse, error) {
	var response TResponse
	var err error

	err = d.uow.Execute(ctx, func(txCtx context.Context) error {
		response, err = d.useCase.Execute(txCtx, req)
		return err
	})

	if err != nil {
		var zero TResponse // Return the zero value for the response type on error.
		return zero, err
	}

	return response, nil
}
