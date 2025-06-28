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

var x string
var y string = "Outside - String"

const z string = "Constant Variable"

//!		we can not use outside function 👇
// xx := "Outside := string"

func main() {

	fmt.Println(BLUE+UNDER_LINE, "--------------- 1-Rule ---------------", RESET)

	//		 var variable_name type = value
	// var a string = 'hello'   	invalid
	var a string = "hello"
	var b int = 1
	var c bool = true
	var d float32 = 1.1
	var e float64 = 1.1

	var f = "hello"
	f = "hello arunesh"

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)

	//		empty variable defiend
	var g int
	fmt.Println("g : ", g)

	g = 9999
	fmt.Println("g : ", g)

	g = 7777
	fmt.Println("g : ", g)

	fmt.Println(BLUE+UNDER_LINE, "--------------- 2-Rule ---------------", RESET)
	//		2-Rule : 	variable_name := value
	aa := "hello"
	bb := 1
	cc := 1.1
	dd := true
	// dd := false
	// dd := "string"

	fmt.Println(aa)
	fmt.Println(bb)
	fmt.Println(cc)
	fmt.Println(dd)

	ee := "string"
	fmt.Println(ee)

	// ee = 11		invalid
	ee = "new string"
	fmt.Println(ee)

	fmt.Println(BLUE+UNDER_LINE, "--------------- Defualt Values ---------------", RESET)

	var aaa string
	var bbb int
	var ccc float64
	var ddd bool
	fmt.Println("aaa : ", aaa) //  ""
	fmt.Println("bbb : ", bbb) // 0
	fmt.Println("ccc : ", ccc) // 0
	fmt.Println("ddd : ", ddd) // false

	aaa = "Hello String"
	fmt.Println("aaa : ", aaa)

	fmt.Println(BLUE+UNDER_LINE, "------------- Access Outside variables -------------", RESET)
	fmt.Println("x : ", x)
	fmt.Println("y : ", y)

	x = "inside function - x : value change"
	y = "inside function - y : value change"

	fmt.Println("x : ", x)
	fmt.Println("y : ", y)

	//ERROR :			testing
	first()

	fmt.Println(BLUE+UNDER_LINE, "------------- Multiple Variable Declaration  -------------", RESET)

	var aaaa, bbbb, cccc int = 1, 2, 3
	fmt.Println("aaaa : ", aaaa)
	fmt.Println("bbbb : ", bbbb)
	fmt.Println("cccc : ", cccc)

	var aaaaa, bbbbb, ccccc = 1, "hello", true
	fmt.Println("aaaaa : ", aaaaa)
	fmt.Println("bbbbb : ", bbbbb)
	fmt.Println("ccccc : ", ccccc)

	var (
		aaaaaa int
		bbbbbb float64 = 1.1
		cccccc string  = "hello"
		dddddd         = true
	)

	fmt.Println("aaaaaa : ", aaaaaa)
	fmt.Println("bbbbbb : ", bbbbbb)
	fmt.Println("cccccc : ", cccccc)
	fmt.Println("dddddd : ", dddddd)

	fmt.Println(BLUE+UNDER_LINE, "------------- Variable Naming Rules  -------------", RESET)

	var age int = 21
	var _age int = 22
	var age1 int = 23
	var age_now int = 24
	fmt.Println(age)
	fmt.Println(_age)
	fmt.Println(age1)
	fmt.Println(age_now)

	//!		Camel Case
	var FirstName string = "Arunesh"
	fmt.Println(FirstName)

	//!		Pascal Case
	var LastName string = "kumar"
	fmt.Println(LastName)

	//!		Snake Case
	var full_name string = "Arunesh kumar"
	fmt.Println(full_name)

	fmt.Println(BLUE+UNDER_LINE, "------------- Constants Variable -------------", RESET)

	//		Untyped Constant
	const first_name = "Arunesh"
	fmt.Println(first_name)

	//		Typed Constant
	const last_name string = "kumar"
	fmt.Println(last_name)

	fmt.Println(z) // 	outside variable

	//!			Multiple Constants Declaration
	const (
		A int = 1
		B     = "hello"
	)

	fmt.Println(A)
	fmt.Println(B)

}

func first() {
	// x = "inside function - x : value change"
	// y = "inside function - y : value change"

	fmt.Println("x : ", x)
	fmt.Println("y : ", y)
}
