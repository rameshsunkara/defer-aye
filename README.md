[![GoTemplate](https://img.shields.io/badge/go/template-black?logo=go)](https://github.com/SchwarzIT/go-template)

# Defer Run (deferrun)

Provides a utility to ensure given function(s) are always executed. 
Given functions are executed in LIFO (Last In First Out) order just like `defer`.

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
func main() {
	fmt.Println("In main")
	r := deferrun.DeferRun{}
	r.OnTerminate(closeDBConnection)
	r.OnTerminate(notifyServiceX)
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

By default, it listens for the following signals:
		
	os.Interrupt, syscall.SIGTERM, syscall.SIGINT

If you want to customize the Signals, simply pass the signals you want to listen on:

```go
	r := deferrun.DeferRun{
		Signals: []os.Signal{syscall.SIGINT, syscall.SIGABRT},
	}
```

```bash
make help
```

## Setup

To get your setup up and running the only thing you have to do is

```bash
make all
```

## Test & lint

Run linting

```bash
make lint
```

Run tests

```bash
make test
```
