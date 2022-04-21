package main

import (
	"fmt"

	"github.com/snowmerak/tson/lexer/jsvalues"
)

func main() {
	t := "{[{\"a\":1,\"b\":2},{\"a\":3,\"b\":4}, 1, 2, 3, true, null, [1, 2, 3, [1, 2, 3]]]}"
	_, v, _, err := jsvalues.FindObject([]byte(t))
	if err != nil {
		panic(err)
	}
	fmt.Println(string(v))
}
