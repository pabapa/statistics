package statistics

// Copyright 2012 G.vd.Schoot. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// WMean calculates the weighted arithmetic mean of a dataset
func WMean(w, data Interface) (wmean float64) {
	var W float64
	size := data.Len()

	for i := 0; i < size; i++ {
		wi := w.Value(i)

		if wi > 0 {
			W += wi
			wmean += (data.Value(i) - wmean) * (wi / W)
		}
	}

	return
}
