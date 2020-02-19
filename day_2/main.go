package main

import "fmt"

// FIO struct
type FIO struct {
	firstName, lastName, secondName string
}

//ExponentialSearch 123
func ExponentialSearch(a []int, x int) int {
	for k, n := 0, len(a); ; n = n - k/2 {
		for k = 1; k <= n && a[n-k] > x; k *= 2 {
		}
		if k == 1 {
			return n
		}
	}

}

//BinarySearch 123
func BinarySearch(a []int, x int) int {
	l, r := 0, len(a)
	for l < r {
		if a[(r-l)/2] > x {
			r = (r + l) / 2
		} else {
			l = (r+l)/2 + 1
		}
	}
	return r
}

//InsertSort 123
func InsertSort(a []int) {
	for i := 1; i < len(a); i++ {
		j := ExponentialSearch(a[:i], a[i])
		buf := a[i]
		copy(a[j+1:i+1], a[j:i])
		a[j] = buf
	}
}

func main() {
	a := []int{12, 8, 22, 11, 1, 3}
	InsertSort(a)
	fmt.Println(a) // [1 3 8 11 12 22]
}
