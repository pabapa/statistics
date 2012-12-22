package statistics

// Copyright 2012 G.vd.Schoot. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Mean calculates the arithmetic mean with the recurrence relation
func Mean(data Interface) (mean float64) {
	Len := data.Len()
	for i := 0; i < Len; i++ {
		mean += (data.Value(i) - mean) / float64(i+1)
	}
	return
}
