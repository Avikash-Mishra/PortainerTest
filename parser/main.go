package main

import (
	"flag"
	"fmt"

	"parser/parser"
)

func main()  {
	var input string
	flag.StringVar(&input, "parse", "", "a string of brackets to parse")
	flag.Parse()
	fmt.Printf("%t\n", parser.ParseStructure(input))
}