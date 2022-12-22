package main

import (
	"fmt"
)

func main() {
	fmt.Println("In Main")
	t := DeferredRun{}
	t.OnTerminate(closeDBConnection)
	t.OnTerminate(closeSomeXResource)
	t.OnTerminate(notifyServiceA)
	/*	defer closeDBConnection()
		defer closeSomeXResource()
		defer notifyServiceA()*/
	// http.ListenAndServe(":8090", nil)
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
