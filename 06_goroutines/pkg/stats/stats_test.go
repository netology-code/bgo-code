package stats

import "testing"

func makeTransactions() []int64 {
	const users = 100_000
	const transactionsPerUser = 100
	const transactionAmount = 1_00
	transactions := make([]int64, users * transactionsPerUser)
	for index := range transactions {
		transactions[index] = transactionAmount
	}
	return transactions
}

func BenchmarkSum(b *testing.B) {
	transactions := makeTransactions()
	want := int64(10_000_000_00)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result := Sum(transactions)
		if result != want {
			b.Errorf("invalid result, got %v, want %v", result, want)
		}
	}
}

func BenchmarkSumConcurrently4(b *testing.B) {
	benchmarkSumConcurrently(b, 4)
}
func BenchmarkSumConcurrently500(b *testing.B) {
	benchmarkSumConcurrently(b, 500)
}

func benchmarkSumConcurrently(b *testing.B, goroutines int) {
	transactions := makeTransactions()
	want := int64(10_000_000_00)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result := SumConcurrently(transactions, goroutines)
		if result != want {
			b.Fatalf("invalid result, got %v, want %v", result, want)
		}
	}
}

