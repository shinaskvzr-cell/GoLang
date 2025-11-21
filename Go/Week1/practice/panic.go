package main

import "fmt"

func main() {
	fmt.Println("Start of program")

	panic("Something went wrong!")

	fmt.Println("End of program")
}
