// 单例模式
// author: baoqiang
// time: 2020/10/29 9:08 下午
package gopat

import (
	"fmt"
	"sync"
)

type singleton map[string]string

var (
	once     sync.Once
	instance singleton
)

func NewSingleton() singleton {
	once.Do(func() {
		instance = make(singleton)
	})
	return instance
}

func RunSingleton() {
	s := NewSingleton()
	s["this"] = "that"
	fmt.Printf("s Got: %s\n", s["this"])

	s2 := NewSingleton()
	fmt.Printf("s2 Got: %s\n", s2["this"])
}
