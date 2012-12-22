package statistics

// Copyright 2012 G.vd.Schoot. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

import (
	"math"
)

func WAbsdev(w, data Interface) float64 {
	wmean := WMean(w, data)
	return WAbsdevMean(w, data, wmean)
}

// WAbsdevMean calculates the weighted absolute deviation of a dataset
func WAbsdevMean(w, data Interface, wmean float64) (wabsdev float64) {

	var W float64
	size := data.Len()

	// calculate the sum of the absolute deviations
	for i := 0; i < size; i++ {
		wi := w.Value(i)

		if wi > 0 {
			delta := math.Abs(data.Value(i) - wmean)
			W += wi
			wabsdev += (delta - wabsdev) * (wi / W)
		}
	}

	return
}
