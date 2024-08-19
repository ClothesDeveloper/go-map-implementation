package main

import (
	"sync"
	"testing"
)

func TestConcurrentWrite(t *testing.T) {
	myMap := NewMap(10)
	var wg sync.WaitGroup
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func() {
			defer wg.Done()
			myMap.Set("dude", 10)
		}()
	}
	wg.Wait()
}

func TestConcurrentRWD(t *testing.T) {
	myMap := NewMap(10)
	someKey := "someKey"

	var wg sync.WaitGroup
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func(i int) {
			defer wg.Done()
			if i%2 == 0 {
				myMap.Get(someKey)
				return
			}

			if i%3 == 0 {
				myMap.Delete(someKey)
				return
			}

			myMap.Set(someKey, i)
		}(i)
	}
}
