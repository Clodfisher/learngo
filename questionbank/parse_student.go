package main

import "fmt"

/*
以下代码有什么问题，说明原因?

考点：foreach
解答：
这样的写法初学者经常会遇到的，很危险！ 与Java的foreach一样，都是使用副本的方式。
所以m[stu.Name]=&stu实际上一致指向同一个指针， 最终该指针的值为遍历的最后一个struct的值拷贝。
就像修改切片元素的属性：
*/

type student struct {
	Name string
	Age  int
}

func parse_student() {
	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	for _, stu := range stus {
		m[stu.Name] = &stu
	}
	fmt.Printf("%#v\n", m)

	for k, v := range m {
		fmt.Println(k, "=>", v.Name, v.Age)
	}

	// 正确
	for i := 0; i < len(stus); i++ {
		m[stus[i].Name] = &stus[i]
	}
	for k, v := range m {
		fmt.Println(k, "=>", v.Name)
	}
}

func main() {
	parse_student()
}
