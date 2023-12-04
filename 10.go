package main

import "fmt"

type Filme struct {
	Titulo     string
	Diretor    string
	Ano        int
	Avaliacoes []int
}

func main() {
	filme := Filme{Titulo: "Avatar 2", Diretor: "James Cameros", Ano: 0001}
	filme.adicionarAvaliacao(10)
	filme.adicionarAvaliacao(9)
	filme.adicionarAvaliacao(10)
	filme.informações()
}

func (f Filme) adicionarAvaliacao(nota int) {
	f.Avaliacoes = append(f.Avaliacoes, nota)
}

func (f Filme) removerAvaliacao(index int) {
	f.Avaliacoes = append(f.Avaliacoes[:index], f.Avaliacoes[index+1:]...)
}

func (f Filme) mediaAvaliacoes() float64 {
	var soma int
	for _, nota := range f.Avaliacoes {
		soma += nota
	}
	return float64(soma) / float64(len(f.Avaliacoes))
}

func (f Filme) informações() {
	fmt.Printf("Título: %s\n", f.Titulo)
	fmt.Printf("Diretor: %s\n", f.Diretor)
	fmt.Printf("Ano: %d\n", f.Ano)
	fmt.Printf("Média de avaliações: %.2f\n", f.mediaAvaliacoes())
}
