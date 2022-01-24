package main

import "fmt"

//Declaro el Array con la cantidad de largo y el tipo de valor que va a contener
//Se inicializa automaticamente con 10 posiciones en cero
var tabla [10]int
var matriz [5][7]int

func main() {
	Slices()
}

func Vectores() {
	tabla2 := [10]int{5, 6, 1, 2, 3, 4, 5, 9, 8, 7}
	fmt.Println(tabla2)
	//con la funcion "len" me da el largo del elemento
	for i := 0; i < len(tabla2); i++ {
		fmt.Println(i)
	}
}
func IniciodeVector() {
	tabla[0] = 1
	tabla[5] = 10
	fmt.Println(tabla)
}

func Matrices() {
	matriz[3][5] = 1
	fmt.Println(matriz)
}

//Slice me sirve para crear matrices y vectores de tamaño dinamico
func Slices() {
	matriz := []int{2, 5, 6}
	fmt.Println(matriz)
	variante4()
}

func variante2() {
	elementos := [5]int{1, 2, 3, 4, 5}
	//Estoy creando un slice con el vector elementos desde la posicion tres en adelante
	porcion := elementos[3:]

	porcion2 := elementos[:2]

	porcion3 := elementos[2:3]

	fmt.Println(porcion)
	fmt.Println(porcion2)
	fmt.Println(porcion3)
}

//Crear slices
func variante3() {
	//creo un slice con make, primero le paso el slice que voy a crear, después el valor y por último la capacidad, que en este caso es de 20 (puede contener hasta 20 elementos)
	//En resumidas cuentas, estoy creando un slice de tamaño 5, pero con capacidad de 20, me reserva la capacidad para este elemnto
	elemento := make([]int, 5, 20)

	//La funcion "cap()" me dice cual es la capacidad del elemento
	fmt.Printf("Largo %d, capacidad %d\n", len(elemento), cap(elemento))
}

func variante4() {
	//Creo un slice vacio => si o si voy a tener que crear un append
	nums := make([]int, 0, 0)
	for i := 0; i < 100; i++ {
		//Append me crea otro slice, pero cambiando su capacidad
		//Le agrega otro valor atrás (concatena) lo voy redimensionando
		nums = append(nums, i)
	}
	fmt.Printf("Largo %d, capacidad %d\n", len(nums), cap(nums))
	//terminal: Largo 100, capacidad 128
	//La capacidad autoasignada, siempre va a ser multiplo de ocho, si el valor no es pot de dos, la capacidad si lo va a ser y mayor que el valor (cap => 0 => 8 => 16 => 32 )
}
