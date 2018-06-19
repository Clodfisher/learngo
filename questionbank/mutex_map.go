package main

import (
	"fmt"
	"sync"
	"time"
)

/*
下面的代码有什么问题?
考点：map线程安全
解答：
可能会出现fatal error: concurrent map read and map write. 修改一下看看效果
func (ua *UserAges) Get(name string) int {
    ua.Lock()
    defer ua.Unlock()
    if age, ok := ua.ages[name]; ok {
        return age
    }
    return -1
}
*/

/*
在go中，匿名成员的方法 会默认被提升为类型本身的方法
*/
type UserAges struct {
	ages map[string]int
	int
	sync.Mutex
}

func (ua *UserAges) Add(name string, age int) {
	ua.Lock()
	defer ua.Unlock()
	ua.ages[name] = age
}

func (ua *UserAges) Get(name string) int {
	if age, ok := ua.ages[name]; ok {
		return age
	}
	return -1
}

func main() {
	s := make([]string, 0)
	s = append(s, "tom", "cat", "dog", "bird", "fish", "pig")
	ua := UserAges{
		ages: make(map[string]int),
	}

	var ua1 UserAges
	ua1.ages = make(map[string]int)
	ua1.int = 10

	for i, name := range s {
		go ua.Add(name, i)
	}

	time.Sleep(time.Second * 1)
	age := ua.Get("fish")
	fmt.Println("fish -> ", age)

}
