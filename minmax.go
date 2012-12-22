package statistics

// Copyright 2012 G.vd.Schoot. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

import (
	"math"
)

// Max finds the first largest member and the members position within the data 
func Max(data Interface) (max float64, max_index int) {
	Len := data.Len()
	max = data.Value(0)

	for i := 0; i < Len; i++ {
		xi := data.Value(i)

		if xi > max {
			max = xi
			max_index = i
		}

		if math.IsNaN(xi) {
			max = xi
			max_index = i
			return
		}
	}

	return
}

// Min finds the first smallest member and the members position within the data 
func Min(data Interface) (min float64, min_index int) {
	Len := data.Len()
	min = data.Value(0)

	for i := 0; i < Len; i++ {
		xi := data.Value(i)

		if xi < min {
			min = xi
			min_index = i
		}

		if math.IsNaN(xi) {
			min = xi
			min_index = i
			return
		}
	}

	return
}

// Minmax finds the first smallest and largest members and 
// the members positions within the data 
func Minmax(data Interface) (min float64, min_index int, max float64, max_index int) {
	Len := data.Len()
	min = data.Value(0)
	max = data.Value(0)

	for i := 0; i < Len; i++ {
		xi := data.Value(i)

		if xi < min {
			min = xi
			min_index = i
		}

		if xi > max {
			max = xi
			max_index = i
		}

		if math.IsNaN(xi) {
			min = xi
			max = xi
			min_index = i
			max_index = i
			break
		}
	}

	return
}
