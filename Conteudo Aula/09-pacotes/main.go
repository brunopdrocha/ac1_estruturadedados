package main

import (
	"fmt"
	"projetinho/utils"
	"projetinho/utils/outroutils"
)

func main() {
	fmt.Println(utils.Somar(4.3, 6.2))
	fmt.Println(utils.Sub(4.3, 6.2))
	fmt.Println(utils.Multi(8, 1.2))
	fmt.Println(outroutils.Div(8, 1.1))
}
