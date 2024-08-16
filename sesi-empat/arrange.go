package main

import (
	"fmt"
	"sync"
	"time"
)

func arrange() {
	var wg sync.WaitGroup
	var bisa interface{}
	var coba interface{}

	bisa = []string{"bisa1", "bisa2", "bisa3"}
	coba = []string{"coba1", "coba2", "coba3"}

	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 4; i++ {
			fmt.Println(bisa, i+1)
			time.Sleep(time.Millisecond * 5)
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < 4; i++ {
			fmt.Println(coba, i+1)
			time.Sleep(time.Millisecond * 5)
		}
	}()
	wg.Wait()
}
