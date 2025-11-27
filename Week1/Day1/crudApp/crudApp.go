package crudApp
import "fmt"

func CrudApp(){
	for {
		fmt.Println("---CRUD APP MENU---")
		fmt.Println("1.Create")
		fmt.Println("2.Read")
		fmt.Println("3.Update")
		fmt.Println("4.Delete")
		fmt.Println("5.Exit")

		var choice int
		fmt.Println("Enter your choice")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Create selected")
		case 2:
			fmt.Println("Read selected")
		case 3:
			fmt.Println("Update selected")
		case 4:
			fmt.Println("Delete selected")
		case 5:
			fmt.Println("Exit selected")
			return
		default:
			fmt.Println("Invalid choice!! Try again")
		}
	}
}