package deferrun_test

import (
	"fmt"
	"reflect"
	"syscall"
	"testing"

	"github.com/rameshsunkara/deferrun"
	"github.com/stretchr/testify/assert"
)

func TestOnTerminate(t *testing.T) {
	sHandler := deferrun.NewSignalHandler()

	sHandler.OnSignal(func() {
		fmt.Println("Clean func 1")
	})
	sHandler.OnSignal(func() {
		fmt.Println("Clean func 2")
	})

	assert.Implements(t, (*deferrun.SignalHandler)(nil), sHandler)

	sHandlerValue := reflect.Indirect(reflect.ValueOf(sHandler))

	signals := sHandlerValue.FieldByName("signals")
	assert.EqualValues(t, 3, signals.Len())

	deferredFuncs := sHandlerValue.FieldByName("deferredFuncs")
	assert.EqualValues(t, 2, deferredFuncs.Len())
}

func TestCustomSignals(t *testing.T) {
	sHandler := deferrun.NewSignalHandler(syscall.SIGTERM, syscall.SIGINT)
	assert.Implements(t, (*deferrun.SignalHandler)(nil), sHandler)

	sHandlerValue := reflect.Indirect(reflect.ValueOf(sHandler))

	signals := sHandlerValue.FieldByName("signals")
	assert.EqualValues(t, 2, signals.Len())

	deferredFuncs := sHandlerValue.FieldByName("deferredFuncs")
	assert.EqualValues(t, 0, deferredFuncs.Len())
}
