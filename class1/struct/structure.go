package main
import "fmt"

type Aluno struct {
	name string
	age int
	mec int
	year string
	email string
	curso string
}

func main () {

	var aluno1 Aluno
	var aluno2 Aluno

	aluno1.name = "Adalberto"
	aluno1.email = "ada@uni.com"
	aluno1.mec = 120967
	aluno1.year = "2º ano"
	aluno1.age =  21
	aluno1.curso = "MIECT"

	
	aluno2.name = "Junior"
	aluno2.email = "juni@uni.com"
	aluno2.mec = 120968
	aluno2.year = "2º ano"
	aluno2.age =  21
	aluno2.curso = "LCI"

	fmt.Println("Dados dos alunos: ")
	fmt.Println("Aluno 1:")
	fmt.Println("Nome: ",aluno1.name)
	fmt.Println("email: ",aluno1.email)
	fmt.Println("mec: ",aluno1.mec)
	fmt.Println("Idade: ", aluno1.age)
	fmt.Println("Curso: ",aluno1.curso)
	fmt.Println("Ano: ",aluno1.year)

	fmt.Println("\nAluno 2:")
	fmt.Println("Nome: ",aluno2.name)
	fmt.Println("email: ",aluno2.email)
	fmt.Println("mec: ",aluno2.mec)
	fmt.Println("Idade: ", aluno2.age)
	fmt.Println("Curso: ",aluno2.curso)
	fmt.Println("Ano: ",aluno2.year)

	fmt.Println("\n\nUsing Fuction: ")
	printStructData(aluno1)
	printStructData(aluno2)

}

func printStructData (aluno Aluno) {
	fmt.Println("Nome: ",aluno.name)
	fmt.Println("email: ",aluno.email)
	fmt.Println("mec: ",aluno.mec)
	fmt.Println("Idade: ", aluno.age)
	fmt.Println("Curso: ",aluno.curso)
	fmt.Println("Ano: ",aluno.year)
}