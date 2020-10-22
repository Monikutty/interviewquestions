package main

import "fmt"

func main() {

	var array_size int

	fmt.Println("Enter the size of array: ")

	fmt.Scan(&array_size)

	fmt.Println("Enter the array elements: ")

	array := make([]int, array_size)

	sum := 0

	for i := 0; i < array_size; i++ {

		fmt.Scan(&array[i])

		sum = sum + array[i]
	}

	avg := float64(sum) / float64(array_size)

	fmt.Printf("\nAverage is : %f", avg)
}
