package statistics

// Copyright 2012 G.vd.Schoot. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

func MedianFromSortedData(sorted_data Interface) (median float64) {
	Len := sorted_data.Len()
	lhs := (Len - 1) / 2
	rhs := Len / 2

	if Len == 0 {
		return 0.0
	}

	if lhs == rhs {
		median = sorted_data.Value(lhs)
	} else {
		median = (sorted_data.Value(lhs) + sorted_data.Value(rhs)) / 2.0
	}

	return
}
