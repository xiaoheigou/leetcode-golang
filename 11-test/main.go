package main

import (
	"fmt"
	"unsafe"
)

type person struct {
	name string
	age  int
}

func main() {
	p := person{
		"asdfssadfasfdsfs爱福家暗示法大大沙发",
		12,
	}
	// s := new(person)
	var s person
	p.name = "hahah"
	s.name = "asdf"
	fmt.Printf("%T\n", p)
	fmt.Printf("%T----%q", s, s.name)
	fmt.Printf("%v", unsafe.Sizeof(p))
}
