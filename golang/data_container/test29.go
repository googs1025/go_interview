package main

import "fmt"

const (
	x = iota
	y
	z = "xx"
	p
	q = iota
)

func main() {
	fmt.Println(x, y, z, p, q)
}
