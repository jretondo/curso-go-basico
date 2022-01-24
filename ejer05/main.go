package main

import "fmt"

//ciclos
func main() {

	otroCiclo04()
}
func cicloBasico() {
	i := 1
	for i < 10 {
		fmt.Println(i)
		i++
	}
}

func otroCiclo01() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
}

func otroCiclo02() {
	numero := 0
	for {
		fmt.Println("Continuo...")
		fmt.Println("Ingrese el número secreto")
		fmt.Scanln(&numero)
		if numero == 100 {
			break
		}
	}
}

func otroCiclo03() {
	var i = 0
	for i < 10 {
		//Sirve para enmascarar la salida
		fmt.Printf("\n valor de i: %d", i)
		if i == 5 {
			fmt.Print("multiplicamos por 2 \n")
			i = i * 2
			//Sirve para torcer la ejecución, para que vuelva para arriba, sino iriía para abajo
			//Hace que se verifique la condición del for nuevamente
			continue
		}
		fmt.Printf("               pasó por acá \n")
		i++
	}
}

func otroCiclo04() {
	var i int = 0
	//Se hace un marcado - etiqueta, no afecta en el código, pero ayuda a identificar
RUTINA:
	fmt.Println("Volvió para acá")
	for i < 10 {
		if i == 4 {
			i = i + 2
			fmt.Println("voy a RUTINA")
			//Vuelve para la etiqueta que le indiqué
			goto RUTINA
		}
		fmt.Printf("Valor de i: %d\n", i)
		i++
	}
}
