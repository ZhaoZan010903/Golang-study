// 声明文件所在的包, 每个go文件必须有归属的包
package main

//引入程序中需要用的包 , 为了使用包下的函数
import "fmt"

// main 主函数  程序的入口
func main() {
	fmt.Print("Hello Golang!") //在控制台打印输出一句话 , 双引号中的内容会原样输出
}
