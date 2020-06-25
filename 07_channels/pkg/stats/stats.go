package stats

import (
	"log"
	"sync"
	"sync/atomic"
)

func Sum(transactions []int64) int64 {
	result := int64(0)
	for _, transaction := range transactions {
		result += transaction
	}
	return result
}

func SumAtomic(transactions []int64, goroutines int) int64 {
	wg := sync.WaitGroup{}
	wg.Add(goroutines)

	total := int64(0)
	partSize := len(transactions) / goroutines
	for i := 0; i < goroutines; i++ {
		part := transactions[i*partSize : (i+1)*partSize]
		go func() {
			log.Println(i)
			sum := Sum(part)
			atomic.AddInt64(&total, sum) // no bug: atomic change
			wg.Done()
		}()
	}

	wg.Wait()
	return atomic.LoadInt64(&total)
}

func SumMutex(transactions []int64, goroutines int) int64 {
	wg := sync.WaitGroup{}
	wg.Add(goroutines)

	mu := sync.Mutex{}
	total := int64(0)

	partSize := len(transactions) / goroutines
	for i := 0; i < goroutines; i++ {
		part := transactions[i*partSize : (i+1)*partSize]
		go func() {
			sum := Sum(part)
			mu.Lock()
			total += sum // no bug: change under mutex
			mu.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()
	mu.Lock()
	result := total
	mu.Unlock()
	return result
}

func SumMutexFIXME(transactions []int64, goroutines int) int64 {
	wg := sync.WaitGroup{}
	wg.Add(goroutines)

	mu := sync.Mutex{}
	total := int64(0)

	partSize := len(transactions) / goroutines
	for i := 0; i < goroutines; i++ {
		part := transactions[i*partSize : (i+1)*partSize]
		go func() {
			mu.Lock() // FIXME: ни в коем случае!
			sum := Sum(part)
			total += sum
			wg.Done()
			mu.Unlock()
		}()
	}

	wg.Wait()
	return total
}

func SumConcurrently(transactions []int64, goroutines int) int64 {
	wg := sync.WaitGroup{}
	wg.Add(goroutines)

	total := int64(0)
	partSize := len(transactions) / goroutines
	for i := 0; i < goroutines; i++ {
		part := transactions[i*partSize : (i+1)*partSize]
		go func() {
			total += Sum(part) // FIXME: shared memory bug, discuss later
			wg.Done()
		}()
	}

	wg.Wait()
	return total
}
