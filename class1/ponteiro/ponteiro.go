package main
import "fmt"

type Creature struct {
	name string
	especie string
}

type Aluno struct{
	name string
}

func (a Aluno) Reset() {
	a.name = ""
}

func (a *Aluno) Update(name string) {
	a.name = name
}

func main () {
	var number int = 24
	var mypointer *int = &number

	fmt.Println("Number: ",number)
	fmt.Println("Number endereço: ",&number)
	fmt.Println("Endereço que o ponteiro aponta: ",mypointer)
	fmt.Println("Endereço do pointer?: ",&mypointer)
	fmt.Println("O conteudo do pointer?: ",*mypointer)

	*mypointer = 15
	fmt.Println("Depois de mudar o valor do ponteiro:")
	fmt.Println("Number: ",number)
	fmt.Println("Number endereço: ",&number)
	fmt.Println("Endereço que o ponteiro aponta: ",mypointer)
	fmt.Println("Endereço do pointer?: ",&mypointer)
	fmt.Println("O conteudo do pointer?: ",*mypointer)

	var creature Creature

	 creature.name = "Golfinho"
	 creature.especie = "Shark"

	fmt.Println("Antes da função: \n",creature.name)
	changeCreatureName(&creature)
	fmt.Println("Depois da função: \n",creature.name)

	var aluno Aluno = Aluno{name: "João"}
	fmt.Println("Passagem por valor: ")
	fmt.Println("Aluno Name: ",aluno.name)
	aluno.Reset()
	fmt.Println("Depois de Reset: ",aluno.name)

	fmt.Println("Passagem por pointer: ")

	var student Aluno = Aluno{name: "Adalberto"}
	fmt.Println("Aluno Name: ",student.name)
	student.Update("Júnior")
	fmt.Println("Depois de Updete: ",student.name)
}

func changeCreatureName (creature *Creature) {
	creature.name = "Tubarão"
	fmt.Println(creature.name)
}