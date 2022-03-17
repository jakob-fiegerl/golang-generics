package main

import "fmt"

type Number interface {
	int | float64
}

func Max[T Number](array []T) T {
	var result T = 0
	for key, _ := range array {
		if array[key] > result {
			result = array[key]
		}
	}
	return result
}

func main() {
	intValues := []int{2, 3, 15, 27, 11, 13}
	fmt.Println("Int values max: ", Max(intValues))
	floatValues := []float64{2.1, 3.2, 15.4, 27.9, 11.1, 13.2}
	fmt.Println("Float values max: ", Max(floatValues))
}
