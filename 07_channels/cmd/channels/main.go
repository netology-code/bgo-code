package main

import (
	"fmt"
	"lectionchannels/pkg/stats"
)

func sum(transactions []int64, result chan <- int64) {
	go func() {
		result <- stats.Sum(transactions)
	}()
}

func main() {
	const users = 10_000_000
	const transactionsPerUser = 100
	const transactionAmount = 1_00
	transactions := make([]int64, users * transactionsPerUser)
	for index := range transactions {
		// для простоты храним только суммы
		// и считаем, что каждая транзакция = 1 руб.
		transactions[index] = transactionAmount
	}

	total := int64(0)
	const partsCount = 10
	partSize := len(transactions) / partsCount
	result := make(chan int64)
	for i := 0; i < partsCount; i++ {
		part := transactions[i * partSize : (i + 1) * partSize]
		sum(part, result)
	}

	finished := 0
	for value := range result {
		total += value
		finished++
		progress := finished * 100 / partsCount
		fmt.Printf("Done: %d%%\n", progress)
		if finished == partsCount {
			break
		}
	}

	fmt.Println(total)
}
