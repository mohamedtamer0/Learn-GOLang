package main

import "fmt"

/*
Boolean type
Numeric Type:
Integers
Float
Complex Numbers
Text types
*/
func main() {
	var n bool = false
	fmt.Printf("%v,%T \n", n, n)

	x := 1 == 1
	y := 1 == 2
	fmt.Printf("%v,%T \n", x, x)
	fmt.Printf("%v,%T \n", y, y)

	var defaultValue bool
	fmt.Printf("%v,%T \n", defaultValue, defaultValue)

	fmt.Println("====================")

	var a int = 10
	var b int8 = 3
	fmt.Println(a + int(b))

	fmt.Println("====================")
	//Integers
	s := 8              //2^3
	fmt.Println(s << 3) //2^3 * 2^3 = 2^6
	fmt.Println(s >> 3) //2^3 / 2^3 = 2^0

	fmt.Println("====================")
	//Float
	var f float64 = 3.14
	f = 13.7e72
	f = 2.1e14
	fmt.Printf("%v, %T \n", f, f)
	fmt.Println("====================")
	f1 := 10.2
	f2 := 3.7
	fmt.Println(f1 + f2)
	fmt.Println(f1 - f2)
	fmt.Println(f1 * f2)
	fmt.Println(f1 / f2)

	fmt.Println("====================")
	//Complex

}
