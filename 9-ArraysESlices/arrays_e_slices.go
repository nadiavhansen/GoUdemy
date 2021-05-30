package main

import "fmt"

func main() {
	fmt.Println("Arrays e Slices")

	var array1[5] int
	fmt.Println(array1)

	array2 := [5]string{"A", "B", "C", "D", "E"}
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3}
	fmt.Println(array3)

	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice)

	slice = append(slice, 6)
	fmt.Println(slice)

	slice2 := array2[1:3]
	fmt.Println(slice2)

	array2[1] = "S"
	fmt.Println(slice2)

	// Arrays internos
	
	slice3 := make([]float32, 10, 11)
	fmt.Println(slice3)

	slice3 =  append(slice3, 5)
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))

	slice3 =  append(slice3, 6)
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))

	slice4 := make([]float32, 5)
	fmt.Println(slice4)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))
	slice4 = append(slice4, 1)
	fmt.Println(slice4)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))
	

}