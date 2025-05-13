package operations

import "fmt"

// funções privadas só são visiveis no mesmo pacote
func multiply(a int, b int) int {
	fmt.Printf("Multiplicando %d * %d: ", a, b)
	return a * b
}

//com M maiusculo é uma função "exportada"
func Multiply(a int, b int) int {
	return multiply(a,b)
}