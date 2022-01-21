package main

import (
	"fmt"
	"strconv"
)

//Las variables que empiezan con minusculas, son variables privadas
//Las variables respetan los scopes
var numero int
var texto string = "Variable inicializada"
var status bool

//Si la variable comienza con mayúscula, es una variable pública
var VarPublica string

func main() {
	//Go me inicializa todas las variables
	//string = ""
	//numbers = 0
	//booleans = false
	fmt.Println(numero, texto, status)

	var numero2 int
	numero3 := 4
	numero3 = 5
	var num1, num2, num3 int
	num1, num2, num3 = 1, 2, 3
	text1, num4, bool1 := "esto", 5, false
	fmt.Println(numero2, numero3)
	fmt.Println(num1, num2, num3)
	fmt.Println(text1, num4, bool1)

	//Demostración del scope
	//Puedo crear variables locales con los mismos nombres que los que están por fuera
	var status bool = true
	fmt.Println("status local", status)
	verStatusGlobal()

	var entero int = 5
	var flotante float32 = 10
	//Hay que tener cuidado al convertir o igualar números, ya que para Go los enteros y los flotantes no son iguales. antes hay que convertir la variable para poder igaualar
	entero = int(flotante)
	fmt.Println(entero)

	//Si se quiere convertir de numero a texto, no se puede hacer así nomás. Lo deja como cadena vacía por ejemplo
	var textoAnum string = ""
	//Ver bien estos métodos para convertit
	textoAnum = fmt.Sprintf("%d", entero)
	fmt.Println(textoAnum)

	//Otra Forma de convertir con otro paquete (strconv)
	textoAnum = strconv.Itoa(entero)
	println(textoAnum)
}

func verStatusGlobal() {
	fmt.Println("status global", status)
}
