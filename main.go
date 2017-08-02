package main

import (
	"flag"
	"demo/param"
)

var (
	testFlag = flag.String("testFlag","hello", "chose test case")
)

func main() {
	flag.Parse()

	switch *testFlag {
	case "hello":
		param.Hello()
	case "bye":
		param.Bye()
	}
}