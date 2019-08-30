package main

import "sync"

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			var value int
			for value != 1 {
				value = (value + 1) * 222 / 22
			}
		}()
	}
	wg.Wait()
}
