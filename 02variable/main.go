package main

import "fmt"

const LoginToken string = "sfmksmkf"

func main() {
	var username string = "Rudra"
	fmt.Println(username)
	fmt.Printf("variable is of type: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("variable is of type: %T \n", isLoggedIn)

	var samallValue uint8 = 255
	fmt.Println(samallValue)
	fmt.Printf("variable is of type: %T \n", samallValue)

	var samallFloat float32 = 255.985735737
	fmt.Println(samallFloat)
	fmt.Printf("variable is of type: %T \n", samallFloat)

	//default value and alias
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("another variable : %T \n", anotherVariable)

	var anotherVariables string
	fmt.Println(anotherVariables)
	fmt.Printf("another variable : %T \n", anotherVariables)

	//implicit type
	var web = "rp"
	fmt.Println(web)
	fmt.Printf("another variable : %T \n", web)

	//no var
	numberOfUser := 3000
	fmt.Println(numberOfUser)

}
