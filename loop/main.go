package main
import "fmt"


func main() {
	days := []string{"Sun", "mon", "tu", "wed"}
	for d := 0; d < len(days); d++ {
		fmt.Println(days[d])
	}
}
