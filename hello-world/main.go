package main

import "fmt"

const englishHelloPrefix = "Hello, "

// Hello function return string "Hello, "+name
func Hello(name string) string {
	if name == "" {
		name = "world"
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("hasan"))
}
