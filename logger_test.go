package log

import (
	"testing"
)

func TestConfig_SetConfiguration(t *testing.T) {
	config := &Config{
		Encoding: "json",
		Level:    "info",
		DevMode:  false,
		Caller:   false,
	}
	config.SetConfiguration("TEST")
	Info("TEST")
}
