package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	var investmentAmount = 1000
	var expectedReturn = 5.5
	var years = 10
	fmt.Print("------\n")
	fmt.Scan(&investmentAmount)
	var futureValue = float64(investmentAmount) * math.Pow(1+expectedReturn/100, float64(years))

	futureRealValue := futureValue / math.Pow(1+inflationRate/100, float64(years))
	fmt.Println(futureValue)
	fmt.Println(futureRealValue)

}
