package main

import "fmt"

var linguagem string

func main() {
	// declaração explícita
	var msg string
	var n1, n2 int

	fmt.Println(msg, n1, n2)

	// declaração implícita
	var n3, n4 = 4, 5

	fmt.Println(n3, n4)

	// declaração curta
	n5 := 10

	fmt.Println(n5)

	const Pi = 3.14

	var n6 complex128

	n6 = 14 + 30.4i

	fmt.Println(n6)
}

/*

INTEIROS
int8	-128 a 127
int16	-32.768 a 32.767
int32
int64
uint8	0 a 255
uint16	0 a 65.535
uint32
uint64

byte	uint8
rune	int32 -> caracteres unicode

int		de acordo com a arq. do computador (32 ou 64)

PONTO FLUTUANTE
float32	6 a 9 dígitos de precisão
float64	15 a 17 dígitos de precisão

quando a declaração é implícita -> float64

COMPLEXOS
complex64
complex128

DEMAIS
bool	false / true
string

nil		-> não possui valor válido
*/
