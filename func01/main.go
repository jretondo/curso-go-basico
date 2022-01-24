package main

import "fmt"

func main() {

	//fmt.Println(uno(5))
	//numero, estado := dos(2)

	//	fmt.Println(numero)
	//	fmt.Println(estado)
	//	fmt.Println(tres(3))

	fmt.Println(calculo(5, 46))
	fmt.Println(calculo(1, 2, 3, 4))
	fmt.Println(calculo(15))
	fmt.Println(calculo(3, 5, 9, 4, 5, 6, 3, 6, 9, 4))
	fmt.Println(calculo(3, 5, 6))

}

//Tal cual typescript (Se declara el tipo que entra y el tipo que sale)
func uno(numero int) int {

	return numero * 2
} //Entradas    Salidas
func dos(numero int) (int, bool) {
	if numero == 1 {
		return numero, false
	} else {
		return 10, true
	}
}

//Puedo nombrar las variables del resultado
func tres(numero int) (z int) {

	return numero * 5
}

//los tres puntos indica que no se sabe cuantos parametros le vamos a pasar (pueden ser uno o mas)
func calculo(numero ...int) int {
	total := 0

	//Range devuelve siempre dos valores, un rango. Se usa mucho para vectores, Arrays
	//Rango de todos los elemntos de mi lista
	//El primer valor es el index y el segundo son mi lista
	//El (_) se usa cuando no quiero usar ese parametro, no reserva memoria
	for item, num := range numero {
		total = total + num
		fmt.Printf("\nItem %d \n", item)
	}
	return total
}
