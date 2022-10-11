package main

import "fmt"

var pl = fmt.Println

type Animal interface {
	SayHello()
	React()
}

type Cat string

func (c Cat) Attack() {
	pl("cat Attacks")
}

func (c Cat) Name() string {
	return string(c)
}

func (c Cat) SayHello() {
	pl("I´m the cat and saying hello")
}

func (c Cat) React() {
	pl("I´m cat, do yo want to play")
}

func main() {
	var myCat Animal
	myCat = Cat("Kitty")
	myCat.SayHello()

	var myCat2 Cat = myCat.(Cat)
	myCat2.Attack()
	pl("Cats Name: ", myCat2.Name())
}
