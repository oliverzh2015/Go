package fib

import "fmt"

func main() {
	myArray := [5]string{"I", "am", "stupid", "and", "weak"}
	for x, y := range myArray {
		fmt.Println(x, y)
	}
}
