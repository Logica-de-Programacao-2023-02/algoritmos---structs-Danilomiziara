package main

import "fmt"

type Triangulo struct {
	base   float64
	altura float64
}

func main() {
	triangulo := Triangulo{base: 5, altura: 3}
	area := AreaC(triangulo)
	fmt.Println("Área do triângulo:", area)
}

func AreaC(t Triangulo) float64 {
	area := t.base * t.altura / 2
	return area
}
