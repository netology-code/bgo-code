package deposit

func Calculate(amount int64) (int64, int64) {
	const minPercent = 4
	minResult := (amount * (100 + minPercent)) / 100
	const maxPercent = 6
	maxResult := (amount * (100 + maxPercent)) / 100
	return minResult, maxResult
}
