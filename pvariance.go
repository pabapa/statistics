package statistics

// Copyright 2012 G.vd.Schoot. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// PVariance finds the pooled variance of two datasets 
func PVariance(data1, data2 Interface) float64 {

	n1 := data1.Len()
	n2 := data2.Len()

	var1 := Variance(data1)
	var2 := Variance(data2)

	return ((float64(n1-1) * var1) + (float64(n2-1) * var2)) / float64(n1+n2-2)
}
