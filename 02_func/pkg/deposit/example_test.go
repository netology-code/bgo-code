package deposit_test

import (
	"fmt"
	"lectionfunctions/pkg/deposit"
)

func ExampleCalc() {
	fmt.Println(deposit.Calculate(1_000))
	fmt.Println(deposit.Calculate(100))
	// Output:
	// 1040 1060
	// 104 106
}


