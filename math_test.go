package main

import (
	"bytes"
	"fmt"
	"io"
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
	r, w, _ := os.Pipe()
	stdout := os.Stdout
	os.Stdout = w

	// Chama a função main
	main()

	// Fecha o writer e restaura a saída padrão
	w.Close()
	os.Stdout = stdout

	// Lê a saída capturada
	var buf bytes.Buffer
	io.Copy(&buf, r)

	// Verifica a saída
	expected := fmt.Sprintln(soma(111, 10))
	if buf.String() != expected {
		t.Errorf("Saída da função main é inválida: Resultado %q. Esperado: %q", buf.String(), expected)
	}
}
