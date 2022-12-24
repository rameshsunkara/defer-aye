package deferrun_test

import (
	"fmt"
	"reflect"
	"syscall"
	"testing"

	"github.com/rameshsunkara/deferrun"
)

func TestOnTerminate(t *testing.T) {
	sHandler := deferrun.NewSignalHandler()

	sHandler.OnSignal(func() {
		fmt.Println("Clean func 1")
	})
	sHandler.OnSignal(func() {
		fmt.Println("Clean func 2")
	})

	sHandlerValue := reflect.Indirect(reflect.ValueOf(sHandler))

	signals := sHandlerValue.FieldByName("signals")
	if 3 != signals.Len() {
		t.Errorf("Expected 3 Signal but got: %d", signals.Len())
	}

	deferredFuncs := sHandlerValue.FieldByName("deferredFuncs")
	if 2 != deferredFuncs.Len() {
		t.Errorf("Expected 2 deferred functions but got: %d", deferredFuncs.Len())
	}
}

func TestCustomSignals(t *testing.T) {
	sHandler := deferrun.NewSignalHandler(syscall.SIGTERM, syscall.SIGINT)

	sHandlerValue := reflect.Indirect(reflect.ValueOf(sHandler))

	signals := sHandlerValue.FieldByName("signals")
	if 2 != signals.Len() {
		t.Errorf("Expected 2 Signal but got: %d", signals.Len())
	}

	deferredFuncs := sHandlerValue.FieldByName("deferredFuncs")
	if 0 != deferredFuncs.Len() {
		t.Errorf("Expected 0 deferred functions but got: %d", deferredFuncs.Len())
	}
}
