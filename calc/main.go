package main

import (
	"fmt"
	//"os"

	"github.com/saraalb/go-essentials/calc/operations"
)

func main() {

 	fmt.Println("Calculadora #GoEssentials")

	result:=operations.Sum(5,10)
	println(result)
	result=operations.Multiply(5,10)
	println(result)
	var err error
	result, err = operations.Divide(5,0)
	if err != nil {
		fmt.Println("Não consegui dividir")
	}
	println(result)
	result=operations.Subtract(5,10)
	println(result) 

}