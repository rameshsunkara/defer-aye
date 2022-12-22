[![GoTemplate](https://img.shields.io/badge/go/template-black?logo=go)](https://github.com/SchwarzIT/go-template)

# Defer Run (deferrun)

Provides a utility to ensure given function(s) are always executed. 

## The Problem

In the below example, methods `closeDBConnection` and `notifyServiceX` will never be executed if the service is terminated.
Try it yourself, run the below code and try to kill it using Ctrl+C or other preferred means.

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

## How this module helps ?

With the below code, you can be assured that whatever logic you want to run on your service termination, it will always be executed.

```go

```

## Customization & Options

By default, it listens for the following signals:
		
	os.Interrupt, syscall.SIGTERM, syscall.SIGINT

If you want to customize the Signals, simply pass the signals you want to listen on:

```go

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
