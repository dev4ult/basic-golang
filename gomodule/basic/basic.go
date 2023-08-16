package basic

import "fmt"

func VariableExamples() {
	// without datatype in defining
	var say = "Hello"
	fmt.Println(say)

	// define explicitly
	var toName string = "John"
	fmt.Println(toName)

	// string concatenation
	fmt.Println(say + " " + toName)

	// multiple declare and define
	var width, height int = 100, 70
	fmt.Println("total area = ", width*height)

	// declared without defined
	var declared int
	fmt.Println("without defined variable output : ", declared)

	// short declaration
	shortVar := "this is a short defined variable"
	fmt.Println(shortVar)
}