package main

import "fmt"

type Purchase struct{
	array []int64
	cashback int64
	sumcashback int64
	sum int64
}

	func(percent Purchase) countCashback(){
		percent.sum = int64(0)
		for _, value := range percent.array{
			percent.sum += value
		}
	percent.sumcashback = percent.cashback * percent.sum / 100
	 fmt.Println(percent.sumcashback)
}

func main() {
	purchases := Purchase{
		array: []int64{100, 50, 70},
		cashback: 15,
		sumcashback: 0,
		sum: 0,
	}
purchases.countCashback()
}