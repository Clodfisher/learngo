package main

import (
	"fmt"
	"sync"
)

/*
下面的迭代会有什么问题？

func (set *threadSafeSet) Iter() <-chan interface{} {
	ch := make(chan interface{})
	go func() {
		set.RLock()

		for elem := range set.s {
			ch <- elem
		}

		close(ch)
		set.RUnlock()

	}()
	return ch
}

输出：
ch0
没有达到迭代器的作用
*/

type threadSafeSet struct {
	sync.Mutex
	s []interface{}
}

func (set *threadSafeSet) Iter() <-chan interface{} {
	ch := make(chan interface{})
	go func() {
		set.Lock()
		for elem := range set.s {
			ch <- elem
		}
		close(ch)
		set.Unlock()
	}()

	return ch
}

func main() {
	th := threadSafeSet{
		s: []interface{}{1, 2},
	}

	v := <-th.Iter()
	fmt.Printf("%s%v\n", "ch", v)
}
