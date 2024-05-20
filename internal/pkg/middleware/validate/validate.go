package validate

import (
	"context"
	"github.com/costa92/k8s-krm-go/pkg/api/zerrors"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/middleware"
)

type validator interface {
	Validate() error
}

// IValidator defines methods to implement a custom validator.
type IValidator interface {
	Validate(ctx context.Context, rq any) error
}

// Validator is a validator middleware.
func Validator(vd IValidator) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, rq any) (reply any, err error) {
			if v, ok := rq.(validator); ok {
				// Kratos validation method
				if err := v.Validate(); err != nil {
					if se := new(errors.Error); errors.As(err, &se) {
						return nil, se
					}

					return nil, zerrors.ErrorInvalidParameter(err.Error()).WithCause(err)
				}
			}

			// Custom validation, specific to the API interface
			if err := vd.Validate(ctx, rq); err != nil {
				if se := new(errors.Error); errors.As(err, &se) {
					return nil, se
				}

				return nil, zerrors.ErrorInvalidParameter(err.Error()).WithCause(err)
			}

			return handler(ctx, rq)
		}
	}
}
