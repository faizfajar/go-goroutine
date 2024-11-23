package latihangolanggoroutine

import (
	"fmt"
	"sync"
	"testing"
)

var counter = 0

func OnlyOnce() {
	counter++
}

func TestOnce(t *testing.T) {
	once := sync.Once{}
	group := sync.WaitGroup{}
	// mutex := &sync.Mutex{}

	for i := 0; i < 100; i++ {
		group.Add(1)
		go func() {

			defer group.Done()
			// mutex.Lock()
			once.Do(OnlyOnce)
			// OnlyOnce()
			// mutex.Unlock()

		}()
	}

	group.Wait()
	fmt.Println("Counter", counter)
}
