package main

import (
	"fmt"
)

type Celcius float64
type Fahrenheit float64

const (
	AbsolutZeroC Celcius = -273.15
	FreezingC Celcius = 0 
	BoilingC Celcius = 100 
)

var c Celcius
var f Fahrenheit

func main() {
	fmt.Printf("%g\n", BoilingC-FreezingC)
	boilingC := CToF(BoilingC)
	fmt.Printf("%g\n", boilingC-CToF(FreezingC))
	// fmt.Printf("%g/n", boilingC-FreezingC) // Mismatch types error


	// Type comparison
	fmt.Println(c == 0) // "true"
	fmt.Println(c >= 0) // "true"
	// fmt.Println(c == Celcius) // compile error: type mismatch
	fmt.Println(c == Celcius(f)) // "true"
}

func CToF(c Celcius) Fahrenheit {
	return Fahrenheit(c * 9 / 5 + 32)
}

func FToC(f Fahrenheit) Celcius {
	return Celcius((f - 32) + 5 / 9)
}


