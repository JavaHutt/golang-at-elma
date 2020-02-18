package main

import "fmt"

// FIO struct
type FIO struct {
	firstName, lastName, secondName string
}

//ExponentialSearch 123
func ExponentialSearch(a []int, x int) int {
	if len(a) == 1 {
		if a[0] > x {
			return 0
		}
		return 1
	}
	k := 1
	for ; k < len(a) && a[len(a)-k] > x; k *= 2 {

	}
	if k == 1 {
		return len(a)
	}
	k /= 2
	for ; k < len(a) && a[len(a)-k] > x; k++ {
	}
	if k == len(a) {
		return 0
	}
	return len(a) - k + 1
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
