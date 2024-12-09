package main

import (
	"fmt"

	filemanager "example.com/price-calculator/fileManager"
	"example.com/price-calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	result := make(map[float64][]float64)

	for _, taxRate := range taxRates {
		fm := filemanager.New("prices.txt", fmt.Sprintf("result%.0f.json", taxRate*100))
		// cmdm := cmdmanager.New() // can swap out these input output packages easily because of ioManager interface (interface Segregation)
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		err := priceJob.Process()

		if err != nil {
			fmt.Println("Could not process job")
			fmt.Println(err)
		}
	}

	fmt.Println(result)
}
