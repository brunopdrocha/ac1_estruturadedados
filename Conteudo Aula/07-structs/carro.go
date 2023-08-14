package main

import "fmt"

type Carro struct {
	Modelo     string
	Cor        string
	Velocidade int
	EstaLigado bool
}

func (c *Carro) Ligar() {
	c.EstaLigado = true
}

func (c *Carro) Desligar() {
	c.EstaLigado = false
	c.Velocidade = 0
}
func (c *Carro) Acelerar(valor int) {
	c.Velocidade += valor
}

func main() {
	C := Carro{
		Modelo:     "Celta",
		Cor:        "Azul",
		Velocidade: 120,
		EstaLigado: false,
	}

	fmt.Println(C.Modelo)
	fmt.Println(C.Cor)
	fmt.Println(C.Velocidade)

	C.Ligar()
	fmt.Println(C)

	C.Acelerar(15)
	fmt.Println(C)

	C.Desligar()
	fmt.Println(C)

}
