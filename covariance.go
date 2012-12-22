package statistics

// Copyright 2012 G.vd.Schoot. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

func covariance(data1, data2 Interface, mean1, mean2 float64) (ret float64) {
	for i := 0; i < data1.Len(); i++ {
		delta1 := (data1.Value(i) - mean1)
		delta2 := (data2.Value(i) - mean2)
		ret += (delta1*delta2 - ret) / float64(i+1)
	}
	return
}

func CovarianceMean(data1, data2 Interface, mean1, mean2 float64) float64 {
	n := data1.Len()
	covariance := covariance(data1, data2, mean1, mean2)
	return covariance * float64(n) / float64(n-1)
}

func Covariance(data1, data2 Interface) float64 {
	mean1 := Mean(data1)
	mean2 := Mean(data2)

	return CovarianceMean(data1, data2, mean1, mean2)
}
