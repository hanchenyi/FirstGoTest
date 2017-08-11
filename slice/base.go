package slice

import "fmt"

func Slice1(){
	var slice1 = [5](int){1,2,3,4,5}
	var slice2 = [](int){1,2,3}
	var slice3 = [](string){3:"tom", 4:"jerry"}

	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)

	for k,v := range slice1 {
		v *=2
		fmt.Printf("The slice1 at %d is %d\n",k,v)
	}

	var slice4 = slice2[1:3]
	fmt.Println(slice4)
}

//[1 2 3 4 5]
//[1 2 3]
//[   tom jerry]
//The slice1 at 0 is 2
//The slice1 at 1 is 4
//The slice1 at 2 is 6
//The slice1 at 3 is 8
//The slice1 at 4 is 10
//[2 3]

func Slice2(){
	sl_from := []int{1, 2, 3}
	sl_to := make([]int, 10)

	n := copy(sl_to, sl_from)
	fmt.Println(sl_to)
	fmt.Printf("Copied %d elements\n", n) // n == 3

	sl3 := []int{1, 2, 3}
	sl3 = append(sl3, 4, 5, 6)
	fmt.Println(sl3)
}

//returns the number of elements copied, which will be the minimum of len(src) and len(dst).
//[1 2 3 0 0 0 0 0 0 0]
//Copied 3 elements
//[1 2 3 4 5 6]

//修改字符串中的某个字符
//因此，您必须先将字符串转换成字节数组，然后再通过修改数组中的元素值来达到修改字符串的目的，最后将字节数组转换回字符串格式。
//例如，将字符串 "hello" 转换为 "cello"：
func Slice3(){
	s := "hello"
	fmt.Printf("original string is %s\n",s)
	c := []byte(s)
	c[0] = 'c'
	s2 := string(c) // s2 == "cello"
	fmt.Printf("replaced string is %s\n",s2)
}