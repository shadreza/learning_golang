package main

import "fmt"

func main() {

	// for arrays we need to know the length
	// len() shows the length of the array
	// by defaults falsy values like 0[int], ""[string], false[bool] are considered if values are not set in array

	var nums [3]int
	fmt.Println(nums)
	nums[1] = 1
	fmt.Println(nums)
	println("=== === ===")

	var names [3]string
	fmt.Println(names)
	names[1] = "shad"
	fmt.Println(names)
	println("=== === ===")

	var isSets [3]bool
	fmt.Println(isSets)
	isSets[1] = true
	fmt.Println(isSets)
	println("=== === ===")

	ages := [3]int{1, 2, 3}
	fmt.Println(ages)
	println("=== === ===")

	ages2d := [3][2]int{{1, 2}, {3, 4}, {5, 6}}
	fmt.Println(ages2d)
	fmt.Println(len(ages2d))
	println("=== === ===")
}
