package user

import "time"

//Definición de un objeto
//En mayúscula para que lo pueda usar desde afuera (lo exporta)
type Usuario struct {
	Id        int
	Nombre    string
	FechaAlta time.Time
	Status    bool
}

//Agregamos un método o función
//Tipo void, no devuelve nada
//Es una función que hace referencia al modulo que apunta
//Hace de puntero al módulo
func (this *Usuario) AltaUsuario(id int, nombre string, fechaalta time.Time, status bool) {
	//Tambien se puede refenciar así, pero no conviene u:=&Usuario
	//Usar siempre this
	this.Id = id
	this.Nombre = nombre
	this.FechaAlta = fechaalta
	this.Status = status
}
