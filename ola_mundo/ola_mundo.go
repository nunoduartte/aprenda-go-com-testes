package ola_mundo

import "aprendendo-go-com-testes/internal/ola_mundo/idioma"

func OlaPessoa(idioma idioma.Idioma, nome string) string {
	return idioma.Ola(nome)
}
