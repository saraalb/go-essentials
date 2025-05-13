package operations

import (
	"errors"
	"fmt"
)

/* func Divide(a int, b int) int {
	fmt.Printf("Dividindo %d / %d: ",a,b)
	if b == 0 {
		//fmt.Println("Não posso fazer divisão por zero!")
		//return 0
		panic("Divisão por zero não permitida") //para a execução completamente
	}
	return a / b
} */

func Divide(a int, b int) (int, error) {
	fmt.Printf("Dividindo %d / %d: ",a,b)
	if b == 0 {
		err:= errors.New("Divisão por zero não permitida")
		return 0, err
	}
	return a / b, nil
}