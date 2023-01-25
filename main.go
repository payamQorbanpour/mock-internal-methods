package main

import "fmt"

type Operation struct{}

type SomeInterface interface {
	DoSomething(number int) (bool, error)
	ParrentFunction()
}

func main() {
	o := &Operation{}
	o.ParrentFunction()
}

func (o *Operation) ParrentFunction() bool {
	done := o.DoSomething(321)
	fmt.Println("Parrent function success ", done)
	return done
}

func (o *Operation) DoSomething(number int) bool {
	fmt.Println("Here in the fucntion")
	return true
}
