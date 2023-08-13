package main

import "fmt" // Formatting string

func main() {
	fmt.Println("Hello world")

	// string
	var nameOne string = "ded"
	var nameTwo string = "My"
	var nameThree string

	fmt.Println(nameOne, nameTwo, nameThree)

	nameFour := "yoshi"
	fmt.Println(nameFour)

	// ints
	var ageOne int = 20
	var ageTwo int = 30
	ageThree := 40

	fmt.Println(ageOne, ageTwo, ageThree)

	// bits, memory
	var numOne int8 = 25
	var numTwo int8 = -128
	var numThree uint8 = 255 // Unsigned int: Impossible to have negative number

	fmt.Println(numOne, numTwo, numThree)

	var scoreOne float32 = 25.98
	var scoreTwo float64 = 8689944984989.7 // apply too much
	scoreThree := 1.5

	fmt.Println(scoreOne, scoreTwo, scoreThree)

	// Print
	fmt.Print("Hello, ")
	fmt.Print("world, \n")

	fmt.Print("Goodbye, \n")
	fmt.Print("world, ")

	// Template string
	name := "Yasuo"
	age := 5

	// Println
	fmt.Println("Name ", age)
	fmt.Println("My name is", name, "and I was", age)

	// Printf (Formatted string) %_ = format specifier
	userId := "d93jdjdiemdewinwncinciss"
	age2 := "20"
	fmt.Printf("{userId: %v } \n", userId) // %v: Default format variable
	fmt.Printf("My age is %q and my name is %q \n", age, name)
	fmt.Printf("My age is %q and my name is %q \n", age2, name)

	fmt.Printf("My age type is %T \n", age) // %T: return the type of variable
	fmt.Printf("Your score is %f point \n", 225.66)
	fmt.Printf("Your score is %0.1f point \n", 225.66)

	// Sprintf (saved formatted string)
	var str = fmt.Sprintf("My age is %v and my name is %v \n", age, name)

	fmt.Println("Save string is: ", str)

	fmt.Println()
	fmt.Println()
	fmt.Println()

	// Array
	var ages [3]int = [3]int{20, 30, 40}
	var names [4]string = [4]string{"Loi", "My", "Miley", "Karma"}

	fmt.Println(ages, len(ages))
	fmt.Println(names, len(names))

	// slices (use arrays under the hood)
	var scores = []int{100, 50, 60}
	scores[2] = 25
	scores = append(scores, 85)

	fmt.Println(scores, len(scores))

	// slices ranges
	rangeOne := names[1:3]
	rangeTwo := names[0:2]

	fmt.Println(rangeOne, rangeTwo)

}
