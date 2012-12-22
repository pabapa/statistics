package statistics

// Copyright 2012 G.vd.Schoot. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

import (
	"testing"
)

// Main test routine
func TestMean(t *testing.T) {

	for s1 := 1; s1 < 4; s1++ {
		s2 := 1
		if s1 >= 3 {
			s2 = s1 - 1
		}

		testFloat64(t, s1, s2)
		testInt64(t, s1, s2)
	}

	testNist()
}
