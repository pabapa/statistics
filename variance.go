package statistics

// Copyright 2012 G.vd.Schoot. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

import (
	"math"
)

func _variance(data Interface, mean float64) (variance float64) {

	// calculate the sum of the squares
	for i := 0; i < data.Len(); i++ {
		delta := data.Value(i) - mean
		// TODO: long double for variance... How to implement in Go?
		variance += (delta*delta - variance) / float64(i+1)
	}

	return
}

func VarianceWithFixedMean(data Interface, mean float64) float64 {
	return _variance(data, mean)
}

func SdWithFixedMean(data Interface, mean float64) float64 {
	variance := _variance(data, mean)
	return math.Sqrt(variance)
}

func VarianceMean(data Interface, mean float64) float64 {
	variance := _variance(data, mean)
	return variance * float64(data.Len()) / float64(data.Len()-1)
}

func SdMean(data Interface, mean float64) float64 {
	variance := _variance(data, mean)
	return math.Sqrt(variance * float64(data.Len()) / float64(data.Len()-1))
}

func Variance(data Interface) float64 {
	mean := Mean(data)
	return VarianceMean(data, mean)
}

func Sd(data Interface) float64 {
	mean := Mean(data)
	return SdMean(data, mean)
}
