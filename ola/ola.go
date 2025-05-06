package main

import "fmt"

const portugues = "portugues"
const frances = "frances"
const espanhol = "espanhol"
const chines = "chines"
const prefixoOlaPortugues = "Olá, "
const prefixoOlaEspanhol = "Hola, "
const prefixoOlaFrances = "Bonjour, "
const prefixoOlaChines = "Nǐ hǎo, "

func selecionarPrefixo(idioma string) (prefixo string) {
	switch idioma {
	case espanhol:
		prefixo = prefixoOlaEspanhol
	case frances:
		prefixo = prefixoOlaFrances
	case chines:
		prefixo = prefixoOlaChines
	default:
		prefixo = prefixoOlaPortugues
	}
	return
}

func Ola(nome string, idioma string) string {
	if nome == "" {
		nome = "Mundo"
	}

	prefixo := selecionarPrefixo(idioma)

	return prefixo + nome
}

func main() {
	fmt.Println(Ola("Gabriel", portugues))
	fmt.Println(Ola("Gabriel", espanhol))
	fmt.Println(Ola("Gabriel", "ingles"))
	fmt.Println(Ola("Gabriel", frances))
}
