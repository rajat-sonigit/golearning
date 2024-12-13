package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	dicenumber := rand.Intn(6) + 1
	fmt.Println("your dice number is ",dicenumber)
	switch dicenumber {
	case 1 : fmt.Println("move 1")
	case 2 : fmt.Println(" move 2")
	case 3 : fmt.Println("move 3")
	case 4 : fmt.Println("move 4")
	case 5 : fmt.Println("move 5")
	case 6 : fmt.Println("move 6")
	}
}
