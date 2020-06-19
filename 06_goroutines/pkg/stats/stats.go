package stats

import "sync"

func Sum(transactions []int64) int64 {
	result := int64(0)
	for _, transaction := range transactions {
		result += transaction
	}
	return result
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
