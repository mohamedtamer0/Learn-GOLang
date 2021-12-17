package main

import "fmt"

func main() {
	var i int
	i = 25
	i = 26

	var x = 30
	y := 20
	fmt.Println("Hello World!")
	fmt.Println(i)
	fmt.Println(x)
	fmt.Println(y)

	var j int = 25
	fmt.Printf("%v, %T", j, j)
}
