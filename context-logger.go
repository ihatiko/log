package log

import (
	"context"
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"go.uber.org/zap"
)

type Log struct {
	Fields []zap.Field
}

func WithLog(fields ...zap.Field) *Log {
	return &Log{
		Fields: fields,
	}
}

func (c *Log) WithLog(fields ...zap.Field) *Log {
	c.Fields = append(c.Fields, fields...)
	return c
}

func (c *Log) Context(fields ...zap.Field) *Log {
	c.Fields = append(c.Fields, fields...)
	return c
}
func (c *Log) WithContext(ctx context.Context) *Log {
	span := opentracing.SpanFromContext(ctx)
	if span == nil {
		return c
	}
	spanContext := span.Context()
	jaegerSpanContext, ok := spanContext.(jaeger.SpanContext)
	if ok {
		c.Fields = append(c.Fields, zap.String("spanId", jaegerSpanContext.SpanID().String()))
	}
	return c
}

func WithContext(ctx context.Context) *Log {
	lg := &Log{}
	return lg.WithContext(ctx)
}

func (c *Log) Info(template string, args ...any) {
	lg.Info(fmt.Sprintf(template, args...), c.Fields...)
}

func (c *Log) Warn(template string, args ...any) {
	lg.Warn(fmt.Sprintf(template, args...), c.Fields...)
}

func (c *Log) Error(template string, args ...any) {
	lg.Error(fmt.Sprintf(template, args...), c.Fields...)
}

func (c *Log) DPanic(template string, args ...any) {
	lg.DPanic(fmt.Sprintf(template, args...), c.Fields...)
}

func (c *Log) Panic(template string, args ...any) {
	lg.Panic(fmt.Sprintf(template, args...), c.Fields...)
}

func (c *Log) Fatal(template string, args ...any) {
	lg.Fatal(fmt.Sprintf(template, args...), c.Fields...)
}
