package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var x int
	var y float64

	fmt.Println("Informe o valor: ")
	fmt.Scanln(&x)
	fmt.Println(x)

	fmt.Println("Informe um float: ")
	fmt.Scanln(&y)
	fmt.Println(y)

	leitor := bufio.NewReader(os.Stdin)
	fmt.Println("Escreva uma mensagem:")
	msg, _ := leitor.ReadString('\n')

	fmt.Println(msg)
}
