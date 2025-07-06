package main

import "time"

func main() {
	// single condition switch
	i := 1
	switch i {
	case 1:
		println("is 1")
	case 2:
		println("is 2")
	default:
		println("not 1 or 2")
	}

	// multiple condition switch
	switch time.Now().Weekday() {
	case time.Saturday, time.Friday:
		println("It is holiday")
	default:
		println("Need to work")
	}

	// type switch
	// interface {} => any
	// variable.(type) => Considers the type of the variable
	whoAmI := func(i interface{}) {
		switch i.(type) {
		case int:
			println("It is an integer")
		case string:
			println("It's a string")
		case bool:
			println("It's a boolean")
		default:
			println("Unknown type")
		}

	}

	whoAmI(1)
	whoAmI(1.2)
	whoAmI("1")
	whoAmI(true)
}
