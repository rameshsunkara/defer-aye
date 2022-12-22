package main

import (
	"fmt"
	"github.com/rameshsunkara/defer-aye/terminator"
	"net/http"
)

func main() {
	fmt.Println("In Main")
	t := terminator.DeferredRun{}
	t.OnTerminate(closeDBConnection)
	t.OnTerminate(closeSomeXResource)
	t.OnTerminate(notifyServiceA)
	/*	defer closeDBConnection()
		defer closeSomeXResource()
		defer notifyServiceA()*/
	http.ListenAndServe(":8090", nil)
}

func closeDBConnection() {
	fmt.Println("closing DB Connection")
}

func closeSomeXResource() {
	fmt.Println("closing resource X")
}

func notifyServiceA() {
	fmt.Println("hello Service A, I am terminating")
}
