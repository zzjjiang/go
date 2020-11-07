package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var i int = 10

	fmt.Println(i)
	// 类型推断，也可以不用刻意的指定变量的类型
	var j = 11 //可以更简化 j:=11
	fmt.Println(j)
	//定义多个变量
	var (
		l = 12
		y = 13
	)
	fmt.Println(l)
	fmt.Println(y)

	var f32 float32 = 2.2
	var f64 float64 = 10.3456
	fmt.Println("f32 is", f32, ",f64 is", f64)

	var bf bool = false
	var bt bool = true
	fmt.Println("bf is", bf, ",bt is", bt)

	var ls string = "hello"
	var ys string = "li yao"
	fmt.Println("ls + ys ---", ls+ys)
	ls += ys
	fmt.Println("ls += ys ---", ls, ys)

	//零值
	fmt.Println("零值")
	var zi int
	var zf float64
	var zb bool
	var zs string
	fmt.Println(zi, zf, zb, zs)

	//指针
	pi := &i
	fmt.Println(*pi)
	//重新赋值
	i = 1117
	fmt.Println("new i = ", i)

	//常量
	const name = "liyao"

	//iota 常量生成器
	const (
		one = iota + 2
		tow
		three
		four
	)
	fmt.Println(one, tow, three, four)

	//字符串和数字互转
	i2s := strconv.Itoa(i)
	fmt.Println(i2s)
	s2i, err := strconv.Atoi(i2s)
	fmt.Println(s2i, err)

	i2f := float64(i)
	f2i := int(f64)
	fmt.Println(i2f, f2i)

	//小用 string
	sIndex := strings.Index("zoujone", "jone")
	fmt.Println(sIndex)
}
