package statistics

// Copyright 2012 G.vd.Schoot. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

import (
	"testing"
)

var data = Float64{1.0}
var data1 = Float64{0.0, 1.0, 2.0, 3.0, 4.0, 5.0}
var data2 = make(Float64, 10000)

var data_I = Int64{1}
var data1_I = Int64{0, 1, 2, 3, 4, 5}
var data2_I = make(Int64, 10000)

func BenchmarkMean(b *testing.B) {
	var f float64
	for i := 0; i < b.N; i++ {
		f = Mean(&data)
	}
	_ = f // make compiler happy
}

func BenchmarkMean1(b *testing.B) {
	var f float64
	for i := 0; i < b.N; i++ {
		f = Mean(&data1)
	}
	_ = f // make compiler happy
}

func BenchmarkMean2(b *testing.B) {
	var f float64
	for i := 0; i < b.N; i++ {
		f = Mean(&data2)
	}
	_ = f // make compiler happy
}

func BenchmarkMeanI(b *testing.B) {
	var f float64
	for i := 0; i < b.N; i++ {
		f = Mean(&data_I)
	}
	_ = f // make compiler happy
}

func BenchmarkMean1_I(b *testing.B) {
	var f float64
	for i := 0; i < b.N; i++ {
		f = Mean(&data1_I)
	}
	_ = f // make compiler happy
}

func BenchmarkMean2_I(b *testing.B) {
	var f float64
	for i := 0; i < b.N; i++ {
		f = Mean(&data2_I)
	}
	_ = f // make compiler happy
}
