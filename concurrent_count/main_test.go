// filepath: d:\GolangProject\demo\concurrent_count\main_test.go
package main

import (
	"sync"
	"testing"
)

func TestAdd(t *testing.T) {
	count := Count{value: 0, mu: sync.Mutex{}}
	count.Add()
	if count.GetValue() != 1 {
		t.Errorf("Expected value 1, got %d", count.GetValue())
	}
}

func TestGetValue(t *testing.T) {
	count := Count{value: 42, mu: sync.Mutex{}}
	if count.GetValue() != 42 {
		t.Errorf("Expected value 42, got %d", count.GetValue())
	}
}

func TestConcurrentAccess(t *testing.T) {
	count := Count{value: 0, mu: sync.Mutex{}}
	var wg sync.WaitGroup

	// Increment the counter 1000 times concurrently
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			count.Add()
		}()
	}

	// Wait for all goroutines to finish
	wg.Wait()

	if count.GetValue() != 1000 {
		t.Errorf("Expected value 1000, got %d", count.GetValue())
	}
}
