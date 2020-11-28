package main

import "fmt"

// 1、map(映射)  K_V key必须是相同类型  无序
// 2、结构为map[k]v
// 3、map的大小 没有容量只有长度
func main() {

	//创建方式一  make
	myMap := make(map[string]int)
	myMap["jone"] = 20
	age := myMap["jone"]
	//创建方式二 字面量
	myMap1 := map[string]int{"yy": 22}
	fmt.Println(age, "----", myMap1["yy"])

	//如果获取的key不存在时，会返回value类型对应的0值，所以我们要先判断key存不存在
	//map []操作符可以返回两个值 第一个为对应的value，第二标记key是否存在，存在则为true
	myMap2 := map[string]int{"ly": 22}
	age, ok := myMap2["ly1"]
	if ok {
		fmt.Println(age)
	}

	//删除 delete(map,"key)
}
