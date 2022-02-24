package main

import (
	echoServer "jasonblanchard/depfootprint/pkg/echo"
)

func main() {

	e, _ := echoServer.MakeEcho()
	e.Logger.Fatal(e.Start(":1323"))
}
