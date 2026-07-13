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

	//

	// lesson Part - 1 Types & Variables

	intVar8 := int8(127)
	intVar16 := int16(22222)
	intVar32 := int32(1243222432)
	intVar64 := int64(1274324324432422222)

	fmt.Println(intVar8)
	fmt.Println(intVar16)
	fmt.Println(intVar32)
	fmt.Println(intVar64)

	// --------------------------------------

	//  overflow

	// negative overflow
	var intVar8Neg int8 = 127
	intVar8Neg += 1
	fmt.Println(intVar8Neg)

	// positive overflow
	var intVar8Pos int8 = 127
	intVar16Pos := int16(intVar8Pos)

	intVar16Pos += 1
	fmt.Println(intVar16Pos)

	// --------------------------------------

	//Type Conversion
	result := fahrenheitToCelsius(celsiusToFahrenheit(100))
	fmt.Println(result)

	// --------------------------------------
	// Part 2 — Functions
	//	 if, else, else if...

	a := int(5)
	b := int(-3)
	c := int(-1)
	if a > 0 {
		fmt.Println("a is greater than 0 max")
	} else {
		fmt.Println("a is not greater than 0")
	}
	if b < 0 {
		fmt.Println("b is less than 0 min")
	} else {
		fmt.Println("b is not less than 0")
	}
	if c < 0 {
		fmt.Println("c negative")
	} else {
		fmt.Println("c is not less than 0")
	}

	/// quadratic ----------------------------

	fmt.Println(quadratic(1, -5, 6))
}

func celsiusToFahrenheit(celsius float64) float64 {
	return (celsius * 9 / 5) + 32
}

func fahrenheitToCelsius(fahrenheit float64) float64 {
	return (fahrenheit - 32) * 5 / 9
}

// ax² + bx + c = 0.
// Test: (1, -5, 6), (1, 0, 1), (1, -2, 1).
// Diskriminant (D) = b² - 4ac

func quadratic(a, b, c float64) (float64, float64) {

	D := (b * b) - (4 * a * c)
	if D < 0 {
		return 0, 0
	} else if D > 0 {
		fmt.Println("D > 0")
	} else {
		fmt.Println("D = 0")
	}
	x1 := (-b + D) / (2 * a)
	x2 := (-b - D) / (2 * a)
	return x1, x2
}
