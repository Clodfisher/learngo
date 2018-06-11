package main

/*
请写出以下输入内容 ?

考点：make默认值和append
解答：
make初始化是由默认值，此处默认值为0

[0 0 0 0 0 1 2 3]
*/
import "fmt"

func main() {
	s := make([]int, 5)
	s = append(s, 1, 2, 3)
	fmt.Println(s)
}
