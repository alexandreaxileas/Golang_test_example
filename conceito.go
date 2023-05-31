package conceito

// ConceitoAluno retorna o conceito a partir da nota
func ConceitoAluno(nota float64) string {
	switch {
	case (nota <= 10) && (nota >= 9):
		return "A"
	case (nota < 9) && (nota >= 7):
		return "B"
	case (nota < 7) && (nota >= 6):
		return "C"
	case (nota < 6) && (nota >= 4):
		return "D"
	case (nota < 4) && (nota >= 3):
		return "E"
	case (nota < 3) && (nota >= 0):
		return "F"
	default:
		return "Inválido"
	}

}
