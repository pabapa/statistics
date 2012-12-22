package statistics

// Copyright 2012 G.vd.Schoot. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

import (
	"math"
)

// runs a t-test between two datasets representing independent
// samples. Tests to see if the difference between means of the
// samples is different from zero.
func TTest(data1, data2 Interface) float64 {

	n1 := float64(data1.Len())
	n2 := float64(data2.Len())

	mean1 := Mean(data1)
	mean2 := Mean(data2)

	pv := PVariance(data1, data2)

	t := (mean1 - mean2) / (math.Sqrt(pv * ((1.0 / n1) + (1.0 / n2))))

	return t
}
