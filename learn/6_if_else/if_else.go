package main

func main() {
	age := 17

	for {
		if age < 18 {
			println("You are not adult yet")
		} else if age == 18 {
			println("You just became adult")
		} else {
			println("You are a grown up")
			break
		}
		age++
	}

	role := "admin"
	hasPermission := true

	if role == "admin" && hasPermission {
		println("YEAH")
	} else {
		println("NAH")
	}

	// variable can be declared inside if construct
	if age := 10; age > 18 {
		println("The age is adult", age)
	} else {
		println("Teenager age", age)
	}

	// go does not have ternary operator
	// have to use normal if else
}
