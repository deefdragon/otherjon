package main

import (
	"fmt"
	"io"
	"time"
)

var stuff map[int]string = make(map[int]string)
var arr []string = make([]string, 0)

func main() {

	h := &HoldieThingine[string]{
		SomeThing: "asdf",
	}

	DoTheDooer(h)

}

func DoTheDooer(thing Doer, p2 io.WriteCloser) {
	thing.Do()
}

type HoldieThingine[T any] struct {
	SomeThing T

	someCreationTime time.Time
	//some other stuff.
}

func (h *HoldieThingine[T]) Do() {
	fmt.Println(h.SomeThing)
}

type Doer interface {
	Do()
}
