package main

import "fmt"

func main() {
	// 1.变量的声明
	var age int
	// 2.变量的赋值
	age = 18
	// 3.变量的使用
	fmt.Println("age=", age)

	//声明和赋值合成一句
	var age2 int = 12
	fmt.Println("age2=", age2)

	/*
		变量的重复定义会报错
		# command-line-arguments
		.\demo01.go:17:6: age redeclared in this block
		.\demo01.go:7:6: other declaration of age

		var age int = 20
		fmt.Println(age)
	*/

	//不能给不匹配的类型
	/*
		var num int = 12.56
		fmt.Println("num = ", num)
	*/
}
