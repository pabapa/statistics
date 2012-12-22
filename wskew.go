package statistics

// Copyright 2012 G.vd.Schoot. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

func WSkew(w, data Interface) float64 {
	wmean := WMean(w, data)
	wsd := WSdMean(w, data, wmean)
	return WSkewMeanSd(w, data, wmean, wsd)
}

// Compute the weighted skewness of a dataset
func WSkewMeanSd(w, data Interface, wmean, wsd float64) (wskew float64) {
	var W float64
	size := data.Len()

	for i := 0; i < size; i++ {
		wi := w.Value(i)

		if wi > 0 {
			x := (data.Value(i) - wmean) / wsd
			W += wi
			wskew += (x*x*x - wskew) * (wi / W)
		}
	}

	return
}
