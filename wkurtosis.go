package statistics

// Copyright 2012 G.vd.Schoot. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

func WKurtosis(w, data Interface) float64 {
	wmean := WMean(w, data)
	wsd := WSdMean(w, data, wmean)
	return WKurtosisMeanSd(w, data, wmean, wsd)
}

// WKurtosisMean calculates the kurtosis of a dataset
func WKurtosisMeanSd(w, data Interface, wmean, wsd float64) float64 {

	var wavg, W float64
	Len := data.Len()

	for i := 0; i < Len; i++ {
		wi := w.Value(i)

		if wi > 0 {
			x := (data.Value(i) - wmean) / wsd
			W += wi
			wavg += (x*x*x*x - wavg) * (wi / W)
		}
	}

	return wavg - 3.0 // makes kurtosis zero for a Gaussian
}
