package main

import "fmt"

//数组  通过类型和长度 最终确定数组的类型 如[5]String
// 数组的长度是固定的
func main() {

	//定义一  长度+初始值
	strings := [5]string{"1", "3", "2", "4", "6"}
	fmt.Println(strings[1])

	//定义二  长度省略，根据初始值数量推断长度
	tow := [...]string{"1", "3", "2", "4", "6"}
	fmt.Println(len(tow))

	//定义三 此种定义不可省略长度
	three := [5]string{1: "a", 3: "c"}
	fmt.Println(three[1])

	//数组循环  传统for循环
	for i := 0; i < len(tow); i++ {
		fmt.Println(i, tow[i])
	}

	// for range
	for i, v := range tow {
		fmt.Printf("索引%d,值为%s\n", i, v)
	}

	//可以range里面未用到到值省略  -- " _  "
	for _, v := range tow {
		fmt.Printf(v)
	}
}
