package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {

	powCh := make(chan []int, 1)
	resCh := make(chan []int, 1)
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {

		defer wg.Done()
		var transferSlice []int
		count := 10
		for i := 0; i < count; i++ {
			transferSlice = append(transferSlice, rand.Intn(101))
		}

		powCh <- transferSlice
	}()

	go func() {
		defer wg.Done()
		var resultSlice []int

		for _, v := range <-powCh {
			resultSlice = append(resultSlice, v*v)
		}

		resCh <- resultSlice
	}()

	wg.Wait()

	for _, num := range <-resCh {
		fmt.Printf("%d ", num)
	}
}
