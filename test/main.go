package main

import (
	"fmt"

	"github.com/snowmerak/tson/strcase"
)

func main() {
	s := "FooBar"
	fmt.Println(strcase.PascalToSnake(s))
}
