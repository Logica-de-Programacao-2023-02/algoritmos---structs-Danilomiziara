package main

import (
	"fmt"
	"math"
)

type Circulo struct {
	raio float64
}

func main() {
	c := Circulo{raio: 5}
	area := Area(c)
	fmt.Printf("Área do círculo: %.2f\n", area)
}

func Area(c Circulo) float64 {
	return math.Pi * c.raio * c.raio
}
