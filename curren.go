package main

import "fmt"

func calCoins(currency int, coins []int) int {
	coin := 0
	remaining := currency
	for i := len(coins) - 1; i >= 0; i-- {
		coins := coins[i]
		count := remaining / coins
		if count > 0 {
			coin += count
			remaining -= count * coins
		}
	}
	return coin
}

func main() {

	coins := []int{1, 5, 10, 50, 100, 500}

	var mn int
	fmt.Println("Enter the currenies : ")
	fmt.Scanln(&mn)

	currencies := make([]int, mn)
	for i := 0; i < mn; i++ {
		fmt.Printf("Enter the currency %d : ", i+1)
		fmt.Scan(&currencies[i])

		// currencies := []int{0, 12, 468, 123456}

		for _, currency := range currencies {
			coin := calCoins(currency, coins)
			fmt.Printf("For %d currency, %d coins are needed.\n", currency, coin)
		}
	}
}
