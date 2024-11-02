package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
 func main(){
 fmt.Println("Hello to our Shop")
 fmt.Println("Enter the rating of the Shop")
  reader := bufio.NewReader(os.Stdin)

  input, _ := reader.ReadString('\n')
  fmt.Println("Your rating is - " ,input)
  number, err  := strconv.ParseInt(strings.TrimSpace(input), 0 , 64 )
  
if err != nil{
	fmt.Println("Error")
}else{
	fmt.Println("Your review with added 1 is ", number+1 )
}
fmt.Println("Thanks for visiting our shop")
 }