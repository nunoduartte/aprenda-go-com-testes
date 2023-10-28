package idioma

type Frances struct {
}

func (f Frances) Ola(nome string) string {
	return "Bonjour, " + nome
}
