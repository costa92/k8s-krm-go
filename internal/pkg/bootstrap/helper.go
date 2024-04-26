package bootstrap

import (
	"context"

	krtlog "github.com/go-kratos/kratos/v2/log"
	"go.opentelemetry.io/otel/trace"
)

// TraceID is the trace id of the request
// nolint: unused
func traceID() krtlog.Valuer {
	return func(ctx context.Context) interface{} {
		if span := trace.SpanContextFromContext(ctx); span.HasTraceID() {
			return span.TraceID().String()
		}
		return ""
	}
}

// SpanID is the span id of the request
// nolint: unused
func spanID() krtlog.Valuer {
	return func(ctx context.Context) interface{} {
		if span := trace.SpanContextFromContext(ctx); span.HasSpanID() {
			return span.SpanID().String()
		}
		return ""
	}
}
