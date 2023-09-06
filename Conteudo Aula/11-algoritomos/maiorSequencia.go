package main

import "fmt"

func maiorSequencia(v []int) int {
	n := len(v)
	seq := 1
	mSeq := 1
	for i := 0; i < n-1; i++ {
		if v[i+1] > v[i] {
			seq++
		} else {
			if seq > mSeq {
				mSeq = seq
			}
			seq = 1
		}
	}
	if seq > mSeq {
		mSeq = seq
	}
	return mSeq
}

func main() {
	a := []int{1, 2, 3, 1, 2}
	fmt.Println(maiorSequencia(a))
}
