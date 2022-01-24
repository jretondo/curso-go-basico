package main

import "fmt"

//Mapeados de elementos serializados
//Aparentemente son comos los objetos en JavaScript
//Un array de Objetos
func main() {
	paises := make(map[string]string, 5)
	fmt.Println(paises)
	paises["Mexico"] = "D.F."
	paises["Argentina"] = "BS As"
	fmt.Println(paises)

	campeonato := map[string]int{
		"Barcelona":   39,
		"Real Madrid": 10,
		"Chivas":      25,
		"Boca":        30}

	campeonato["River"] = 25
	campeonato["Chivas"] = 60
	//Lo ordena por clave (1-9)(A-Z) cuando lo imprimo en pantalla
	//Con el range lo hace desordenado

	//Borrar elemento
	delete(campeonato, "Real Madrid")

	fmt.Println(campeonato)

	//Recorrer mapa => con range lo ordena aleatoriamiente
	for equipo, puntaje := range campeonato {
		fmt.Printf("El equipo %s, tiene un puntaje de: %d\n", equipo, puntaje)
	}

	//Ver si existe un valor en mi mapa
	puntaje, existe := campeonato["Mineiro"]
	fmt.Printf("El puntaje capturado es %d, y el equipo existe %t\n", puntaje, existe)
}
