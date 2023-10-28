package main

import (
	"aprendendo-go-com-testes/internal/ola_mundo/idioma"
	"aprendendo-go-com-testes/ola_mundo"
	"fmt"
)

func main() {
	f := idioma.Frances{}
	fmt.Println(ola_mundo.OlaPessoa(f, "nuno"))
}
