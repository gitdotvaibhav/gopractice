package main

import(
	"sync"
	"time"
	"fmt"
)


var(

	wg sync.WaitGroup;

)

func main(){

	runner := make(chan int,1)

	wg.Add(1)


	go raley(runner)

	runner <- 1

	wg.Wait()


}



func raley(runner chan int){

	player := <- runner
	if(player == 4){
		wg.Done()
		close(runner)
		return	
	}

	time.Sleep(100*time.Millisecond)

	player = player+1;
	runner <- player
	go raley(runner)


	fmt.Printf("Exited %d",player)
	return



}