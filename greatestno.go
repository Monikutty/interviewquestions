package main

import "fmt"

func main() {

	var array_size int

	fmt.Println("Enter the size of array: ")

	fmt.Scan(&array_size)

	fmt.Println("Enter the array elements: ")

	array := make([]int, array_size)

	for i := 0; i < array_size; i++ {

		fmt.Scan(&array[i])

	}

	max := array[0]

	for i := 1; i < array_size; i++ {

		if max < array[i] {

			max = array[i]
		}
	}

	fmt.Printf("\nMax element is : %d", max)
}
