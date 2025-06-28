package main

import "fmt"

const (
	RESET      = "\u001B[0m"
	UNDER_LINE = "\u001B[21m"
	RED        = "\u001B[31m"
	GREEN      = "\u001B[32m"
	YELLOW     = "\u001B[33m"
	BLUE       = "\u001B[34m"
)

func main() {
	fmt.Println(BLUE, UNDER_LINE, "---------- Print() ----------", RESET)

	fmt.Print("Second")

	fmt.Printf("Second\n")

	fmt.Print("Second\n")

	fmt.Print("First\nSecond\n")

	var a, b = "First", "Second"
	fmt.Print(a, b, "\n")
	fmt.Print(a, " ", b, "\n")

	var c, d int = 1, 2
	fmt.Print(c, d, "\n")

	fmt.Println(BLUE, UNDER_LINE, "---------- Print() ----------", RESET)

	fmt.Println(a, b)
	fmt.Println(a, " ", b)

	fmt.Println(BLUE, UNDER_LINE, "---------- Printf() ----------", RESET)

	fmt.Printf("a has value : \"%v\" and type : '%T'\n", a, a)
	fmt.Printf("c value : \"%v\" and type : \"%T\"\n", c, c)

	//!			Formatting Vervs for Printf()
	fmt.Println(BLUE, UNDER_LINE, "---------- Formatting Vervs for Printf() ----------", RESET)
	const full_name string = "Arunesh kumar"
	fmt.Printf("full_name : '%v' type : '%T'\n", full_name, full_name)
	fmt.Printf("full_name : %#v\n", full_name)
	fmt.Printf("full_name %% %#v\n", full_name)

	//!			Reference Link
	// https://www.w3schools.com/go/go_formatting_verbs.php
}
