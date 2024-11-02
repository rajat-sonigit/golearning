package main

import "fmt"

func main() {
	var username string = "Rajat"
	fmt.Println(username)
	fmt.Printf("The variable type is : %T \n", username)

	// bool type 
	var isLogged bool = true
	fmt.Println(isLogged)
	fmt.Printf("The variable type is : %T \n", isLogged)

	// integer 
	const  myint uint16 = 9099
	fmt.Println(myint)
	fmt.Printf("The variable type is : %T \n", myint)

	// integer 
	const  myFloat float32 = 90.99
	fmt.Println(myFloat)
	fmt.Printf("The variable type is : %T \n", myFloat)

	// no var concept 
	numberofUser := 700000
	fmt.Println(numberofUser)
}
