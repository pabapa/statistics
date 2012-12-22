package statistics

// Copyright 2012 G.vd.Schoot. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

import (
	"math"
)

func wvariance(w, data Interface, wmean float64) (wvariance float64) {
	var W float64
	size := data.Len()

	// the sum of the squares
	for i := 0; i < size; i++ {
		wi := w.Value(i)

		if wi > 0 {
			delta := data.Value(i) - wmean
			W += wi
			wvariance += (delta*delta - wvariance) * (wi / W)
		}
	}

	return
}

func factor(w Interface) (factor float64) {

	var a, b float64
	size := w.Len()

	// the sum of the squares
	for i := 0; i < size; i++ {
		wi := w.Value(i)

		if wi > 0 {
			a += wi
			b += wi * wi
		}
	}

	factor = (a * a) / ((a * a) - b)

	return
}

func WVarianceWithFixedMean(w, data Interface, wmean float64) float64 {
	return wvariance(w, data, wmean)
}

func WsdWithFixedMean(w, data Interface, wmean float64) float64 {
	wvariance := wvariance(w, data, wmean)
	return math.Sqrt(wvariance)
}

func WVarianceMean(w, data Interface, wmean float64) float64 {
	variance := wvariance(w, data, wmean)
	scale := factor(w)

	return scale * variance
}

func WSdMean(w, data Interface, wmean float64) float64 {
	variance := wvariance(w, data, wmean)
	scale := factor(w)
	return math.Sqrt(scale * variance)
}

func WSd(w, data Interface) float64 {
	wmean := WMean(w, data)
	return WSdMean(w, data, wmean)
}

func WVariance(w, data Interface) float64 {
	wmean := WMean(w, data)
	return WVarianceMean(w, data, wmean)
}
