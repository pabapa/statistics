package statistics

// Copyright 2012 G.vd.Schoot. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

func Skew(data Interface) float64 {
	mean := Mean(data)
	sd := SdMean(data, mean)
	return SkewMeanSd(data, mean, sd)
}

// SkewMeanSd calculates the skewness of a dataset
func SkewMeanSd(data Interface, mean, sd float64) (skew float64) {
	for i := 0; i < data.Len(); i++ {
		x := (data.Value(i) - mean) / sd
		skew += (x*x*x - skew) / float64(i+1)
	}
	return
}
