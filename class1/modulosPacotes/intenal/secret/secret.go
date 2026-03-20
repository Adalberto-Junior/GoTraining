package secret

//token não é exportado (minúsculo)

func token() string {
	return "abc-123"
}

//TokenForLog expõe algo controlado para uso interno do módulo.
func TokenForLog() string {
	return token()
}
