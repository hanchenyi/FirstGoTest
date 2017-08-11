package param

import (
	"fmt"
	"strconv"
	"time"
	"strings"
)

func Hello() {
	var str = "hello world"
	println(str)
	fmt.Println(len(str))
}

func Bye() {
	println("bye bye")
}


//strconv.Atoi(s string) (i int, err error) 将字符串转换为 int 型。
//strconv.ParseFloat(s string, bitSize int) (f float64, err error) 将字符串转换为 float64 型。

func StrConv() {
	var orig string = "666"
	var an int
	var newS string
	fmt.Printf("The size of ints is: %d\n", strconv.IntSize)
	an, _ = strconv.Atoi(orig)
	fmt.Printf("The integer is: %d\n", an)
	an = an + 5
	newS = strconv.Itoa(an)
	fmt.Printf("The new string is: %s\n", newS)
}

// Duration 类型表示两个连续时刻所相差的纳秒数，类型为 int64
var week time.Duration
func TimeConv() {
	t := time.Now()
	fmt.Println(t) // e.g. Wed Dec 21 09:52:14 +0100 RST 2011
	fmt.Printf("%02d.%02d.%4d\n", t.Day(), t.Month(), t.Year())
	// 21.12.2011
	//协调世界时（Coordinated Universal Time)
	t = time.Now().UTC()
	fmt.Println(t) // Wed Dec 21 08:52:14 +0000 UTC 2011
	fmt.Println(time.Now()) // Wed Dec 21 09:52:14 +0100 RST 2011
	// calculating times:
	week = 60 * 60 * 24 * 7 * 1e9 // must be in nanosec
	week_from_now := t.Add(week)
	fmt.Println(week_from_now) // Wed Dec 28 08:52:14 +0000 UTC 2011
	// formatting times:
	fmt.Println(t.Format(time.RFC822)) // 21 Dec 11 0852 UTC
	fmt.Println(t.Format(time.ANSIC)) // Wed Dec 21 08:56:34 2011
	fmt.Println(t.Format("02 Jan 2006 15:04")) // 21 Dec 2011 08:52
	s := t.Format("20060102")
	fmt.Println(t, "=>", s)
	// Wed Dec 21 08:52:14 +0000 UTC 2011 => 20111221
}

//闭包, 匿名函数
//它们被允许调用定义在其它环境下的变量。闭包可使得某个函数捕捉到一些外部状态，例如：函数被创建时的状态。
//另一种表示方式为：一个闭包继承了函数所声明时的作用域。这种状态（作用域内的变量）都被共享到闭包的环境中，
//因此这些变量可以在闭包中被操作，直到被销毁
func F() {
	for i := 0; i < 4; i++ {
		g := func(i int) { fmt.Printf("%d ", i) } //此例子中只是为了演示匿名函数可分配不同的内存地址，在现实开发中，不应该把该部分信息放置到循环中。
		g(i)
		fmt.Printf(" - g is of type %T and has value %v\n", g, g)
	}
}

// 返回值是一个函数，函数的入参是在最后入值的。
func MakeAddSuffix(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}