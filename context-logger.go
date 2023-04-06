package log

import (
	"context"

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
		c.Fields = append(c.Fields, zap.String("traceid", jaegerSpanContext.TraceID().String()))
	}
	return c
}

func WithContext(ctx context.Context) *Log {
	lg := &Log{}
	return lg.WithContext(ctx)
}

func (c *Log) InfoF(template string, args ...any) {
	lg.With(c.Fields...).Sugar().Infof(template, args...)
}
func (c *Log) DebugF(template string, args ...any) {
	lg.With(c.Fields...).Sugar().Debugf(template, args...)
}
func (c *Log) PanicF(template string, args ...any) {
	lg.With(c.Fields...).Sugar().Panicf(template, args...)
}
func (c *Log) ErrorF(template string, args ...any) {
	lg.With(c.Fields...).Sugar().Errorf(template, args...)
}
func (c *Log) DPanicF(template string, args ...any) {
	lg.With(c.Fields...).Sugar().DPanicf(template, args...)
}
func (c *Log) WarnF(template string, args ...any) {
	lg.With(c.Fields...).Sugar().Warnf(template, args...)
}
func (c *Log) FatalF(template string, args ...any) {
	lg.With(c.Fields...).Sugar().Fatalf(template, args...)
}

func (c *Log) Info(args ...any) {
	lg.With(c.Fields...).Sugar().Info(args...)
}
func (c *Log) Debug(args ...any) {
	lg.With(c.Fields...).Sugar().Debug(args...)
}
func (c *Log) Panic(args ...any) {
	lg.With(c.Fields...).Sugar().Panic(args...)
}
func (c *Log) Error(args ...any) {
	lg.With(c.Fields...).Sugar().Error(args...)
}
func (c *Log) DPanic(args ...any) {
	lg.With(c.Fields...).Sugar().DPanic(args...)
}
func (c *Log) Warn(args ...any) {
	lg.With(c.Fields...).Sugar().Warn(args...)
}
func (c *Log) Fatal(args ...any) {
	lg.With(c.Fields...).Sugar().Fatal(args...)
}
