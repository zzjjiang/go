package main

import "fmt"

//切片就是一个动态数组   底层是数组   数组的任意分割就可以得到一个切片
func main() {
	// 原始数组
	array := [5]string{"1", "a", "c", "3", "g"}
	// 基于数组生成切片 slice:=array[start:end]  包含索引start，不包含索引end
	slice := array[2:5]
	for i, v := range slice {
		fmt.Printf("索引为%d,值为%s\n", i, v)
		/*
			左闭右开
			索引为0,值为c
			索引为1,值为3
			索引为2,值为g
		*/
	}

	//方式一 切片创建   类型+长度+容量  容量 > 长度（切片内元素个数）
	slice1 := make([]string, 4, 8)
	fmt.Println(len(slice1), cap(slice1))
	//方式二 通过字面量创建
	slice2 := []string{"qa", "11", "17", "yy"}
	fmt.Println(len(slice2), cap(slice2))

	//切面元素追加 append会自动处理切片容量不足需要扩容到问题
	slice2 = append(slice2, "f")
	fmt.Println(len(slice2))
}
