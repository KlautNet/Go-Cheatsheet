package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {

	//Hello world
	fmt.Println("Hello World")

	//Data Types
	var x int = 5
	var y int
	y = 7
	var sum int = x + y

	//Array
	var a [5]int
	b := [5]int{5, 4, 3, 2, 1}

	//Index 0..
	a[3] = 5

	//Slices = Abstract version of Array -> No fixed Lenght
	sl := []int{6, 5, 4}
	sl = append(sl, 3)

	//Maps
	exm := make(map[string]int)
	exm["three"] = 3
	exm["four"] = 4
	delete(exm, "four")

	//Emit Type
	z := 9

	//Statements
	if x > 6 {
		fmt.Println("Greater than 6")
	} else {
		fmt.Println("Smaller than 6")
	}

	//For Loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	//Array Iteration
	arr := []string{"apple", "cucumber", "melon"}

	for index, value := range arr {
		fmt.Println("index:", index, "value:", value)
	}

	//Map Iteration
	exmap := make(map[string]string)
	exmap["a"] = "alpha"
	exmap["b"] = "beta"

	for key, value := range exmap {
		fmt.Println("key:", key, "value:", value)
	}

	//While Loop
	c := 0
	for c < 5 {
		fmt.Println(c)
		c++
	}

	//Exec functions
	print("Test")

	result := add(2, 3)
	fmt.Println(result)

	result1, err := sqrt(16)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result1)
	}

	//Definde Structs
	p := person{name: "Johannes", age: 16}

	//Pointers -> func inc(x int) {x++} wont do anything to i := 7 because structure path is incremented

	t := 7
	inc(&t)
	fmt.Println(t)

	//Print
	fmt.Println(z)
	fmt.Println(sum)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(sl)
	fmt.Println(exm["three"])
	fmt.Println(p.age)
}

//Functions
func print(message string) {
	fmt.Println("[System]", message)
}

//Return
func add(a int, b int) int {
	return a + b
}

//Multiple Returns
func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Undefined for negative inputs")
	}

	return math.Sqrt(x), nil
}

//Struct Types
type person struct {
	name string
	age  int
}

//Pointers
func inc(x *int) {
	*x++
}
