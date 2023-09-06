package main

import "fmt"

func nIntAlvo(v []int, n int, alvo int) (int, int) {

	for i := 0; i <= n-2; i++ {
		for j := 0; j <= n-1; j++ {
			if v[i]+v[j] == alvo {
				return v[i], v[j]
			} //fim do indice j
		}

	} //Fim  indice i
	return -1, -1
} //fim NintAlvo

func main() {
	v := []int{1, 2, 3}

	fmt.Println(nIntAlvo(v, 3, 5))
	fmt.Println(nIntAlvo(v, 3, 1))

}
