package main

import "fmt"

func main() {
	var name string = "golang"
	fmt.Println("The name is :", name)

	// infer type
	var new_name = "new go"
	fmt.Println("The name is :", new_name)

	var is_aged = true
	fmt.Println("is aged :", is_aged)

	// shorthand syntax
	age := 12
	fmt.Println("age", age)

	// if value not assigned then type needs to be added
	var old_name string
	old_name = "b => c"
	fmt.Println("Old name", old_name)

	var price float32
	price = 10.9
	fmt.Println(price)
}
