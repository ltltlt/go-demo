// GOMAXPROCS=1 go run thisfile
package main

import "sync"
import "runtime"
import "fmt"

func test() {
	for i := 0; i < 5; i++ {
		runtime.Gosched()
		fmt.Println(i)
	}
}
func main() {
	var wg sync.WaitGroup
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			test()
		}()
	}
	wg.Wait()
}
