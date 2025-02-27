package main

import (
	"bytes"
	"fmt"
	"os"
	"testing"
)

func TestSoma(t *testing.T) {
	total := soma(15, 15)
	if total != 30 {
		t.Errorf("Resultado da soma é inválido: Resultado %d. Esperado: %d", total, 30)
	}
}

func TestMainFunction(t *testing.T) {
	// Captura a saída padrão
	var buf bytes.Buffer
	stdout := os.Stdout
	os.Stdout = &buf
	defer func() { os.Stdout = stdout }()

	// Chama a função main
	main()

	// Verifica a saída
	expected := fmt.Sprintln(soma(111, 10))
	if buf.String() != expected {
		t.Errorf("Saída da função main é inválida: Resultado %q. Esperado: %q", buf.String(), expected)
	}
}
