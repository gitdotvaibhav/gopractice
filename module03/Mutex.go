package main

import(
	"sync"
	"fmt"
	"runtime"
)


var(
	counter int
	wg sync.WaitGroup
	mutex sync.Mutex

)


func main(){

	runtime.GOMAXPROCS(2)
	
	wg.Add(2)

	go increamentCount()

	go increamentCount()

	fmt.Printf("%d\n", counter)

	wg.Wait()

	fmt.Printf("%d\n", counter)

}



func increamentCount(){
	defer wg.Done()

	fmt.Printf("%d\n", counter)


	for i:=0; i < 2;i++ {

		mutex.Lock()
		{
			value :=counter

			runtime.Gosched()

			value = value+1
			counter =value

		}

		mutex.Unlock()

	}

}