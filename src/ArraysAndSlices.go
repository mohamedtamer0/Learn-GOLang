package main

import "fmt"

/*
Arrays:
Creation
built-in function
working with arrays
Slices:
Creation
built-in function
Working with slices
*/
func main() {
	//Or[...] any size
	grades := [3]int{97, 85, 93}
	fmt.Printf("Grades: %v \n", grades)

	fmt.Println("=================")

	var students [3]string
	fmt.Printf("Students: %v \n", students)
	students[0] = "Tamer"
	students[1] = "Mohamed"
	fmt.Printf("Students: %v \n", students)
	fmt.Printf("Students: %v \n", students[1])
	fmt.Printf("Number of Students: %v \n", len(students))

	fmt.Println("=================")
	var identityMatrix = [3][3]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}
	fmt.Println(identityMatrix)
	fmt.Println("=================")
	var identityMatrix2 [3][3]int
	identityMatrix2[0] = [3]int{1, 0, 0}
	identityMatrix2[1] = [3]int{0, 1, 0}
	identityMatrix2[2] = [3]int{0, 0, 1}
	fmt.Println(identityMatrix2)

	fmt.Println("=================")

	//a := [...]int{1, 2, 3}
	//b := a
	//b[1] = 5
	//fmt.Println(a)
	//	fmt.Println(b)

	fmt.Println("=================")
	//Slices
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b := a[:]
	c := a[3:]
	d := a[:6]
	e := a[3:6]
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)

}
