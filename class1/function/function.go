package main
import ("fmt")

func isEven (num int) bool {
	// if num % 2 == 0 {
	// 	return true
	// } else {
	// 	return false
	// }
	return num % 2 == 0
}

func addNumb (x int, y int) (add int) {
	add = x + y

	return 
}

func multOf12 (x int) int {
	if x > 12 {
		return 1
	}
	fmt.Println(x," * 12 = ",x*12)

	return multOf12((x +1))
}

func main () {
	fmt.Println("O número: 2 é par? ", isEven(2))
	fmt.Println("O número: 533 é par? ", isEven(533))
	fmt.Println("O número: 334 é par? ", isEven(334))

	fmt.Println("3 + 45 =  ", addNumb(3,45))
	fmt.Println("10 + 13 = ", addNumb(10,13))
	fmt.Println("334 + 234 =  ", addNumb(334,234))

	multOf12(1)
}

