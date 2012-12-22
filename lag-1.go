package statistics

// Copyright 2012 G.vd.Schoot. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

func Lag1Autocorrelation(data Interface) float64 {
	mean := Mean(data)
	return Lag1AutocorrelationMean(data, mean)
}

func Lag1AutocorrelationMean(data Interface, mean float64) float64 {
	var r1, q float64
	Len := data.Len()
	v := (data.Value(0) - mean) * (data.Value(0) - mean)

	for i := 1; i < Len; i++ {
		delta0 := data.Value(i-1) - mean
		delta1 := data.Value(i) - mean
		q += (delta0*delta1 - q) / float64(i+1)
		v += (delta1*delta1 - v) / float64(i+1)
	}

	r1 = q / v

	return r1
}
