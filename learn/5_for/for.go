package main

// The only construct => for no while or any other
func main() {

	// normal for loop
	count := 5
	for i := 0; i < count; i++ {
		println(i)
	}

	println("=== === ===")

	// while loop
	j := 0
	for j < 5 {
		println(j + 1)
		j++
		break
	}

	println("=== === ===")

	// while loop
	k := 0
	for k < 5 {
		println(k + 1)
		k = k + 1
		continue
		k += 1
	}

	// // infininte loop
	// for {
	// 	println("1")
	// }

	println("=== === ===")

	for i := range 10 {
		println(i)
	}
}
