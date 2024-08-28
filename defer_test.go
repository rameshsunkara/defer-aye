package deferrun_test

import (
	"reflect"
	"syscall"
	"testing"

	"github.com/rameshsunkara/deferrun"
)

func TestNewSignalHandler(t *testing.T) {
	sHandler := deferrun.NewSignalHandler()

	sHandlerValue := reflect.Indirect(reflect.ValueOf(sHandler))
	signals := sHandlerValue.FieldByName("signals")
	if 3 != signals.Len() {
		t.Errorf("Expected 3 Signal but got: %d", signals.Len())
	}

	deferredFuncs := sHandlerValue.FieldByName("deferredFuncs")
	if 0 != deferredFuncs.Len() {
		t.Errorf("Expected 0 deferred functions but got: %d", deferredFuncs.Len())
	}
}

func TestCustomSignals(t *testing.T) {
	sHandler := deferrun.NewSignalHandler(syscall.SIGTERM, syscall.SIGINT)

	sHandlerValue := reflect.Indirect(reflect.ValueOf(sHandler))

	signals := sHandlerValue.FieldByName("signals")
	if 2 != signals.Len() {
		t.Errorf("Expected 2 Signal but got: %d", signals.Len())
	}
}

func TestOnSignal(t *testing.T) {
	sHandler := deferrun.NewSignalHandler()

	sHandler.OnSignal(func() {
		t.Logf("method 1 executed")
	})
	sHandler.OnSignal(func() {
		t.Logf("method 2 executed")
	})

	sHandlerValue := reflect.Indirect(reflect.ValueOf(sHandler))
	deferredFuncs := sHandlerValue.FieldByName("deferredFuncs")
	if 2 != deferredFuncs.Len() {
		t.Errorf("Expected 2 deferred functions but got: %d", deferredFuncs.Len())
	}

	// syscall.Kill(syscall.Getpid(), syscall.SIGINT)
}
