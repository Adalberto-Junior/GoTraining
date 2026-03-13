package main
import ("fmt")

func main () {
	var student1 string = "Adalberto" // type is string
	var student2 = "Esha" //type is inferret
	age := 24  // type is inferret
	var relaction string
	var years int
	var validate bool 


	fmt.Println(student1)
	fmt.Println(student2)
	fmt.Println(age)
	fmt.Println("relaction",relaction)
	fmt.Println(years)
	fmt.Println(validate)
	fmt.Println("After Give a value")
	relaction = "Friend"
	fmt.Println("relaction: ",relaction)

	var x,y,z int = 1,2,3
	var a,nome = 15, "Ana"
	
	fmt.Println("x: ",x)
	fmt.Println("y: ",y)
	fmt.Println("z: ",z)
	fmt.Println("a: ",a)
	fmt.Println("nome: ",nome)

	const PI  = 3.14

	fmt.Println("PI: ", PI)
}