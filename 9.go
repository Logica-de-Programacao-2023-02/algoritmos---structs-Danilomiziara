package main

import "fmt"

type Aluno struct {
	nome  string
	idade int
	notas []float64
}

func main() {
	aluno := Aluno{
		nome:  "Pedro Felipe",
		idade: 16,
		notas: []float64{9.9, 8.0, 2.5},
	}

	aluno.imprimirInformacoes()

	// Adicionar uma nova nota
	aluno.adicionarNota(8.0)

	// Remover a segunda nota
	aluno.removerNota(1)

	aluno.imprimirInformacoes()
}

func (a Aluno) adicionarNota(nota float64) {
	a.notas = append(a.notas, nota)
}

func (a Aluno) removerNota(indice int) {
	if indice >= 0 && indice < len(a.notas) {
		a.notas = append(a.notas[:indice], a.notas[indice+1:]...)
	}
}

func (a Aluno) calcularMedia() float64 {
	total := 0.0
	for _, nota := range a.notas {
		total += nota
	}
	return total / float64(len(a.notas))
}

func (a Aluno) imprimirInformacoes() {
	fmt.Println("Nome:", a.nome)
	fmt.Println("Idade:", a.idade)
	fmt.Println("Média:", a.calcularMedia())
}
