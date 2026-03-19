package main

import (
	"fmt"
	"time"
)


func sayHello() {
	for i := 0; i < 3; i++ {
		fmt.Println("Hello!")
		time.Sleep(100 * time.Microsecond)
	}
}

func sayWorld(){
	for i:=0; i < 3; i++ {
		fmt.Println("World!")
		time.Sleep(100 * time.Microsecond)
	}
}

// func main() {

// 	go sayHello()
// 	go sayWorld()

// 	time.Sleep(1 * time.Second)

// 	fmt.Println("End!")
// }

// func main() {

// 	go func() {
// 		for i:= 0; i < 3; i++ {
// 			fmt.Println("Goroutine: ",i)
// 			time.Sleep(50 * time.Millisecond)
// 		}
// 	}()

// 	for i := 0; i < 3; i++ {
// 		fmt.Println("main: ",i)
// 		time.Sleep(50 * time.Millisecond)
// 	}

// 	time.Sleep(200 * time.Millisecond)
// }


func process(id int, data string) {
	fmt.Printf("Gorotine %d: Processing: %s\n",id, data)
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("Gorotine %d: ended\n",id)
}


// func main() {

// 	for i := 1; i <= 5; i++ {
// 		go process(i, fmt.Sprintf("Tarefa-%d",i))
// 	}

// 	time.Sleep(1 * time.Second)
// 	fmt.Println("End of all Gorotine")
// }