package main

import (
	"fmt"
	"math"
)

type Pessoa struct {
	Nome   string
	Idade  int
	Altura float64
}

func (p *Pessoa) AvancaIdade() {
	p.Idade++
}

type Ponto struct {
	X, Y int
}

type Retangulo struct {
	Ponto
	Largura, Altura int
}

type Circulo struct {
	Raio float64
}

func (c Circulo) Area() float64 {
	return math.Pi * c.Raio * c.Raio
}

func main() {
	circ := Circulo{
		Raio: 1.45,
	}
	fmt.Println(circ.Area())

	p := Pessoa{
		Nome:   "Alice",
		Idade:  25,
		Altura: 1.65,
	}
	fmt.Println(p)
	fmt.Println(p.Nome)
	fmt.Println(p.Altura)

	p.AvancaIdade()
	fmt.Println(p)

	r := Retangulo{
		Ponto:   Ponto{X: 2, Y: 4},
		Largura: 60,
		Altura:  30,
	}
	fmt.Println(r)
}
