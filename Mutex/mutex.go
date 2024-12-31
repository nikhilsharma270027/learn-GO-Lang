package main // race condition
import (
	"fmt"
	"sync"
)

// example social media
type post struct {
	views int
	mu    sync.Mutex
}

func (p *post) inc(wg *sync.WaitGroup) {
	// defer wg.Done()
	defer func() {
		p.mu.Unlock()
		wg.Done()
	}()
	p.mu.Lock() // always been where the data is been modified
	p.views += 1
	// p.mu.Unlock()
}

func main() {
	var wg sync.WaitGroup
	mypost := post{views: 0}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go mypost.inc(&wg)
		// suppose if i use go routine
	}
	wg.Wait()

	fmt.Println(mypost.views)

}
