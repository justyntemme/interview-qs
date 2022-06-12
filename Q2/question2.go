package main

import "fmt"

func addToValue(a *int) {
	//a = 5
	*a = 1
}

func main() {
	a := 4
	addToValue(&a)
	fmt.Println(a)
}

func addToValueNo(a int) {
	//a = 5
}

func addToval() {
	a := 4
	addToValue(&a)
	fmt.Println(a)
}
