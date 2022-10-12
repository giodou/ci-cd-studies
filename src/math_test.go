package main

import "testing"

func TestSoma(t *testing.T){
	total := Soma(15,15)
	
	if total != 30 {
		t.Errorf("Resultado da soma é inválido: Resultado %d. Esperado: %d", total, 30)
	}
}

func TestSubtracao(t *testing.T){
	total := Subtracao(15,5)
	
	if total != 10 {
		t.Errorf("Resultado da Subtracao é inválido: Resultado %d. Esperado: %d", total, 30)
	}
}