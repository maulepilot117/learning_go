package main

import "fmt"

func main() {
	adams := 42
	fmt.Printf("42 as binary, %b \n", adams)
	fmt.Printf("42 as hex, %x \n", adams)

	a, b, c, d, e, f := 0, 1, 56734, 3, 4, 5
	fmt.Printf(`
    Values as binary:
    a: %b,
    b: %b,
    c: %b,
    d: %b,
    e: %b,
    f: %b`, a, b, c, d, e, f)
	fmt.Printf(`
    Values as decimal:
    a: %d,
    b: %d,
    c: %d,
    d: %d,
    e: %d,
    f: %d`, a, b, c, d, e, f)
}
