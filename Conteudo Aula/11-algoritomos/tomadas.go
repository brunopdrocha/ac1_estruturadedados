package main

import "fmt"

func main() {
	var t1, t2, t3, t4 int

	fmt.Scan(&t1, &t2, &t3, &t4)
	totalTomadas := t1 + t2 + t3 + t4

	fmt.Println(totalTomadas - 3)
}
