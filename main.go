package main

import (
	"fmt"
	"os"
)

func main() {
	path := "./test.txt"
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println(err)
	}
	expressions := parser(lexer(data))
	fmt.Println(expressions)
}
