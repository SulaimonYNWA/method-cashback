package main

import "fmt"

type Purchase struct{
	array []int64
	cashback int64
	sumcashback int64
}

func(percent Purchase) countCashback(cashback int64, sum int64){
	 percent.sumcashback = cashback*sum/100
	 fmt.Println(percent.sumcashback)
}

func main() {
	
	purchases := Purchase{
		// array: []int64{100, 50, 70},
		// cashback: 15,
	}
	purchases.array = []int64{100, 50, 70}
	purchases.cashback = 15
	purchases.sumcashback = 0

	var sum int64 = 0
	for _, value := range purchases.array{
		sum += value
	}

fmt.Println(sum)
finalSum:=sum*(100-purchases.cashback)/100
fmt.Println(finalSum)

purchases.countCashback(purchases.cashback, sum)

}