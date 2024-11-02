package main
import (
	"fmt"
	"time"
)
func main() {
	fmt.Println("Hey we are going to study the time")
	presenttime := time.Now()
	fmt.Println(presenttime.Format("01-02-2006 Monday"))
	 var fruit[6] string
	 fruit[0] = "Apple"
	  fruit[1] = "banana"
	  fruit[2] = "mango"
	  fruit[3] = "orange"
	  fruit[4] = "strawberry"
	  fruit[5] = "gauava"
  fmt.Println(fruit)
  var fruitList = append(fruit[1:4])
 fmt.Println(fruitList)
}
