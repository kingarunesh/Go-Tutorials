package main

import "fmt"

var a int
var b int = 1
var c = 2

func main() {

	fmt.Println("----- Main Function -----")
	fmt.Println("a : ", a)
	fmt.Println("b : ", b)
	fmt.Println("c : ", c)

	// 		test function called
	test()

}

func test() {
	fmt.Println("----- Test Function -----")
	fmt.Println("a : ", a)
	fmt.Println("b : ", b)
	fmt.Println("c : ", c)

	a = 9
	b = 99
	c = 999

	fmt.Println("a : ", a)
	fmt.Println("b : ", b)
	fmt.Println("c : ", c)
}
