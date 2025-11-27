package evenOdd

import "fmt"

func EvenOrOdd() {
	var number int
	fmt.Println("Enter a number")
	fmt.Scan(&number)

	if number%2 == 0 {
		fmt.Printf("%d is an Even Number", number)
	} else {
		fmt.Printf("%d is an Odd Number\n", number)
	}
}
