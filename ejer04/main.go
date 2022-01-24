package main

import (
	"bufio"
	"fmt"
	"os"
)

var numero1 int
var numero2 int
var resultado int
var leyenda string

func main() {
	fmt.Println("Ingrese número 1: ")
	fmt.Scanf("%d", &numero1)

	fmt.Println("Ingrese número 2: ")
	fmt.Scanf("%d", &numero2)

	fmt.Println(numero1 + numero2)

	fmt.Printf("Escriba una leyenda:")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		leyenda = scanner.Text()
	}

	resultado = numero1 + numero2

	fmt.Println(leyenda, resultado)
}

func otraForma() {
	fmt.Println("Ingrese número 1: ")
	fmt.Scanln(&numero1)

	fmt.Println("Ingrese número 2: ")
	fmt.Scanln(&numero2)

	fmt.Printf("%d", numero1+numero2)
}
