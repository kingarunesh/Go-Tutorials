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

	fmt.Println(BLUE, UNDER_LINE, "---------- Arrays ----------", RESET)

	//NOTE : 		length is defiend
	var arr_1 = [5]int{1, 2, 3}
	var arr_2 = [5]int{1, 2, 3, 4, 5}
	// var arr_2 = [5]int{1, 2, 3, 4, 5, 6, 7}   error

	arr_3 := [5]int{1, 2, 3, 4}
	arr_4 := [3]int{1, 2, 3}
	// arr_4 := [3]int{1, 2, 3, 4}   error

	fmt.Println(arr_1)
	fmt.Println(arr_2)
	fmt.Println(arr_3)
	fmt.Println(arr_4)

	//NOTE : 	  length is inferred
	var arr_5 = [...]int{1, 2, 3}
	var arr_6 = [...]int{1, 2, 3, 4, 5}
	fmt.Println(arr_5)
	fmt.Println(arr_6)

	//NOTE : 	 Examples
	var users = [4]string{"Aohan", "Bohan", "Cohan", "Dohan"}
	// var users = [4]string{"Aohan", "Bohan", "Cohan"}
	// var users = [...]string{"Aohan", "Bohan", "Cohan"}
	fmt.Println(users)

	//NOTE : 		Access Arrays Element
	fmt.Println(BLUE, UNDER_LINE, "---------- Access Arrays Ele ----------", RESET)
	var a = [4]int{11, 22, 33, 44}

	b := [4]int{55, 66, 77, 88}
	// const b = [4]int{11, 22, 33, 44}  - we can not define array in const

	fmt.Println(a)
	fmt.Println(a[0])
	fmt.Println(a[1])
	fmt.Println(a[2])
	fmt.Println(a[3])

	fmt.Println(b)
	fmt.Println(b[0])
	fmt.Println(len(b))
	fmt.Println(b[len(b)-1])

	//NOTE : 		change Array Elements
	fmt.Println(BLUE, UNDER_LINE, "---------- Change Array Ele ----------", RESET)

	var ages = [4]int{2005, 1999, 2002, 2007}

	fmt.Println(ages)

	ages[1] = 2000
	ages[1] = 2025
	fmt.Println(ages)

	//NOTE : 		 Array Initialization
	fmt.Println(BLUE, UNDER_LINE, "---------- Array Initialization ----------", RESET)

	var a = [4]int{}           //	 not initialized
	var b = [4]int{11, 22}     // partially initialized
	var c = [4]int{1, 2, 3, 4} // fully initialized

	var d = [5]int{1: 22, 3: 44}

	var e = [...]int{1: 11, 3: 33}
	var f = [...]int{1: 11, 3: 33, 5: 55, 7: 77, 9: 99}

	g := [...]int{1: 11, 3: 33, 5: 55, 7: 77, 9: 99}

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)

	fmt.Println(len(e))
	fmt.Println(len(f))

	fmt.Println(g)
	fmt.Println(len(g))

}
