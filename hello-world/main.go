package main

import "fmt"

// Hello function return string "Hello, "+name
func Hello(name string) string {
	return "Hello, " + name
}

func main() {
	fmt.Println(Hello("hasan"))
}
