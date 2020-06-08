package main

import (
	"fmt"
	"lectionfunctions/pkg/deposit"
)

func main() {
	amount := int64(100_000_00)
	min, max := deposit.Calculate(amount)
	fmt.Println(min, max)
}
