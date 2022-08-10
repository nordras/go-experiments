package main

import "fmt"

func main()  {
	fmt.Print("Hello world \n")

	// variables
	const userName string = "Igor";
	var age int = 30
	occupation := "software engineer"

	fmt.Println("Igor", age)
	fmt.Printf("%v, %t", occupation, occupation)
}