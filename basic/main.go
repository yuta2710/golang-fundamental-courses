package main

import (
	"fmt" // Formatting string
	"math"
	"sort"
	"strings"
)

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

	// Standard library
	greeting := "Hello concac"
	fmt.Println(strings.Contains(greeting, "Hello"))
	fmt.Println(strings.Contains(greeting, "hello"))

	// fmt.Println(strings.ReplaceAll(greeting, "Hello", "hi"))

	fmt.Println(strings.ToUpper(greeting))

	// The original variable does not change
	fmt.Println("Original value of greeting variable = ", greeting)
	fmt.Println(strings.Index(greeting, "concac"))
	fmt.Println(strings.Split(greeting, " ")) // Split greeting into the array

	fmt.Println()
	fmt.Println()

	// Sort packages
	userNames := []string{"Loi", "Karma", "My", "Miley"}
	quantities := []int{10, 5, 4, 33, 99, 44}

	sort.Strings(userNames)
	sort.Ints(quantities)

	fmt.Println(userNames)
	fmt.Println(quantities)

	index := sort.SearchInts(quantities, 44)
	nonExistIndex := sort.SearchInts(quantities, 45)

	fmt.Println(index)
	fmt.Println(nonExistIndex)

	fmt.Println()
	fmt.Println()

	// For loop
	x := 0
	for x < 5 {
		fmt.Println("Value of x is: ", x)
		x++
	}
	// Traditional for loop
	for i := 0; i < 5; i++ {
		fmt.Println("Value of i is: ", i)
	}

	objects := []string{"Mikelodic", "DLow", "Hydra", "Double2T"}

	for i := 0; i < len(objects); i++ {
		fmt.Println(objects[i])
	}
	fmt.Println()
	fmt.Println()

	for index, value := range objects {
		fmt.Printf("The value at index %v is %v \n", index, value)
	}
	fmt.Println()
	fmt.Println()

	for _, value := range objects {
		fmt.Printf("The champion in this table is %v \n", value)
		value = "New string" // cannot update or modify the value due to a "value" is a copy version of original value
	}

	fmt.Println()
	fmt.Println()

	// Boolean & Conditional
	fmt.Println(age <= 50)
	fmt.Println(age >= 50)
	fmt.Println(age == 50)
	fmt.Println(age != 40)

	states := []string{"Deo on xiu nao", "Toi cam thay hom nay nhu cc", "Cai dit con me", "Dit con me cuoc doi"}

	for index, value := range states {
		if index == 1 {
			fmt.Printf("Continuing at pos %v and val %v", index, value)
		} else if index == 2 {
			fmt.Print("Cai dit me no thiet chu troi oi")
		}
	}

	fmt.Println()
	fmt.Println()

	// Function
	sayGreeting("Loi")
	sayGoodbye("Loi")

	cycleNames(names[:], sayGreeting)
	cycleNames(names[:], sayGoodbye)

	fmt.Println(circleArea(5.5))

	fmt.Println(getInitials("Loi My"))

	for _, v := range names {
		sayHello(v)
	}

	k := 50

	for _, v := range points { // access the "points" variable indirectly
		var data int = v * k
		fmt.Printf("Result = %v \n", data)
	}

	fmt.Println()
	fmt.Println()

	// Map (like Hashmap in Java)
	menu := map[string]float64{
		"soup":           4.99,
		"pie":            7.99,
		"salad":          6.99,
		"Toffee Pudding": 3.55,
	}

	fmt.Println(menu)
	fmt.Println(menu["pie"])
	fmt.Println(menu["salad"])
	fmt.Println(menu["an noi xa lo"]) // Output: 0

	// Looping maps
	for k, v := range menu {
		fmt.Println(k, "-", v)
	}

	menu["pie"] = 55.99 // Can be modified

	for k, v := range menu {
		fmt.Println(k, "-", v)
	}

	name2 := "tifa"

	updateName(name2) // It just update the copy version\
	fmt.Println("Memory Address: ", &name2)

	m := &name2

	fmt.Println("Memory address of m: ", m)
	fmt.Println("Value at memory address ", *m)

	// Solution for this update the original
	// 1. Assign
	fmt.Println(name2)
	name2 = updateNameSolution1(name2) // assign the function update to the original variable
	fmt.Println(name2)

	// 2. Pointer
	name3 := "Loi"

	fmt.Println(name3)
	updateNameSolutionPointer(&name3)
	fmt.Println(name3)

	// Update the pointer wrapper value
	updateMenu(menu)
	fmt.Println(menu)

	bill := newBill("Nguyen Phuc Loi")
	bill.updateTip(10)

	fmt.Println(bill.format())

	shapes := []shape{
		square{length: 15.2},
		circle{radius: 700},
	}

	for _, v := range shapes {
		printShapeInfo(v)
		fmt.Println("---")
	}
}

type shape interface {
	area() float64
	circumf() float64
}
type square struct {
	length float64
}

type circle struct {
	radius float64
}

func (s square) area() float64 {
	return s.length * s.length
}

func (s square) circumf() float64 {
	return s.length * 4
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) circumf() float64 {
	return 2 * math.Pi * c.radius
}

func printShapeInfo(s shape) {
	fmt.Printf("Area of %T is: %0.2f \n", s, s.area())
	fmt.Printf("Circumreference of %T is : %0.2f \n", s, s.circumf())
}

func updateName(x string) {
	x = "wedge"
}

func updateNameSolution1(x string) string {
	x = "Wet ass pussy"

	return x
}

func updateNameSolutionPointer(x *string) {
	*x = "Data"
}

func updateMenu(x map[string]float64) {
	x["coffee"] = 66999.99
}

func sayGreeting(data string) {
	fmt.Printf("Hello bro %v \n", data)
}

func sayGoodbye(data string) {
	fmt.Printf("Goodbye bro %v \n", data)
}

func cycleNames(names []string, f func(string)) {
	for _, v := range names {
		f(v)
	}
}

func circleArea(r float64) float64 {
	return math.Pi * r * r
}

// Multiple return values

func getInitials(data string) (string, string) {
	s := strings.ToUpper(data)

	splitter := strings.Split(s, " ")

	var initials []string

	for _, v := range splitter {
		initials = append(initials, v[:1])
	}

	if len(initials) > 1 {
		return initials[0], initials[1]
	}

	return initials[0], "_"
}
