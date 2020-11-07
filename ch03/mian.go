package main

import "fmt"

func main() {
	//if
	i := 6
	if i >= 10 {
		fmt.Println("i > 10")
	} else if i > 5 && i < 10 {
		fmt.Println("5<i<=10")
	} else {
		fmt.Println("i<=5")
	}

	//switch case自带break 如需要则fallthrough
	switch j := 7; {
	case j > 10:
		fmt.Println("j>10")

	case j > 5 && j <= 10:
		fmt.Println("10>=j>5")
	default:
		fmt.Println("j<=5")
	}

	//for 循环 1 + 100
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println("sum = ", sum)
}
