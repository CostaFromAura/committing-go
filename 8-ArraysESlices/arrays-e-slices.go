package main

import "fmt"

func main() {
	fmt.Println("Arrays e Slices")

	var array1 [5]int
	array1[0] = 1
	fmt.Println(array1)

	array2 := [5]string{"Posição 1"}
	fmt.Println(array2)

	slice := []int{10, 11, 12, 13, 14, 15, 16, 17}
	fmt.Println(slice)

	slice = append(slice, 18)
	fmt.Println(slice)

	// Arrays Internos

	slice3 := make([]float32, 10, 11)
	fmt.Println(slice3)

	slice3 = append(slice3, 5)
	slice3 = append(slice3, 6)

	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))
}
