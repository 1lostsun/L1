package main

import (
	"fmt"
	"math"
)

func main() {
	arr := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	temperatureDiapason(arr)
}

func temperatureDiapason(arr []float64) {
	mp := make(map[float64][]float64)
	for _, v := range arr {
		subZeroKey := math.Ceil(v/10) * 10
		aboveZeroKey := (math.Ceil(v/10) - 1) * 10

		if math.Signbit(v) {
			mp[subZeroKey] = append(mp[subZeroKey], v)
		} else {
			if aboveZeroKey+10 == v {
				mp[aboveZeroKey+10] = append(mp[aboveZeroKey+10], v)
				continue
			}
			mp[aboveZeroKey] = append(mp[aboveZeroKey], v)
		}
	}

	fmt.Println(mp)
}
