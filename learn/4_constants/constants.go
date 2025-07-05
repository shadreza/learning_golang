package main

import "fmt"

const age = 30

// const age_new := 30

func main() {
	const name = "golang"

	// name = "new name"

	fmt.Println(name)
	fmt.Println(age)

	const (
		port = 5000
		host = "localhost"
	)

	fmt.Println(port, host)
}
