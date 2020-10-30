// 对象池：instantiates and maintains a group of objects instances of the same type
// author: baoqiang
// time: 2020/10/29 8:33 下午
package gopat

import (
	"log"
	"time"
)

var workChan chan int

// 具体干活的对象
type Worker struct {
	id int
}

func (w Worker) Do() {
	for {
		select {
		case i := <-workChan:
			log.Printf("Worker(%d) Do Task-%v", w.id, i)
		}
		if len(workChan) == 0 {
			break
		}
	}
}

// 新建一个对象池
type Pool []*Worker

func NewPool(num int) Pool {
	var p []*Worker
	for i := 0; i < num; i++ {
		p = append(p, &Worker{
			id: i,
		})
	}
	return p
}

// 干活
func work() {
	p := NewPool(2)

	for _, w := range p {
		go w.Do()
	}

}

func RunPool() {
	work()

	workChan = make(chan int, 3)

	for i := 0; i < 10; i++ {
		workChan <- i
	}

	close(workChan)
	time.Sleep(time.Second * 5)
}
