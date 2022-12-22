package terminate

import (
	"os"
	"os/signal"
	"sync"
	"syscall"
)

type DeferredRun struct {
	Signals       []os.Signal
	deferredFuncs []any
	runOnce       sync.Once
}

func (dr *DeferredRun) OnTerminate(deferFunc any) {
	dr.deferredFuncs = append([]any{deferFunc}, dr.deferredFuncs...)
	dr.runOnce.Do(func() {
		dr.run()
	})
}

func (dr *DeferredRun) run() {
	signalsChan := make(chan os.Signal, 1)

	if dr.Signals == nil {
		dr.Signals = []os.Signal{os.Interrupt, syscall.SIGTERM, syscall.SIGINT}
	}

	signal.Notify(signalsChan, dr.Signals...)

	go func() {
		<-signalsChan
		for _, f := range dr.deferredFuncs {
			f.(func())()
		}
		os.Exit(0)
	}()
}
