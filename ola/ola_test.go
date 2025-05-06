package main

import "testing"

func TestOla(t *testing.T) {
	verificaMensagemCorreta := func(t *testing.T, resultado, esperado string) {
		t.Helper()
		if resultado != esperado {
			t.Errorf("resultdo '%s', esperado '%s'", resultado, esperado)
		}
	}

	t.Run("diz olá para as pessoas", func(t *testing.T) {
		resultado := Ola("Chris", "portugues")
		esperado := "Olá, Chris"
		verificaMensagemCorreta(t, resultado, esperado)
	})

	t.Run("diz 'Olá, Mundo' quando uma string vazia for passada", func(t *testing.T) {
		resultado := Ola("", "portugues")
		esperado := "Olá, Mundo"
		verificaMensagemCorreta(t, resultado, esperado)
	})

	t.Run("em espanhol", func(t *testing.T) {
		resultado := Ola("Elodie", "espanhol")
		esperado := "Hola, Elodie"
		verificaMensagemCorreta(t, resultado, esperado)
	})

	t.Run("idioma desconhecido retorna mensagem com prefixo em portugues", func(t *testing.T) {
		resultado := Ola("Elodie", "ingles")
		esperado := "Olá, Elodie"
		verificaMensagemCorreta(t, resultado, esperado)
	})
}

func TestSelecionarPrefixo(t *testing.T) {
	verificarPrefixoCorreto := func(t *testing.T, resultado, esperado string) {
		t.Helper()
		if resultado != esperado {
			t.Errorf("resultdo '%s', esperado '%s'", resultado, esperado)
		}
	}

	t.Run("retorna profixo em portugues", func(t *testing.T) {
		resultado := SelecionarPrefixo("portugues")
		esperado := "Olá, "
		verificarPrefixoCorreto(t, resultado, esperado)
	})

	t.Run("retorna prefixo em espanhol", func(t *testing.T) {
		resultado := SelecionarPrefixo("espanhol")
		esperado := "Hola, "
		verificarPrefixoCorreto(t, resultado, esperado)
	})

	t.Run("retorna prefixo com idioma desconhecido retorna em portugues", func(t *testing.T) {
		resultado := SelecionarPrefixo("ingles")
		esperado := "Olá, "
		verificarPrefixoCorreto(t, resultado, esperado)
	})
}
