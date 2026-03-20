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
	data, err := readDoc("oQueEstudar.txt")

	if err != nil {
		fmt.Errorf("Falha ao inicializar data: %w",err)
	}
	fmt.Println("Data: ",data)

	return nil
}

type Produto struct {
	name string
	price int
	stock int
}

func validateProduct(p Produto) error{
	var erro []error

	if p.name == ""{
		erro = append(erro, fmt.Errorf("nome é obrigatório"))
	}

	if p.price < 0 {
		erro = append(erro, fmt.Errorf("O preço deve ser positivo"))
	}

	if p.stock < 0 {
		erro = append(erro, fmt.Errorf("estoque não pode ser negativo"))
	}

	return errors.Join(erro...)
}

func main() {

	result, err := dividerNumb(2,0)

	if err != nil {
		fmt.Println("Erro: ",err)
		// return
	}

	fmt.Println("Resultado: ",result)

	errs := runData()

	if errs != nil {
		fmt.Println(err)
	}

	p := Produto{name: "",price: -120, stock: -15}

	erro := validateProduct(p)

	if erro != nil {	
		fmt.Println(erro)
	}
	
}