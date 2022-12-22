package terminator

import (
	"os"
	"os/signal"
	"sync"
	"syscall"
)

var runOnce sync.Once

var deferFuncs []func()

var SignalToListen = []os.Signal{os.Interrupt, syscall.SIGTERM, syscall.SIGINT}

func OnTerminate(deferFunc func()) {
	deferFuncs = append(deferFuncs, deferFunc)
	runOnce.Do(func() {
		run()
	})
}

func run() {
	signalsChan := make(chan os.Signal, 1)
	signal.Notify(signalsChan, SignalToListen...)

	go func() {
		<-signalsChan
		for _, f := range deferFuncs {
			f()
		}
		os.Exit(0)
	}()
}
