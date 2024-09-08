package main

import "fmt"

func main() {

	priceSlice := []float64{10, 20, 30}
	taxRateSlice := []float64{0, 0.7, 1, 30, 50} // tax rates are in percentage
	priceWithTax := make(map[float64][]float64)

	for _, rate := range taxRateSlice {
		temp := make([]float64, len(priceSlice))
		for j, price := range priceSlice {
			temp[j] = (rate/100)*price + price
		}
		priceWithTax[rate] = temp
	}

	fmt.Println(priceWithTax)
}
