package statistics

/* statistics/float64_test.go
 * 
 * Copyright (C) 1996, 1997, 1998, 1999, 2000 Jim Davies, Brian Gough
 * 
 * This program is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 2 of the License, or (at
 * your option) any later version.
 * 
 * This program is distributed in the hope that it will be useful, but
 * WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
 * General Public License for more details.
 * 
 * You should have received a copy of the GNU General Public License
 * along with this program; if not, write to the Free Software
 * Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston, MA 02110-1301, USA.
 */

import (
	"math"
	"sort"
	"testing"
)

// Main test routine
func testFloat64(t *testing.T, stridea, strideb int) {

	/* sample sets of doubles */
	var i int
	na := 14
	nb := 14

	rawa := Float64{
		.0421, .0941, .1064, .0242, .1331,
		.0773, .0243, .0815, .1186, .0356,
		.0728, .0999, .0614, .0479}

	rawb := Float64{
		.1081, .0986, .1566, .1961, .1125,
		.1942, .1079, .1021, .1583, .1673,
		.1675, .1856, .1688, .1512}

	raww := Float64{
		.0000, .0000, .0000, 3.000, .0000,
		1.000, 1.000, 1.000, 0.000, .5000,
		7.000, 5.000, 4.000, 0.123}

	groupa1 := make(Float64, stridea*na)
	groupa := NewStrider(&groupa1, stridea)
	groupb1 := make(Float64, strideb*nb)
	groupb := NewStrider(&groupb1, strideb)
	w1 := make(Float64, strideb*na)
	w := NewStrider(&w1, strideb)

	rel := 1e-10

	for i = 0; i < na; i++ {
		groupa1[i*stridea] = rawa[i]
	}

	for i = 0; i < na; i++ {
		w1[i*strideb] = raww[i]
	}

	for i = 0; i < nb; i++ {
		groupb1[i*strideb] = rawb[i]
	}

	{
		mean := Mean(groupa)
		expected := 0.0728
		gsl_test_rel(mean, expected, rel, "Mean()")
	}

	{
		mean := Mean(groupa)
		varc := VarianceWithFixedMean(groupa, mean)
		expected := 0.00113837428571429
		gsl_test_rel(varc, expected, rel, "VarianceWithFixedMean")
	}

	{
		mean := Mean(groupa)
		varc := SdWithFixedMean(groupa, mean)
		expected := 0.0337398026922845
		gsl_test_rel(varc, expected, rel, "SdWithFixedMean")
	}

	{
		varc := Variance(groupb)
		expected := 0.00124956615384615
		gsl_test_rel(varc, expected, rel, "Variance")
	}

	{
		sd := Sd(groupa)
		expected := 0.0350134479659107
		gsl_test_rel(sd, expected, rel, "Sd")
	}

	{
		absdev := Absdev(groupa)
		expected := 0.0287571428571429
		gsl_test_rel(absdev, expected, rel, "Absdev")
	}

	{
		skew := Skew(groupa)
		expected := 0.0954642051479004
		gsl_test_rel(skew, expected, rel, "Skew")
	}

	{
		kurt := Kurtosis(groupa)
		expected := -1.38583851548909
		gsl_test_rel(kurt, expected, rel, "Kurtosis")
	}

	{
		wmean := WMean(w, groupa)
		expected := 0.0678111523670601
		gsl_test_rel(wmean, expected, rel, "WMean")
	}

	{
		wmean := WMean(w, groupa)
		wvar := WVarianceWithFixedMean(w, groupa, wmean)
		expected := 0.000615793060878654
		gsl_test_rel(wvar, expected, rel, "WVarianceWithFixedMean")
	}

	{
		est_wvar := WVariance(w, groupa)
		expected := 0.000769562962860317
		gsl_test_rel(est_wvar, expected, rel, "WVariance")
	}

	{
		wsd := WSd(w, groupa)
		expected := 0.0277409978706664
		gsl_test_rel(wsd, expected, rel, "WSd")
	}

	{
		wabsdev := WAbsdev(w, groupa)
		expected := 0.0193205027504008
		gsl_test_rel(wabsdev, expected, rel, "WAbsdev")
	}

	{
		wskew := WSkew(w, groupa)
		expected := -0.373631000307076
		gsl_test_rel(wskew, expected, rel, "WSkew")
	}

	{
		wkurt := WKurtosis(w, groupa)
		expected := -1.48114233353963
		gsl_test_rel(wkurt, expected, rel, "WKurtosis")
	}

	{
		c := Covariance(groupa, groupb)
		expected := -0.000139021538461539
		gsl_test_rel(c, expected, rel, "Covariance")
	}

	{
		pv := PVariance(groupa, groupb)
		expected := 0.00123775384615385
		gsl_test_rel(pv, expected, rel, "PVariance")
	}

	{
		t := TTest(groupa, groupb)
		expected := -5.67026326985851
		gsl_test_rel(t, expected, rel, "TTest")
	}

	{
		max, _ := Max(groupa)
		expected := 0.1331
		gsl_test(max != expected,
			"Max (%g observed vs %g expected)",
			max, expected)
	}

	{
		min, _ := Min(groupa)
		expected := 0.0242
		gsl_test(min != expected,
			"Min (%g observed vs %g expected)",
			min, expected)
	}

	{
		min, _, max, _ := Minmax(groupa)
		expected_max := 0.1331
		expected_min := 0.0242

		gsl_test(max != expected_max,
			"Minmax max (%g observed vs %g expected)",
			max, expected_max)
		gsl_test(min != expected_min,
			"Minmax min (%g observed vs %g expected)",
			min, expected_min)
	}

	{
		_, max_index := Max(groupa)
		expected := 4
		gsl_test(max_index != expected,
			"Max (%d observed vs %d expected)",
			max_index, expected)
	}

	{
		_, min_index := Min(groupa)
		expected := 3
		gsl_test(min_index != expected,
			"Min (%d observed vs %d expected)",
			min_index, expected)
	}

	{
		_, min_index, _, max_index := Minmax(groupa)
		expected_max_index := 4
		expected_min_index := 3

		gsl_test(max_index != expected_max_index,
			"Minmax max (%u observed vs %u expected)",
			max_index, expected_max_index)
		gsl_test(min_index != expected_min_index,
			"Minmax min (%u observed vs %u expected)",
			min_index, expected_min_index)
	}

	sorted1 := make(Float64, stridea*na)
	sorted := NewStrider(&sorted1, stridea)

	for i = 0; i < na; i++ {
		sorted.SetValue(i, groupa.Value(i))
	}

	sort.Sort(sorted)

	{
		median := MedianFromSortedData(sorted)
		expected := 0.07505
		gsl_test_rel(median, expected, rel,
			"MedianFromSortedData (even)")
	}

	{
		zeroth := QuantileFromSortedData(sorted, 0.0)
		expected := 0.0242
		gsl_test_rel(zeroth, expected, rel,
			"QuantileFromSortedData (0)")
	}

	{
		top := QuantileFromSortedData(sorted, 1.0)
		expected := 0.1331
		gsl_test_rel(top, expected, rel,
			"QuantileFromSortedData (100)")
	}

	{
		median := QuantileFromSortedData(sorted, 0.5)
		expected := 0.07505
		gsl_test_rel(median, expected, rel,
			"QuantileFromSortedData (50even)")
	}

	// Test for IEEE handling - set third element to NaN

	groupa.SetValue(3, math.NaN())

	{
		max, max_index := Max(groupa)
		expected := math.NaN()
		expected_index := 3

		gsl_test(!math.IsNaN(max),
			"Max NaN (%g observed vs %g expected)",
			max, expected)

		gsl_test(max_index != expected_index,
			"Max NaN index (%d observed vs %d expected)",
			max_index, expected_index)
	}

	{
		min, min_index := Min(groupa)
		expected := math.NaN()
		expected_index := 3

		gsl_test(!math.IsNaN(min),
			"Min NaN (%g observed vs %g expected)",
			min, expected)

		gsl_test(min_index != expected_index,
			"Min NaN index (%d observed vs %d expected)",
			min_index, expected_index)
	}

	{
		min, min_index, max, max_index := Minmax(groupa)
		expected_max := math.NaN()
		expected_min := math.NaN()
		expected_max_index := 3
		expected_min_index := 3

		gsl_test(!math.IsNaN(max),
			"Minmax max NaN (%g observed vs %g expected)",
			max, expected_max)
		gsl_test(!math.IsNaN(min),
			"Minmax min NaN (%g observed vs %g expected)",
			min, expected_min)

		gsl_test(max_index != expected_max_index,
			"Minmax max index NaN (%u observed vs %u expected)",
			max_index, expected_max_index)
		gsl_test(min_index != expected_min_index,
			"Minmax min index NaN (%u observed vs %u expected)",
			min_index, expected_min_index)
	}
}
