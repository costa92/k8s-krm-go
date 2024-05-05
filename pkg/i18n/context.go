package i18n

import "context"

type translator struct {
}

// NewContext returns a new context with the I18n instance.
func NewContext(ctx context.Context, i *I18n) context.Context {
	return context.WithValue(ctx, translator{}, i)
}

// FromContext returns the I18n instance from the context.
func FromContext(ctx context.Context) *I18n {
	if i, ok := ctx.Value(translator{}).(*I18n); ok {
		return i
	}
	return New()
}
