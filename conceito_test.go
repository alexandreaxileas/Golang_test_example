package conceito

import "testing"

const erroRetorno = "Erro ao retornar o Conceito. Valor (%v) não aceito"

func TestConceitoAluno(t *testing.T) {
	//
	t.Parallel()
	nota := []struct {
		valor float64
	}{
		{9.0},
		{8.8},
		{4.9},
		{11.5},
		{1.5},
		{6.5},
		{3.5},
	}
	//
	for _, n := range nota {
		//
		t.Logf("Nota testada %f", n)
		c := ConceitoAluno(n.valor)
		//
		if (c != "A") && (c != "B") && (c != "C") && (c != "D") &&
			(c != "E") && (c != "F") && (c != "Inválido") {
			t.Errorf(erroRetorno, "de A a F ou Inválido") // The quantity of parameters it is relation with the const used erroRetorno
			// Realize that erroRetorno has just one "%"
		}
	}

}

func TestConceitoVariosAluno(t *testing.T) {
	//Criando o slice
	vAluno1 := make([]Alunos, 0) //
	//Populando a variável vAluno
	vAluno1 = append(vAluno1, Alunos{"Alexandre G.", 2.90})
	vAluno1 = append(vAluno1, Alunos{"Alexandre", 10.00})
	vAluno1 = append(vAluno1, Alunos{"Felipe", 9.00})
	vAluno1 = append(vAluno1, Alunos{"Alini", 1.50})
	vAluno1 = append(vAluno1, Alunos{"Natália", 7.00})
	//
	ConceitoVariosAluno(vAluno1)
}
