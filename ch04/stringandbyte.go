package main

import "fmt"

//string是一个不可变的字节序列，所以可以直接转换为字节切片
func main() {
	s := "hello碌碌无为"
	fmt.Println(len(s))
	fmt.Println(s[15])
	bs := []byte(s)
	fmt.Println(bs)

	//二维数组
	aa := [3][3]int{}
	aa[0][0] = 1
	aa[0][1] = 2
	aa[0][2] = 3
	aa[1][0] = 4
	aa[1][1] = 5
	aa[1][2] = 6
	aa[2][0] = 7
	aa[2][1] = 8
	aa[2][2] = 9
	fmt.Println(aa)

}
