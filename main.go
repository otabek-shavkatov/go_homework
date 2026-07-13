package main

import (
	"Go-practise/auth"
	"fmt"
)

func main() {
	auth.Register()

	// var a, b, c, d int = 1, 2, 3, 4

	// array := [2]int{1, 2}
	array2 := [...]int{32, 1, 2, 3}

	appendArray := append(array2[:], 4, 5, 6, 7, 8, 9, 10)

	slice := appendArray[4:]

	// [:] - slice operator, if we use it look like that so it is returned look like this response
	// [:, len(appendArray)] - it will return all elements of the array
	// [:,4] it is equal to [0, 4] it will return first 3 becouse lasted element is not included in the slice
	// [4:] it will return all elements starting from 4th index to the end of the array
	// [4:len(appendArray)] it will return all elements starting from 4th index to the end of the array

	fmt.Print(
		// array[1],
		appendArray,
		slice,
	)

	// fmt.Println(a)
	// fmt.Println(b)
	// fmt.Println(c)
	// fmt.Println(d)
	// fmt.Println(as)

	// if
}
