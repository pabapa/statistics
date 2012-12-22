package statistics

// Copyright 2012 G.vd.Schoot. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

func Kurtosis(data Interface) float64 {
	mean := Mean(data)
	est_sd := SdMean(data, mean)
	return KurtosisMainSd(data, mean, est_sd)
}

func KurtosisMainSd(data Interface, mean, sd float64) float64 {
	var avg, kurtosis float64

	for i := 0; i < data.Len(); i++ {
		x := (data.Value(i) - mean) / sd
		avg += (x*x*x*x - avg) / float64(i+1)
	}

	kurtosis = avg - 3.0 // makes kurtosis zero for a Gaussian

	return kurtosis
}
