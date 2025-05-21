package main

import "fmt"

const (
	c0 = iota
	c1 = iota
	c2 = iota
)

const (
	a = iota
	b
	c
	d
	e
	f
)

func main() {
	fmt.Println(c0, c1, c2)
	fmt.Printf("%d \t %b\n", a, a)
	fmt.Printf("%d \t %b\n", 1<<a, 1<<a)
	fmt.Printf("%d \t %b\n", 1<<b, 1<<b)
	fmt.Printf("%d \t %b\n", 1<<c, 1<<c)
	fmt.Printf("%d \t %b\n", 1<<d, 1<<d)
	fmt.Printf("%d \t %b\n", 1<<e, 1<<e)
	fmt.Printf("%d \t %b\n", 1<<f, 1<<f)
}
