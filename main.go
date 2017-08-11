package main

import (
	"flag"
	"demo/param"
	"demo/slice"
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
	case "strConv":
		param.StrConv()
	case "timeConv":
		param.TimeConv()
	case "f":
		param.F()
	case "makeAddSuffix":
		addBmp := param.MakeAddSuffix(".bmp")
		addJpeg := param.MakeAddSuffix(".jpeg")
		println(addBmp("file"))
		// returns: file.bmp
		println(addJpeg("file"))
		// returns: file.jpeg
	case "slice1":
		slice.Slice1()
	case "copy":
		slice.Slice2()
	case "editSlice":
		slice.Slice3()

	}
}