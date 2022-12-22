package main

import (
	"fmt"
	"os"
	"syscall"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOnTerminate(t *testing.T) {
	term := DeferredRun{}
	term.OnTerminate(func() {
		fmt.Println("Test func 1")
	})
	term.OnTerminate(func() {
		fmt.Println("Test func 2")
	})
	assert.EqualValues(t, 3, len(term.Signals))
	assert.EqualValues(t, []os.Signal{os.Interrupt, syscall.SIGTERM, syscall.SIGINT}, term.Signals)
	assert.EqualValues(t, 2, len(term.deferredFuncs))
}

func TestCustomSignals(t *testing.T) {
	term := DeferredRun{
		Signals: []os.Signal{os.Interrupt, syscall.SIGTERM},
	}
	assert.EqualValues(t, 2, len(term.Signals))
	assert.EqualValues(t, []os.Signal{os.Interrupt, syscall.SIGTERM}, term.Signals)
}
