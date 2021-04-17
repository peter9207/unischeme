package main

import "fmt"
import "github.com/peter9207/unischeme/lexer"

func main() {

	p, err := lexer.Parse(`5`)
	if err != nil {
		panic(err)
	}

	fmt.Println(p)
}
