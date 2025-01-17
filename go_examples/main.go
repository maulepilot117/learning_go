package main

import "fmt"

func main() {
	const name, age = "Chris", 37
	fmt.Printf("%s is %d years old. 😂\n", name, age)
	
	fmt.Printf(`	This is a raw string literal. 😜
	We can put anything in here without newlines and escaping.
	It's pretty legit.`)
}
