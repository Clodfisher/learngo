package main

import "fmt"

/*
考点：golang的方法集
 解答：
 编译不通过！
一句话：golang的方法集仅仅影响接口实现和方法表达式转化，与通过实例或者指针调用方法无关。
如果是值接收者，实体类型的值和指针都可以实现对应的接口；如果是指针接收者，那么只有类型的指针能够实现对应的接口。
类型的值只能实现值接收者的接口；指向类型的指针，既可以实现值接收者的接口，也可以实现指针接收者的接口。
*/

type Peopel interface {
	Speak(string) string
}

type Student struct {
}

func (stu *Student) Speak(think string) (talk string) {
	if think == "hello" {
		talk = "Hello"
	} else {
		talk = "ni hao"
	}

	return
}

func main() {
	var peo Peopel = Student{}
	fmt.Println(peo.Speak("hello"))
}
