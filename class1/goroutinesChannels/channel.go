package main

import (
	"fmt"
	"time"
)


func worker(id int, job <- chan int, result chan <- int){
	for j := range job{
		fmt.Printf("Worker %d: processing %d\n",id,j)
		time.Sleep(time.Second)
		result <- j *2

	}
}

func main() {

	message := make(chan string)

	go func() {
		message <- "Helo guy. How are you?"
	}()

	 sms := <- message

	fmt.Println(sms)

	data := make(chan string,2)

	data <- "Hello 1"
	data <- "Hello 2"

	fmt.Println(<- data)
	fmt.Println(<- data)


	const numbJob = 10
	const numWorker = 3

	job := make(chan int, numbJob)
	results := make(chan int, numbJob)

	for w := 1; w <= numWorker; w++ {
		go worker(w,job,results)
	}

	//enviar o trabalho
	for j := 1; j <= numbJob; j++ {
		job <- j
	}
	close(job)

	//coletar o resultado
	for a := 1; a <= numbJob; a++{
		result := <- results

		fmt.Printf("Resultado: %d\n",result)
	}



}