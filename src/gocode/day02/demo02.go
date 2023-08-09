package main

import "fmt"

// 全局变量: 定义在函数外的变量
// 全局与局部变量重名会使用局部变量
var n7 = 100
var n8 = 9.7

var (
	n9  = 100
	n10 = "netty"
)

func main() {
	//定义在{}中的变量叫局部变量

	//第一种：变量的使用方式
	var num int = 18
	fmt.Println(num) //18

	//第二种：指定变量的类型，但不赋值，使用默认值打印
	var num2 int
	fmt.Println(num2) //0

	//第三种：如果没有写变量的类型，那么会根据=后面的值进行判定变量的类型 (自动类型推断)
	var num3 = "tom"
	println(num3)

	//第四种：省略var关键字,注意 := 不能写为=
	sex := "男"
	fmt.Println(sex) //男

	fmt.Println("---------------------------------------")

	// 声明多个变量
	var n1, n2, n3 int
	fmt.Println(n1, n2, n3)

	name1 := true
	println(name1)
	fmt.Println(name1)

	// 声明多个变量且赋值
	var n4, name, n5 = 10, "jack", 7.8
	fmt.Println(n4, name, n5)

	n6, height := 6.9, 100.6
	fmt.Println(n6, height)

	fmt.Println(n7, n8)
	fmt.Println(n9, n10)
}
