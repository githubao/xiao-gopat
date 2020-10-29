// 构造器
// author: baoqiang
// time: 2020/10/29 7:42 下午
package gopat

import "log"

type Speed float64

const (
	MPH Speed = 1
	KPH       = 1.60934
)

type Color string

const (
	Blue  Color = "blue"
	Green       = "green"
	Red         = "red"
)

type Wheels string

const (
	SportsWheels Wheels = "sports"
	SteelWheels         = "steel"
)

// 定义一些可以被build的参数，传入不同的参数会生成不同的类
type Builder interface {
	SetColor(color Color) Builder
	SetWheels(wheels Wheels) Builder
	SetSpeed(speed Speed) Builder
	Build() Interface
}

// 生成的这些类，都需要实现一些通用的方法
type Interface interface {
	Drive() error
	Stop() error
}

// impl
type Car struct {
	color  Color
	wheels Wheels
	speed  Speed
}

func (c Car) SetColor(color Color) Builder {
	c.color = color
	return c
}
func (c Car) SetWheels(wheels Wheels) Builder {
	c.wheels = wheels
	return c
}
func (c Car) SetSpeed(speed Speed) Builder {
	c.speed = speed
	return c
}
func (c Car) Build() Interface {
	switch c.color {
	case Blue:
		return BlueCar{}
	}
	return nil
}

type BlueCar struct {
}

func (b BlueCar) Drive() error {
	log.Printf("blue drive")
	return nil
}

func (b BlueCar) Stop() error {
	log.Printf("blue stop")
	return nil
}

func NewCar() Car {
	return Car{}
}

func RunBuilder() {
	b := NewCar()
	blueCar := b.SetColor(Blue).Build()
	blueCar.Drive()
	blueCar.Stop()
}
