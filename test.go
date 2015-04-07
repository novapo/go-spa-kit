package main

import (
	"fmt"
	"strconv"
)

type Foo interface{
	String() string
}

type bar struct {
	a int	
}

func (b bar) String() string {
	return strconv.Itoa(b.a)
}


func main() {
	b := bar{5}
	Print(b)
}


func Print(d Foo) {
	fmt.Println(d.String())
}
