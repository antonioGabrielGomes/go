package iteracao

var quantidadeRepeticoes = 5

func Repetir(caractere string, quantidadeRepeticao int) string {
	var repeticoes string

	if quantidadeRepeticao != 0 {
		quantidadeRepeticoes = quantidadeRepeticao
	}

	for i := 0; i < quantidadeRepeticoes; i++ {
		repeticoes += caractere
	}
	return repeticoes
}
