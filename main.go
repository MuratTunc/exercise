package main

import (
	"fmt"
)

func main() {

	nums := []int{10, 20, 30}

	for index, value := range nums {
		fmt.Println(index, value)
	}

	original := []int{1, 2, 3}
	copySlice := make([]int, len(original))
	copy(copySlice, original)

	fmt.Println("Original:", original)
	fmt.Println("Copy:", copySlice)

	s := []int{1, 2, 33, 4, 5}

	fmt.Println("Slice:", s)
	size := len(s)
	fmt.Println("size:", size)

	//max element find
	var max = 0
	for i := 0; i < size-1; i++ {
		for m := i + 1; m < size; m++ {
			if s[i] > s[m] {
				max = s[i]
			} else if s[i] <= s[m] {
				max = s[m]
				i++
				break
			}
		}

	}
	fmt.Println("max_number:", max)

	//reverse slice
	slice2 := make([]int, len(s))

	for index, value := range s {
		slice2[size-index-1] = value

	}

	fmt.Println("Reversed:", slice2)

}
