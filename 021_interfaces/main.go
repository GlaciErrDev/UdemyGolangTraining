package main

import "fmt"

type S struct {
	val int
}

func (s S) Get() int {
	return s.val
}

func (s *S) Set(v int) {
	s.val = v
}

type Foo interface {
	Get() int
	Set(v int)
}

func Bar(f Foo) {
	fmt.Println(f.Get())
	f.Set(5)
	fmt.Println(f.Get())
}

func main() {
	s := S{}
	Bar(&s) // passing address !!!
	//Bar(s) // S does not implement Foo (Set method has pointer receiver)
}
