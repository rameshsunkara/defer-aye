[![GoTemplate](https://img.shields.io/badge/go/template-black?logo=go)](https://github.com/SchwarzIT/go-template)
[![Build Status](https://github.com/rameshsunkara/deferrun/actions/workflows/main.yml/badge.svg)](https://github.com/rameshsunkara/deferrun/actions/workflows/main.yml?query=branch%3Amain)
[![Go Report Card](https://goreportcard.com/badge/github.com/rameshsunkara/deferrun)](https://goreportcard.com/report/github.com/rameshsunkara/deferrun)
[![codecov](https://codecov.io/gh/rameshsunkara/deferrun/branch/main/graph/badge.svg?token=WDFGFOXNNC)](https://codecov.io/gh/rameshsunkara/deferrun)

# Defer Run (deferrun)

A utility to execute given functions when the configured signal(s) fire.
<br>Functions are executed in LIFO order just like `defer` statements.

The default behavior is designed to facilitate executing function(s) on application **termination**. 

Usage Example: [go-rest-api-example](https://github.com/rameshsunkara/go-rest-api-example/blob/main/main.go#L71)

## The Problem

In the below example, methods `closeDBConnection` and `notifyServiceX` will never be executed if the service is terminated.
Try it yourself, run the below code and try to kill it using Ctrl+C.

```go
package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("In main")
	defer closeDBConnection()
	defer notifyServiceX()
	http.ListenAndServe(":8090", nil)
}

func closeDBConnection() {
	fmt.Println("Closing DB Connection")
}

func notifyServiceX() {
	fmt.Println("Hey Service X, I (Service A) am terminating.")
}
```

### Observed Output

```shell
In main
^C
```

## How this module helps ?

With the below code, you can be assured that whatever logic you want to run on your service termination, it will always be executed.

```go
package main

import (
	"fmt"
	"net/http"
	"github.com/rameshsunkara/deferrun"
)

func main() {
	fmt.Println("In main")
	r := deferrun.NewSignalHandler()
	r.OnSignal(closeDBConnection)
	r.OnSignal(notifyServiceX)
	http.ListenAndServe(":8090", nil)
}

func closeDBConnection() {
	fmt.Println("Closing DB Connection")
}

func notifyServiceX() {
	fmt.Println("Hey Service X, I (Service A) am terminating.")
}
```

### Observed Output

```shell
In main
^CHey Service X, I (Service A) am terminating.
Closing DB Connection
```

## Customization | Options

By default, it listens for following signals:
		
	os.Interrupt, syscall.SIGTERM, syscall.SIGINT

If you want to customize the Signals you want to listen, simply pass whichever signals you want:

```go
r := deferrun.NewSignalHandler(syscall.SIGINT, syscall.SIGABRT)
```

## Limitations

By design, `OnSignal` accepts only functions with no parameters or return values as handling the return value should be consumers responsibility and not something that can be generalized.
<br> If the method you want to execute has parameters or return values simply wrap it and then use `OnSignal`.

For example, say `closeDBConnection` can return a error. So wrap `closeDBConnection` in no-arg, no-return value function and pass it to `OnSignal`.

```go
package main

import (
	"fmt"
	"net/http"
	"github.com/rameshsunkara/deferrun"
)

func main() {
	fmt.Println("In main")
	r := deferrun.NewSignalHandler()
	r.OnSignal(func() {
		if err := closeDBConnection(); err!= nil {
			fmt.Println("Too bad, notify everyone as it can cause havoc")
		}
	})
	http.ListenAndServe(":8090", nil)
}

func closeDBConnection() error {
	fmt.Println("Closing DB Connection")
	return nil
}
```

## Development

```bash
make help
```

### Setup

To get your setup up and running the only thing you have to do is

```bash
make all
```

### Test & lint

Run linting

```bash
make lint
```

Run tests

```bash
make test
```
