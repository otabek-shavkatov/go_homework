package main

import (
	"Go-practise/auth"
	"fmt"
	"math"
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
	uintVar8 := uint8(255)
	uintVar16 := uint16(65535)
	intVar32 := int32(1243222432)
	intVar64 := int64(1274324324432422222)
	intVarFloat := float64(127.1234567890123456789)

	fmt.Println(intVar8)
	fmt.Println(intVar16)
	fmt.Println(uintVar8)
	fmt.Println(uintVar16)
	fmt.Println(intVar32)
	fmt.Println(intVar64)
	fmt.Println(intVarFloat)

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

	// --------------------------------------

	var x int8 = 127
	fmt.Println(x)

	x++
	fmt.Println("x: overflow version", x)

	// --------------------------------------

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

	//  PrintLn() - it will print the value and add a new line at the end of the value
	// Print() - it will print the value without adding a new line at the end of the value
	// Printf() - it will print the value without adding a new line at the end of the value and it will format the value according to the format specifier

	// Part 3
	numberTest := 42
	fmt.Printf("The number is: %d\n", numberTest)
	fmt.Printf("The number is: %b\n", numberTest)
	fmt.Printf("The number is: %o\n", numberTest)
	fmt.Printf("The number is: %x\n", numberTest)
	fmt.Printf("The number is: first the number convert to binary, so and added zero to laft side until all size to 8 %08b\n", numberTest)

	///

	sum, carry := halfAdder(true, true)
	fmt.Println("Sum:", sum, "Carry:", carry)

	sum1, carry1 := halfAdder(true, false)
	fmt.Println("Sum:", sum1, "Carry:", carry1)

	sum2, carry2 := halfAdder(false, true)
	fmt.Println("Sum:", sum2, "Carry:", carry2)

	sum3, carry3 := halfAdder(false, false)
	fmt.Println("Sum:", sum3, "Carry:", carry3)
	///

	A := true
	B := false

	sum, carry = halfAdder(A, B)
	fmt.Printf("A: %t, B: %t, Sum: %t, Carry: %t\n", A, B, sum, carry)

	A1 := true
	B1 := false

	sum, carry = halfAdder(A1, B1)
	fmt.Printf("A: %t, B: %t, Sum: %t, Carry: %t\n", A1, B1, sum, carry)

	A2 := true
	B2 := false

	sum, carry = halfAdder(A2, B2)
	fmt.Printf("A: %t, B: %t, Sum: %t, Carry: %t\n", A2, B2, sum, carry)

	A3 := true
	B3 := false

	sum, carry = halfAdder(A3, B3)
	fmt.Printf("A: %t, B: %t, Sum: %t, Carry: %t\n", A3, B3, sum, carry)

	numberList := []int{0, 1, 10, 42, 127, 128, 200, 255, -42}

	for number := range numberList {
		fmt.Printf("Dec The number is: %d\n", numberList[number])
		fmt.Printf("Bin The number is: %b\n", numberList[number])
		fmt.Printf("Oct The number is: %o\n", numberList[number])
		fmt.Printf("Hex The number is: %x\n", numberList[number])
		fmt.Printf("The number is: first the number convert to binary, so and added zero to laft side until all size to 8 %08b\n", numberList[number])
		fmt.Println()
	}

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
	x1 := (-b + math.Sqrt(D)) / (2 * a)
	x2 := (-b - math.Sqrt(D)) / (2 * a)
	return x1, x2
}

// part 4
func nand(a, b bool) bool {
	return !(a && b)
}

func andFromNand(a, b bool) bool {

	m := nand(a, b)
	and := nand(m, m)
	return and
}

func halfAdder(a, b bool) (bool, bool) {
	m := nand(a, b)
	n := nand(a, m)
	o := nand(b, m)

	carry := andFromNand(n, o)
	sum := nand(n, o)
	return sum, carry
}
