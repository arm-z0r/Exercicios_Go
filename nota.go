package main

import "fmt"

func LeFlutuante() float64 {
	var i float64

	_, err := fmt.Scanf("%v", &i)
	for err != nil {
		_, err = fmt.Scanf("%v", &i)
	}
	return i
}

func main() {
	println("Digite a nota da AD1: ")
	ad1 := LeFlutuante()
	println("Digite a nota da AP1: ")
	ap1 := LeFlutuante()
	println("Digite a nota da AD2:")
	ad2 := LeFlutuante()
	objetivo := 6*2 - (ad1*0.2 + ap1*0.8 - (ad2*0.2)/0.8)
	fmt.Printf("Voce precisa tirar %.2f ", objetivo)
}

//Regra para o cálculo da média: [#01] - N1 e N2 = (ADx4 + APx6)/10; N = ( N1 + N2 )/2; Se N >=6 então NF = N senão NF = [ N + AP3 ]/2
