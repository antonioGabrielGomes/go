package main

import "testing"

func TestPerimetro(t *testing.T) {
	retangulo := Retangulo{ 
		Largura: 10.0,
		Altura: 10.0,
	}
	resultado := Perimetro(retangulo)
	esperado := 40.0

	if resultado != esperado {
		t.Errorf("resultado %.2f esperado %.2f", resultado, esperado)
	}
}

/**
func TestArea(t *testing.T) {
	verificaArea := func(t *testing.T, forma Forma, esperado float64) {
		t.Helper()
		resultado := forma.Area()

		if resultado != esperado {
			t.Errorf("resultado %.2f, esperado %.2f", resultado, esperado)
		}
	}

	t.Run("retângulo", func(t *testing.T) {
		retangulo := Retangulo{ 
			Largura: 12.0,
			Altura: 6.0,
		}
		verificaArea(t, retangulo, 72.0)
	})

	t.Run("circulos", func(t *testing.T) {
		circulo := Circulo{
			Raio: 10,
		}
		verificaArea(t, circulo, 314.1592653589793)
	})
}
*/

func TestArea(t *testing.T) {
	testesArea := []struct {
		nome string
		forma Forma		
		temArea float64
	}{
		{
			nome: "Retângulo",
			forma: Retangulo{ 
				Largura: 12.0,
				Altura: 6.0,
			}, 
			temArea: 72.0,
		},
		{
			nome: "Circulo",
			forma: Circulo{
				Raio: 10,
			}, 
			temArea: 314.1592653589793,
		},
		{
			nome: "Triangulo",
			forma: Triangulo{
				Base: 12,
				Altura: 6,
			}, 
			temArea: 36.0,
		},
	}

	for _, tt := range testesArea {
		t.Run(tt.nome, func(t *testing.T) {
			resultado := tt.forma.Area()
			if resultado != tt.temArea {
				t.Errorf("resultado %.2f, esperado %.2f", resultado, tt.temArea)
			}
		})
	}
}