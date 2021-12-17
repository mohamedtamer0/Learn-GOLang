package main

/*
Variable declaration :
- var foo int
- var foo int = 25
- foo := 25
Can't redeclare variables , but can shadow them
All variables must be used
Visibility :
lower case first letter for package scope
upper case first letter to export
no private scope
*/

import (
	"fmt"
	"strconv"
)

var g int = 99

var (
	firstName string = "Mohamed"
	lastName  string = "Tamer"
)

func main() {
	var i int
	i = 25
	i = 26
	fmt.Println("Hello World!")
	fmt.Println(i)

	var x = 30
	y := 20
	fmt.Println(x)
	fmt.Println(y)

	fmt.Println("===")

	var j int = 25
	fmt.Printf("%v, %T \n", j, j)

	fmt.Printf("%v, %T \n", g, g)

	fmt.Println("===")

	var c string
	c = strconv.Itoa(j)
	fmt.Printf("%v, %T \n", c, c)
}
