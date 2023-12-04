package main

import (
	"fmt"
	"time"
)

type Musica struct {
	titulo  string
	artista string
	duracao time.Duration
}

type Playlist struct {
	nome    string
	musicas []Musica
}

func main() {
	musicas := []Musica{
		{titulo: "Música 1", artista: "Artista 1", duracao: 3 * time.Minute},
		{titulo: "Música 2", artista: "Artista 2", duracao: 4 * time.Minute},
		{titulo: "Música 3", artista: "Artista 3", duracao: 5 * time.Minute},
	}

	playlist := Playlist{
		nome:    "Minha Playlist",
		musicas: musicas,
	}

	imprimirPlaylist(playlist)
}

func imprimirPlaylist(pl Playlist) {
	fmt.Println("Nome da Playlist:", pl.nome)

	tempoTotal := TempoTotal(pl.musicas)
	fmt.Println("Tempo Total da Playlist:", tempoTotal)

	fmt.Println("Músicas:")
	for _, musica := range pl.musicas {
		fmt.Println("Título:", musica.titulo)
		fmt.Println("Artista:", musica.artista)
		fmt.Println("Duração:", musica.duracao)
		fmt.Println()
	}
}

func TempoTotal(musicas []Musica) time.Duration {
	var total time.Duration
	for _, musica := range musicas {
		total += musica.duracao
	}
	return total
}
