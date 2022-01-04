package main

import "fmt"

/*
Maps:
What are they?
Creating
manipulation

Structs:
What are they?
Creating
Naming conventions
Embedding
Tags
*/
func main() {
	statePopulations := make(map[string]int)
	statePopulations = map[string]int{
		"java":   1,
		"kotlin": 2,
		"C++":    3,
		"Go":     4,
		"python": 5,
		"A":      6,
	}
	m := map[[3]int]string{}
	fmt.Println(statePopulations, m)
}
