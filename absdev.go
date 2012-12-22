package statistics

// Copyright 2012 G.vd.Schoot. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

import (
	"math"
)

func Absdev(data Interface) float64 {
	mean := Mean(data)
	return AbsdevMean(data, mean)
}

// AbsdevMean finds the absolute deviation of the data interface
func AbsdevMean(data Interface, mean float64) float64 {

	var sum float64

	// the sum of the absolute deviations 
	for i := 0; i < data.Len(); i++ {
		sum += math.Abs(data.Value(i) - mean)
	}

	return sum / float64(data.Len())
}
