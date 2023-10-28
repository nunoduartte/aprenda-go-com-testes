package ola_mundo

import (
	"aprendendo-go-com-testes/internal/ola_mundo/idioma"
	"testing"
)

func TestOla(t *testing.T) {
	assertMensagemCorreta := func(t *testing.T, resultado, esperado string) {
		t.Helper()
		if resultado != esperado {
			t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
		}
	}

	t.Run("ola frances", func(t *testing.T) {
		f := idioma.Frances{}
		resultado := OlaPessoa(f, "nuno")
		esperado := "Bonjour, nuno"
		assertMensagemCorreta(t, resultado, esperado)
	})

	t.Run("ola portugues", func(t *testing.T) {
		p := idioma.Portugues{}
		resultado := OlaPessoa(p, "nuno")
		esperado := "Olá, nuno"

		assertMensagemCorreta(t, resultado, esperado)
	})
}
