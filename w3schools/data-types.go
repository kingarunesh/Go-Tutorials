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

	fmt.Println(BLUE, UNDER_LINE, "---------- ( Basic ) Data Types ----------", RESET)

	var a bool = true
	var b int = 1
	var c float64 = 1.1
	var d string = "hello"

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	fmt.Println(BLUE, UNDER_LINE, "-------- ( Boolean ) Data Types --------", RESET)

	var e bool // default is false
	var f bool = true
	var g = false
	h := true

	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(h)

	fmt.Println(BLUE, UNDER_LINE, "-------- ( Integer ) Data Types --------", RESET)
	//   Signed integers   - can store both positive and negative values
	//   Unsigned integers - can only positive values

	//!			Signed Integers
	var a int = 1 // we can store max/min values based on system 32bit or 64bit
	var aa int = -1
	fmt.Printf("Value : %v , Type : %T\n", a, a)
	fmt.Printf("Value : %v , Type : %T\n", aa, aa)

	var b int8 = 127   // not more than 127
	var bb int8 = -128 // not less than -128
	fmt.Printf("Value : %v , Type : %T\n", b, b)
	fmt.Printf("Value : %v , Type : %T\n", bb, bb)

	var c int16 = 32767
	var cc int16 = -32768
	fmt.Printf("Value : %v , Type : %T\n", c, c)
	fmt.Printf("Value : %v , Type : %T\n", cc, cc)

	var d int32 = 2147483647
	var dd int32 = -2147483648
	fmt.Printf("Value : %v , Type : %T\n", d, d)
	fmt.Printf("Value : %v , Type : %T\n", dd, dd)

	var e int64 = 9223372036854775807
	var ee int64 = -9223372036854775808
	fmt.Printf("Value : %v , Type : %T\n", e, e)
	fmt.Printf("Value : %v , Type : %T\n", ee, ee)

	//!			UnSigned Integers
	var aaa uint = 1
	fmt.Printf("Value : %v , Type : %T\n", aaa, aaa)

	var bbb uint8 = 127 // not more than 127
	fmt.Printf("Value : %v , Type : %T\n", bbb, bbb)

	var ccc uint16 = 32767
	fmt.Printf("Value : %v , Type : %T\n", ccc, ccc)

	var ddd uint32 = 2147483647
	fmt.Printf("Value : %v , Type : %T\n", ddd, ddd)

	var eee uint64 = 9223372036854775807
	fmt.Printf("Value : %v , Type : %T\n", eee, eee)

	fmt.Println(BLUE, UNDER_LINE, "-------- ( Float ) Data Types --------", RESET)

	var a = 1.1   // default float type is float64
	var aa = -1.1 // default float type is float64
	var b float32 = 1234.5678912345
	var bb float32 = -1234.5678912345
	var c float64 = 1234.5678912345
	var cc float64 = -1234.5678912345

	fmt.Printf("Value : %v , Type : %T\n", a, a)
	fmt.Printf("Value : %v , Type : %T\n", aa, aa)
	fmt.Printf("Value : %v , Type : %T\n", b, b)
	fmt.Printf("Value : %v , Type : %T\n", bb, bb)
	fmt.Printf("Value : %v , Type : %T\n", c, c)
	fmt.Printf("Value : %v , Type : %T\n", cc, cc)

	fmt.Println(BLUE, UNDER_LINE, "-------- ( String ) Data Types --------", RESET)
	var a = "First"
	var b string = "Second"
	c := "third"
	var d string

	fmt.Printf("Value : %v , Type : %T\n", a, a)
	fmt.Printf("Value : %v , Type : %T\n", b, b)
	fmt.Printf("Value : %v , Type : %T\n", c, c)
	fmt.Printf("Value : \"%v\" , Type : %T\n", d, d)
}
