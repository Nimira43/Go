package main

import "math"

func main() {
	var investmentAmount = 10000
	var expectedReturnRate = 6.5
	var years = 10

	var futureValue = float64(investmentAmount) * math.Pow(1+expectedReturnRate/100, float64(years))

}
