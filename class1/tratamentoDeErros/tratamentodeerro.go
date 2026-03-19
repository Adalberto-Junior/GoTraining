package main

import(
	"fmt"
	"errors"
	"os"
)

func dividerNumb (a,b float64) (float64, error) {
	
	if b == 0 {
		return 0, errors.New("Divisão por zero não suportada")
	}
	return  a/b, nil
}

var fileNotFound = errors.New("Arquivo não encontrado")

func readDoc(path string) ([]byte, error) {
	data, err := os.ReadFile(path)

	if err != nil {
		return nil, fmt.Errorf("ao ler arquivo %s: %w",path,err)
	}
	return  data, nil
}

func runData() error {
	data, err := readDoc("oueEstudar.txt")

	if err != nil {
		fmt.Errorf("Falha ao inicializar data: %w",err)
	}
	fmt.Println("Data: ",data)

	return nil
}

func main() {

	result, err := dividerNumb(2,0)

	if err != nil {
		fmt.Println("Erro: ",err)
		return
	}

	fmt.Println("Resultado: ",result)

	errs := runData()

	if errs != nil {
		fmt.Println(err)
	}
}