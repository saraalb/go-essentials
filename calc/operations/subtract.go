package operations

import "fmt"

func Subtract(a int, b int) int {
	fmt.Printf("Subtraindo %d - %d: ",a,b)
	return a - b
}