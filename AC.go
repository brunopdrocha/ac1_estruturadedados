package main

import "fmt"

//Numero Primo
func e_primo(n int) bool {
	var t bool
	t = true
	for i := 2; i <= n-1; i++ {
		if n%i == 0 {
			t = false
			println(i)
		}
	}
	return t
}

//Elemento Fibonnaci
func fibo(n int) int {

	if n < 2 {
		return 1
	}
	return fibo(n-1) + fibo(n-2)
}

//Funcao numero informa dia da semana

func dia(n int) string {

	if n == 1 {
		return "Domingo"
	}

	if n == 2 {
		return "Segunda"
	}

	if n == 3 {
		return "Terca"
	}
	if n == 4 {
		return "Quarta"
	}

	if n == 5 {
		return "Quinta"

	}

	if n == 6 {
		return "Sexta"
	}

	if n == 7 {
		return "Sabado"
	} else {
		return "Valor Invalido"
	}
}

//Verificar se determinado ano eh bisexto

func e_bissexto(n int) string {
	n %= 4
	println(n)
	if n == 0 {
		return "Ano Bissexto"
	} else {
		return "N eh ano Bissexto"
	}
}

func main() {

	fmt.Println(e_primo(7))  //true
	fmt.Println(e_primo(10)) //false

	fmt.Println("=-=-=-===-=-=-=-=--=-=-")

	fmt.Println(fibo(1)) //1
	fmt.Println(fibo(4)) //5

	fmt.Println("=-=-=-===-=-=-=-=--=-=-")

	fmt.Println(dia(3))  //Terca
	fmt.Println(dia(10)) //Invalido

	fmt.Println("=-=-=-===-=-=-=-=--=-=-")

	fmt.Println(e_bissexto(1995))
	fmt.Println(e_bissexto(2016))
}
