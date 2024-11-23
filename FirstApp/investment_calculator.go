package main

import (
	"fmt"
	"math"
)

const inflationRate = 4.5

func main() {

	var investmentAmount float64 
	var years float64 
	var expectedReturnRate float64 
	

	outputText("Investment Amount: ")
	fmt.Scan(&investmentAmount)
	outputText("Years: ")
	fmt.Scan(&years)
	outputText("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	futureValue, futureRealValue :=  calculateFutureValues(investmentAmount, expectedReturnRate, years);

	formattedFV := fmt.Sprintf("future value: %.2f\n" , futureValue)
	formattedRFV := fmt.Sprintf("future value (adjusted for inflation): %.2f\n" , futureRealValue)

	fmt.Print(formattedFV, formattedRFV);
}

func outputText(text string) {
	fmt.Print(text)
}

func calculateFutureValues(investmentAmount float64, expectedReturnRate float64, years float64 ) (float64, float64) {

	fv := investmentAmount * math.Pow(1 + expectedReturnRate / 100, years) 
	rfv := fv / math.Pow(1 + inflationRate / 100, years)

	return fv, rfv
}