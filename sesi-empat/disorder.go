package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

func disorder() {
	var wg sync.WaitGroup
	var bisa interface{}
	var coba interface{}

	bisa = []string{"bisa1", "bisa2", "bisa3"}
	coba = []string{"coba1", "coba2", "coba3"}

	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 4; i++ {
			time.Sleep(time.Duration(rand.IntN(100)) * time.Millisecond)
			fmt.Println(bisa, i+1)
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < 4; i++ {
			time.Sleep(time.Duration(rand.IntN(100)) * time.Millisecond)
			fmt.Println(coba, i+1)
		}
	}()
	wg.Wait()
}
