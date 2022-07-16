package controllers

import "fmt"

type USB interface {
	start(types string)
	stop(types string)
}

type Worker struct {
}

func (u *Worker) start(types string) {
	fmt.Println("computer start", types)
}

func (u *Worker) stop(types string) {
	fmt.Println("computer stop", types)
}

func InterfaceTest() {
	var computer USB = &Worker{}
	computer.stop("computer")
	computer.start("computer")

	var worker Worker
	var a1 USB = &worker
	a1.stop("phone")
	a1.start("phone")
}
