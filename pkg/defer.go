package pkg

import (
	"os"
	"os/signal"
	"syscall"
)

var deferFuncs []func()

var SignalsToListen = []os.Signal{os.Interrupt, syscall.SIGTERM, syscall.SIGINT}

func DeferAlways(deferFunc func()) {
	deferFuncs = append(deferFuncs, deferFunc)
}

func init() {
	signals := make(chan os.Signal)
	signal.Notify(signals, SignalsToListen...)

	go func() {
		select {
		case <-signals:
			for _, f := range deferFuncs {
				f()
			}
			os.Exit(0)
		}
	}()

}
