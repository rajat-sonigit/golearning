package main

import "fmt"

func main() {
	fmt.Println("Hey its the struct")
	rajat := User{"Rajat", "Soni",19,"rjr045",true}
  fmt.Println(rajat)

   count := 29
  var result string

  if  count<25{
	result = "gey"
  } else{
	result = "popo"
  }
  fmt.Println(result)
}

type User struct{
	Name string 
    Cast string 
	Age  int 
    Email string 
	Status bool
}