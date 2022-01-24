package main

import "fmt"

var estado bool

func main() {
	estado = true
	if estado == true {
		fmt.Println(estado)
	} else {
		fmt.Println(estado)
	}
	convAntes(estado)
	ejemploSwitch()
}

//Puedo convertir la variable antes de evaluarla
func convAntes(estado bool) {
	if estado = false; estado == true {
		fmt.Println(estado)
	} else {
		fmt.Println(estado)
	}
}

func ejemploSwitch() {
	switch numero := 3; numero {
	case 1:
		fmt.Println(numero)
	case 2:
		fmt.Println(numero)
	case 3:
		fmt.Println(numero)
	case 4:
		fmt.Println(numero)
	case 5:
		fmt.Println(numero)
	default:
		fmt.Println("No hay otro")
	}
}
