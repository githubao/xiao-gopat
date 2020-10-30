// 代理模式：provides a surrogate for an object to control it's actions
// author: baoqiang
// time: 2020/10/29 9:18 下午
package gopat

import "fmt"

// 代理和实体都得实现相同接口
type IObject interface {
	Do(action string)
}

type RealObj struct{}

func (o RealObj) Do(action string) {
	fmt.Printf("real do: %v", action)
}

type ProxyObj struct{
	obj RealObj
}

func (p ProxyObj) Do(action string) {
	fmt.Printf("proxy do: \n")
	p.obj.Do(action)
}

func RunProxy(){
	obj := RealObj{}
	p := ProxyObj{obj: obj}
	p.Do("hello")
}