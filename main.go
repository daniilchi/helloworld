package main

import "fmt"

// test
func main() {
	fmt.Println(HelloWorld())
	fmt.Println(SecondString())
	fmt.Println(ThirdString())
}

func HelloWorld() string {
	return "Hello World!"
}

func SecondString() string {
	return "This is second line!"
}

func ThirdString() string {
	return "This is third line!"
}
