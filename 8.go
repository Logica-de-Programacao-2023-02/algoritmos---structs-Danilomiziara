package main

import (
	"fmt"
	"time"
)

type Viagem struct {
	Origem  string
	Destino string
	Data    time.Time
	Preco   float64
}

func main() {
	viagens := []Viagem{
		{Origem: "São Paulo", Destino: "Rio de Janeiro", Data: time.Now(), Preco: 300.0},
		{Origem: "São Paulo", Destino: "Salvador", Data: time.Now(), Preco: 500.0},
		{Origem: "São Paulo", Destino: "Recife", Data: time.Now(), Preco: 400.0},
	}

	fmt.Println(viagemMaisCara(viagens))
}

func viagemMaisCara(viagens []Viagem) Viagem {
	var maisCara Viagem
	for i, viagem := range viagens {
		if i == 0 || viagem.Preco > maisCara.Preco {
			maisCara = viagem
		}
	}
	return maisCara
}
