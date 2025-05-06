package main

import "fmt"

const prefixoOlaPortugues = "Olá, "

func SelecionarPrefixo(idioma string) string {
	if idioma != "portugues" {
		switch idioma {
		case "espanhol":
			return "Hola, "
		default:
			return prefixoOlaPortugues
		}
	}

	return prefixoOlaPortugues
}

func Ola(nome string, idioma string) string {
	prefixo := SelecionarPrefixo(idioma)

	if nome == "" {
		nome = "Mundo"
	}

	return prefixo + nome
}

func main() {
	fmt.Println(Ola("Gabriel", "portugues"))
	fmt.Println(Ola("Gabriel", "espanhol"))
	fmt.Println(Ola("Gabriel", "ingles"))
}
