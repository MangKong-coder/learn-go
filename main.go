package main

import (
	"fmt"
	"strings"
)

func main() {
	for {
	var desiredExample string;

	fmt.Println("Which example would you like to run?")
	fmt.Println("1. Arrays")
	fmt.Println("2. Slices")
	fmt.Println("3. Maps")
	fmt.Println("4. Structs")
	fmt.Println("5. If Statement")
	fmt.Println("6. Switch Statement")
	fmt.Println("7. Loops")
	fmt.Println("8. Defer Panic Recover")
	fmt.Println("9. Pointers")
	fmt.Scanln(&desiredExample)

	switch strings.ToLower(desiredExample) {
	case "1", "arrays":
		array_example()
	case "2", "slices":
		slice_example()
	case "3", "maps":
		map_example()
	case "4", "structs":
		struct_example()
	case "5", "if statement":
		if_statement_example()
	case "6", "switch statement":
		switch_example()
	case "7", "loops":
		loops_example()
	case "8", "defer panic recover":
		defer_panic_recover_example()
	case "9", "pointers":
		pointers_example()
	default:
		fmt.Println("Invalid input. Please try again.")
		continue
	}

	fmt.Println("Would you like to run another example? (y/n)")
	var runAnotherExample string;
	fmt.Scanln(&runAnotherExample)
	if strings.ToLower(runAnotherExample) == "n" {
		break
	} else {
		continue
	}

}
	
}

func array_example() {
	//Arrays
	arr := [5]int{}
	arr[0] = 1
	arr[1] = 5
	arrLen := len(arr)
	fmt.Println(arr)
	fmt.Println(arrLen)
}

func slice_example() {
	//Slices
	sli := []int{}
	sli = append(sli, 1, 2, 3)
	slice := make([]int, 0)
	slice = append(slice, 1, 2, 3)
	fmt.Println(slice)
	fmt.Println(sli)
}

func map_example() {
	//Maps 	
	// maps := make(map[string]int)
	maps := map[string]int{
		"California": 1,
		"Texas":      2,
		"Florida":    3,
		"New York":   4,
		"New Jersey": 5,
	}
	fmt.Println(maps["California"])
}

func struct_example() {
	//Structs
	type Person struct {
		Name string
		Age  int
	}

	leighne := Person{
		Name: "Leighne",
		Age:  17,
	}

	fmt.Println(leighne)
	fmt.Println("this is the age of person leighne: ", leighne.Age)

}

func if_statement_example() {
	x := 5;

	if x > 10 {
		fmt.Println("x is greater than 10")
	} else {
		fmt.Println("x is less than 10")
	}
}

func switch_example() {
	var x int;

	fmt.Scanln(&x)

	switch x {
	case 1:
		fmt.Println("x is 1")
	case 2:
		fmt.Println("x is 2")
	case 3:
		fmt.Println("x is 3")
		fallthrough
	default:
		fmt.Println("x is not 1, 2, or 3")
	}
}

func loops_example() {
	// for loop
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	// j is scoped to the whole function
	// j := 10
	// for ; j < 20; j++ {
	// 	fmt.Println(j)
	// }

	// while loop in golang
	// k := 0
	// c := true
	// for c {
	// 	if k > 10 {
	// 		fmt.Println("k is greater than 10")
	// 		break
	// 	} else {
	// 		fmt.Println("k is less than 10")
	// 		k++
	// 	}
	// }

	// iterating through a collection or a for range loop

	// slice or can be an array
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i, v := range s {
		fmt.Println(i, v)
	}

	// map 
	m := map[string]int{
		"California": 1,
		"Texas":      2,
		"Florida":    3,
		"New York":   4,
		"New Jersey": 5,
	}

	for k, v := range m {
		fmt.Println(k, v)
	}
}

func defer_panic_recover_example() {
	// // defer: deleay execution of a function until the surrounding function returns
	// fmt.Println("start")
	// defer fmt.Println("middle");
	// fmt.Println("end");


	// // panic: stops the execution of the surrounding function and returns the error
	// fmt.Println("start")
	// panic("We made a boo boo");
	// fmt.Println("end");
	// // recover: recovers from a panic
}

func pointers_example() {
	a := 42
	// equivalent to: pointer to a
	b := &a

	// *b dereferences the pointer or gets the value of the pointer
	fmt.Println(a, *b)
	a = 43
	fmt.Println(a, *b)
}



