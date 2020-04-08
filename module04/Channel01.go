package main


import(
	"sync"
	"fmt"
	"math/rand"
)

var(

	wg sync.WaitGroup

)

func main(){

	court := make(chan int)
	wg.Add(2)

	go player("vaibhav",court)
	go player("sanap",court)

	court <- 1

	wg.Wait()
}



func player(name string, court chan int){

	defer wg.Done()

	for {

		ball ,ok := <- court

		if !ok {

			fmt.Printf("Player %s won\n", name)
		}


		n :=rand.Intn(200)
		if(n % 13 == 0){
			fmt.Printf("ball dropped from player %s\n", name)
			close(court)
			return
		}

		fmt.Printf("Player %s hit the ball %s",name,ball )

		ball++

		court <- ball

	}
	

}