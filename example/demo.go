package example

import "fmt"

type Foo interface {
	Bar(a string, b string)
}

type foo struct {
}

func (f *foo) Bar(a string, b string) {
	fmt.Println(a, b)
}
