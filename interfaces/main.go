package main

import "fmt"

type serVivo interface {
	estaVivo() bool
}

type humano interface {
	respirar()
	pensar()
	comer()
	sexo() string
	estaVivo() bool
}

type animal interface {
	respirar()
	comer()
	esCarnivoro() bool
	estaVivo() bool
}

type vegetal interface {
	clasificacionVeg() string
	estaVivo() bool
}

/* Genero Humano */
type person struct {
	edad       int
	altura     float32
	peso       float32
	respirando bool
	pensando   bool
	comiendo   bool
	esHombre   bool
	vivo       bool
}

type hombre struct {
	person
}

type mujer struct {
	person
}

func (h *hombre) respirar() {
	h.respirando = true
}
func (h *hombre) pensar() {
	h.pensando = true
}
func (h *hombre) comer() {
	h.comiendo = true
}
func (h *hombre) sexo() string {
	if h.esHombre == true {
		return "Hombre"
	} else {
		return "Mujer"
	}
}
func (h *hombre) estaVivo() bool {
	return h.vivo
}

func (h *mujer) respirar() {
	h.respirando = true
}
func (h *mujer) pensar() {
	h.pensando = true
}
func (h *mujer) comer() {
	h.comiendo = true
}
func (h *mujer) sexo() string {
	if h.esHombre == true {
		return "Hombre"
	} else {
		return "Mujer"
	}
}
func (h *mujer) estaVivo() bool {
	return h.vivo
}

func HumanoRespirando(hu humano) {
	hu.respirar()
	fmt.Printf("Soy un/a %s, y ya estoy respirando \n", hu.sexo())
}

/* Genero Animal */

type perro struct {
	respirando bool
	comiendo   bool
	carnivoro  bool
	vivo       bool
}

func (p *perro) respirar() {
	p.respirando = true
}
func (p *perro) comer() {
	p.comiendo = true
}
func (p *perro) esCarnivoro() bool {
	return p.carnivoro
}
func AnimalRespirando(an animal) {
	an.respirar()
	fmt.Println("Soy un animal y estoy respirando")
}
func AnimalCarnivoro(an animal) int {
	if an.esCarnivoro() == true {
		return 1
	}
	return 0
}
func (an *perro) estaVivo() bool {
	return an.vivo
}

/*  Ser Vivo  */

func estoyVivo(v serVivo) bool {
	return v.estaVivo()
}

func main() {
	totalcarnivoros := 0
	Dogo := new(perro)
	Dogo.carnivoro = true
	AnimalRespirando(Dogo)
	Dogo.vivo = true
	totalcarnivoros += AnimalCarnivoro(Dogo)
	fmt.Printf("Total carnivoros %d \n", totalcarnivoros)
	fmt.Printf("Estoy vivo = %t \n", estoyVivo(Dogo))
}

func personas() {
	Pedro := new(hombre)
	Pedro.esHombre = true
	HumanoRespirando(Pedro)
	Maria := new(mujer)
	HumanoRespirando(Maria)
}
