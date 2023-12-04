package main

import "fmt"

type Animal struct {
	nome    string
	espécie string
	idade   int
	som     string
}

func (a *Animal) modificarSom(novoSom string) {
	a.som = novoSom
}

func (a Animal) imprimirInformações() {
	fmt.Println("Nome:", a.nome)
	fmt.Println("Espécie:", a.espécie)
	fmt.Println("Idade:", a.idade)
	fmt.Println("Som:", a.som)
}

func main() {
	animal := Animal{
		nome:    "galinha",
		espécie: "ave",
		idade:   2,
		som:     "popo",
	}

	animal.imprimirInformações()

	// Modificar o som do animal
	animal.modificarSom("bocejo")

	fmt.Println("Novo som:", animal.som)
}
