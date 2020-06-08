package main

func main() {
	cpuPrice := 146.88
	ramPrice := 71.28
	hddPrice := 6.25
	count := 1.0

	total := (cpuPrice + ramPrice + hddPrice) * count
	println(total)
}

// в вашем коде не должно быть закомментированных функций
// мы их оставляем лишь для того, чтобы не плодить "кучу проектов"
// и не заставлять вас перепечатывать с лекций
//func main() {
//	balance := 1.6
//	balance += 0.3
//	total := 1.9
//	equal := balance == total
//	println(equal)
//}

