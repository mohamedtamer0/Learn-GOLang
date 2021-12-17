package main

import (
	"fmt"
)

/*
Naming convention
typed constants
untyped constants
Enumerated constants
Enumeration constants
*/

const aa int16 = 27

const (
	one   = iota
	two   = iota
	three = iota
)

func main() {

	const myConst int = 25
	fmt.Printf("%v, %T \n", myConst, myConst)

	//const myConst2 float64 = math.Sin(1.57)
	//fmt.Printf("%v, %T \n", myConst2, myCons2)

	const a = 14
	const b string = "foo"
	const c float32 = 3.14
	const d bool = true
	fmt.Printf("%v \n", a)
	fmt.Printf("%v \n", b)
	fmt.Printf("%v \n", c)
	fmt.Printf("%v \n", d)

	fmt.Println("===============")
	const aa int = 14
	fmt.Printf("%v, %T \n", aa, aa)
	var bb int = 14
	fmt.Printf("%v, %T \n", aa+bb, aa+bb)

	fmt.Println("===============")
	fmt.Printf("%v \n", one)
	fmt.Printf("%v \n", two)
	fmt.Printf("%v \n", three)

}
