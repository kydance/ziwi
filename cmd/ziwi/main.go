package main

import (
	"fmt"

	"github.com/kydance/ziwi/pkg/strutil"
)

func main() {
	fmt.Printf("strutil.CamelCase(\"hello world\"): %v\n", strutil.CamelCase("hello world"))

	str := "*kyden*"
	fmt.Printf("strutil.UnWarp(str, \"*\"): %v\n", strutil.UnWarp(str, "*"))
	fmt.Println(str)
}
