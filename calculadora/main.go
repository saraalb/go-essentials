package main

import "fmt"

func main() {
	num1:= 1
	num2:= 0
	operacao:= "divisao"

//  Usando if
	if (operacao == "adicao") {
		fmt.Printf("%d + %d = %d",num1, num2, num1+num2)
	 } else if (operacao == "multipiclacao") {
		fmt.Printf("%d * %d = %d",num1, num2, num1*num2)
	 } else if (operacao == "divisao") && (num2!=0) {
		fmt.Printf("%d / %d = %d",num1, num2, num1/num2)
	 } else if (operacao == "subtracao") {
		fmt.Printf("%d - %d = %d",num1, num2, num1-num2)
	 } else if (operacao == "resto" || operacao == "modulo") {
		fmt.Printf("%d %% %d = %d",num1, num2, num1%num2)
	 } else {
		fmt.Printf("insira uma opçao valida")
	 }

// Usando switch
	 switch operacao {
	 case "adicao":
		fmt.Printf("%d + %d = %d",num1, num2, num1+num2)
	 case "subtracao":
		fmt.Printf("%d - %d = %d",num1, num2, num1-num2)
	 case "multiplicacao":
		fmt.Printf("%d * %d = %d",num1, num2, num1*num2)
	 case "divisao":
		if num2!=0 {
		fmt.Printf("%d / %d = %d",num1, num2, num1/num2)
		} else {
			fmt.Println("Divisão por zero. Digite um número válido")
		}
	 case "resto","modulo":
		fmt.Printf("%d %% %d = %d",num1, num2, num1%num2)
	 default:
		fmt.Println("insira uma opção válida.")
	 }
	 
}