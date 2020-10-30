// 装饰器模式：Adds behavior to an object, statically or dynamically
// author: baoqiang
// time: 2020/10/29 9:13 下午
package gopat

import (
	"fmt"
	"log"
)

type RealFunc func(i int) int

func Decorator(f RealFunc) RealFunc {
	return func(i int) int {
		log.Printf("start run")
		res := f(i)
		log.Printf("end run")
		return res
	}
}

// 实现方法
func Double(i int) int {
	res := i * 2
	fmt.Printf("Got Res: %v\n", res)
	return res
}

func RunDecorator() {
	deco := Decorator(Double)
	deco(3)
}
