package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	fmt.Print("Welcome to the store")
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("The the Item you want")
	input, _ := reader.ReadString('\n')
	fmt.Println("Your Ordered Material is ",input)
}
