package main
import "fmt"

func main () {

	var aluno = map[string]string{"name": "Adalberto","mec": "118923", "cours":"MIECT","years":"5 ano"}
	universidade := map[string]int{"UA": 2001,"Losiada":1999,"Nova Lisboa":1990,"UC":1975}

	fmt.Println("Map:")
	fmt.Println("aluno: \n",aluno)
	fmt.Println("Universidade: \n",universidade)

	var cidade = make(map[string]string)  //is empty

	cidade["Cidede de Aveiro"] = "Aveiro"
	cidade["Albergaria-a-velha"] ="Aveiro"
	cidade["Aguida"] = "Aveiro"
	cidade["Coimbra"] = "Coimbra"

	fmt.Println("Com make(): \n",cidade)

	var pais = make(map[string]string)
	var continente map[string]string

	fmt.Println(pais == nil)
	fmt.Println(continente == nil)

	fmt.Println("Pais: \n",pais)
	fmt.Println("Continente: \n",continente)
}