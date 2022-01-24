package main

import (
	"fmt"
	"time"

	us "./user"
)

//Herencia
type pepe struct {
	us.Usuario
}

//Definición de un objeto
func main() {

	u := new(pepe)
	u.AltaUsuario(1, "Javier", time.Now(), true)
	fmt.Println(u.Usuario)
}
