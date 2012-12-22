package statistics

// Copyright 2012 G.vd.Schoot. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

func QuantileFromSortedData(sorted_data Interface, f float64) (result float64) {
	n := sorted_data.Len()
	index := f * float64(n-1)
	lhs := int(index)
	delta := index - float64(lhs)

	if n == 0 {
		return 0.0
	}

	if lhs == n-1 {
		result = sorted_data.Value(lhs)
	} else {
		result = (1-delta)*sorted_data.Value(lhs) + delta*sorted_data.Value(lhs+1)
	}

	return
}
