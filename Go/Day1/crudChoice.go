package main
import "fmt"

func main(){
	fmt.Println("CRUD operations")
	fmt.Println("1. Create Item")
	fmt.Println("2. Read Item")
	fmt.Println("3. Update Item")
	fmt.Println("4. Delete Item")
	fmt.Println("5. Exit")

	var choice int
	fmt.Println("Enter your choice")
	fmt.Scan(&choice)

	switch choice{
	case 1:
		fmt.Println("Create operation selected")
	case 2:
		fmt.Println("Read operation selected")
	case 3:		
		fmt.Println("Update operation selected")
	case 4:
		fmt.Println("Delete operation selected")
	case 5:
		fmt.Println("Exiting....")
		return
	default:
		fmt.Println("Invalid choice")
	}

}