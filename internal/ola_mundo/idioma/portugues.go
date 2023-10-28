package idioma

type Portugues struct {
}

func (p Portugues) Ola(nome string) string {
	return "Olá, " + nome
}
