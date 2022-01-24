package main

import "fmt"

//Funciones anonimas y clousures

//A las funciones anonimas puedo modificarlas en tiempo de ejecución

//Puedo declarar variables de tipo función
var calculo func(int, int) int = func(num1 int, num2 int) int {
	return num1 + num2
}

func main() {
	fmt.Printf("Sumo 5 + 7 = %d \n", calculo(5, 7))

	//Redefinimos la función
	//Por más que se redifina, siempre hay que respetar las variables de entrada y salida
	calculo = func(num1 int, num2 int) int {
		return num1 - num2
	}
	fmt.Printf("Resto 6 - 4 = %d \n", calculo(6, 4))

	Operaciones()

	//Clousures
	tablaDos := 2
	MiTabla := Tabla(tablaDos)
	for i := 1; i < 11; i++ {
		fmt.Println(MiTabla())
	}
}

func Operaciones() {
	resultado := func() int {
		var a int = 23
		var b int = 10
		return a + b
	}

	fmt.Printf("Resultado %d\n", resultado())
}

//Clousures => Isolaciñon de código
func Tabla(valor int) func() int {
	numero := valor
	secuencia := 0

	return func() int {
		secuencia += 1
		return numero * secuencia
	}
}
