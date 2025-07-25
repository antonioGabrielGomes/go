package main

import "math"

// retangulo
type Retangulo struct {
	Largura float64
	Altura float64
}

func (r Retangulo) Area() float64 {
	return r.Largura * r.Altura
}

// circulo
type Circulo struct {
	Raio float64
}

func (c Circulo) Area() float64 {
	return math.Pi * c.Raio * c.Raio
}

// triangulo
type Triangulo struct {
	Base float64
	Altura float64
}

func (t Triangulo) Area() float64 {
	return (t.Base * t.Altura) * 0.5
}

// perimetro
func Perimetro(retangulo Retangulo) float64 {
	return 2 * (retangulo.Largura + retangulo.Altura)
}

// forma
type Forma interface {
	Area() float64
}