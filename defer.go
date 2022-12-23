package deferrun

import (
	"os"
	"os/signal"
	"sync"
	"syscall"
)

type SignalHandler interface {
	OnSignal(deferFunc func())
}

func NewSignalHandler(signalsToListen ...os.Signal) SignalHandler {
	if signalsToListen == nil {
		signalsToListen = []os.Signal{os.Interrupt, syscall.SIGTERM, syscall.SIGINT}
	}
	return &deferRun{
		signals: signalsToListen,
	}
}

// deferRun - Implements SignalHandler
type deferRun struct {
	signals       []os.Signal
	deferredFuncs []func()
	runOnce       sync.Once
}

func (dr *deferRun) OnSignal(deferFunc func()) {
	// Prepend as we have to iterate in LIFO order
	dr.deferredFuncs = append([]func(){deferFunc}, dr.deferredFuncs...)
	dr.runOnce.Do(func() {
		dr.run()
	})
}

func (dr *deferRun) run() {
	signalsChan := make(chan os.Signal, 1)
	signal.Notify(signalsChan, dr.signals...)

	go func() {
		<-signalsChan
		for _, f := range dr.deferredFuncs {
			f()
		}
		os.Exit(0)
	}()
}
