// package main

// import (
// 	"fmt"
// 	"time"
// )

// // func task(id int) {
// // 	fmt.Println("doing task", id)

// // }

// func main() {
// 	for i:= 0; i <= 10; i++ {
// 		// go task(i)   // runs concurrently/ parallely

// 		go func(i int){
// 			fmt.Println(i)

// 		}(i)

// 	}
// 	time.Sleep(time.Second + 2)
// }

package main

import (
	"fmt"
	"sync"
	// "time"
)

func task(id int, w *sync.WaitGroup) {
	defer w.Done()
	fmt.Println("doing task", id)

}

func main() {
	var wg sync.WaitGroup
	for i:= 0; i <= 10; i++ {
		wg.Add(1) // counter type mechanise
		go task(i, &wg)   // runs concurrently/ parallely
	}

	wg.Wait()
}