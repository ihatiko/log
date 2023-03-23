package log

import (
	"context"
	"github.com/ihatiko/tracer"
	"github.com/opentracing/opentracing-go"
	"testing"
)

func TestConfig_SetConfiguration(t *testing.T) {
	config := &Config{
		Encoding: "json",
		Level:    "info",
		DevMode:  false,
		Caller:   true,
	}
	config.SetConfiguration("TEST")
	Info("TEST")
}

func TestConfig_Jaeger(t *testing.T) {
	config := &Config{
		Encoding: "json",
		Level:    "info",
		DevMode:  false,
		Caller:   true,
	}
	config.SetConfiguration("TEST")
	cfgJaeger := tracer.Config{LogSpans: false, Host: "localhost:9000"}
	_, err := cfgJaeger.NewTracer("TEST")
	if err != nil {
		Fatal(err)
	}
	span, ctx := opentracing.StartSpanFromContext(context.Background(), "TESTSPAN")
	span.Finish()
	WithLog().
		WithLog().
		WithContext(ctx).
		Info("ERROR")
}
