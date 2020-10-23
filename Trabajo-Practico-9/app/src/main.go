package main

import "fmt"

// Msj : Simple Example
func Msj() string {
	return "Hello, Word."
}

// MsjWithParams : Example with params
func MsjWithParams(a float64, b float64) float64 {
	return a * b
}

// WithIf : Example with confitionals
func WithIf(a string) string {
	lenght := len(a)
	if lenght > 10 {
		return "Fizz"
	} else if lenght > 5 {
		return "Bozz"
	} else {
		return "Perfect"
	}
}

func main() {
	a := Msj()
	fmt.Println(a)
	b := MsjWithParams(2, 4)
	fmt.Println(b)
	c := WithIf("FizzBozz")
	fmt.Println(c)
}
