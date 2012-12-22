package statistics

/* statistics/int64_test.go
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
	"sort"
	"testing"
)

func testInt64(t *testing.T, stridea, strideb int) {

	ina := 20
	inb := 20

	// sample sets of integers

	raw1 := Int64{1, 2, 3, 4, 5, 6}

	irawa := Int64{
		17, 18, 16, 18, 12,
		20, 18, 20, 20, 22,
		20, 10, 8, 12, 16,
		16, 18, 20, 18, 21}

	irawb := Int64{
		19, 20, 22, 24, 10,
		25, 20, 22, 21, 23,
		20, 10, 12, 14, 12,
		20, 22, 24, 23, 17}

	test1_org := make(Int64, stridea*6)
	test1 := NewStrider(&test1_org, stridea)
	igroupa_org := make(Int64, stridea*ina)
	igroupa := NewStrider(&igroupa_org, stridea)
	igroupb_org := make(Int64, strideb*inb)
	igroupb := NewStrider(&igroupb_org, strideb)

	rel := 1e-10

	for i := 0; i < ina; i++ {
		igroupa.SetValue(i, float64(irawa[i]))
	}

	for i := 0; i < inb; i++ {
		igroupb.SetValue(i, float64(irawb[i]))
	}

	for i := 0; i < 6; i++ {
		test1.SetValue(i, float64(raw1[i]))
	}

	{
		mean := Mean(igroupa)
		expected := 17.0
		gsl_test_rel(mean, expected, rel, "Mean(integer)")
	}

	{
		mean := Mean(test1)
		expected := 3.5
		gsl_test_rel(mean, expected, rel, "Mean(fractional)")
	}

	{
		mean := Mean(igroupa)
		variance := VarianceWithFixedMean(igroupa, mean)
		expected := 13.7
		gsl_test_rel(variance, expected, rel, "variance_with_fixed_mean")
	}

	{
		mean := Mean(igroupa)
		sd := SdWithFixedMean(igroupa, mean)
		expected := 3.70135110466435
		gsl_test_rel(sd, expected, rel, "sd_with_fixed_mean")
	}

	{
		variance := Variance(igroupa)
		expected := 14.4210526315789
		gsl_test_rel(variance, expected, rel, "variance")
	}

	{
		sd_est := Sd(igroupa)
		expected := 3.79750610685209
		gsl_test_rel(sd_est, expected, rel, "sd")
	}

	{
		absdev := Absdev(igroupa)
		expected := 2.9
		gsl_test_rel(absdev, expected, rel, "absdev")
	}

	{
		skew := Skew(igroupa)
		expected := -0.909355923168064
		gsl_test_rel(skew, expected, rel, "skew")
	}

	{
		kurt := Kurtosis(igroupa)
		expected := -0.233692524908094
		gsl_test_rel(kurt, expected, rel, "kurtosis")
	}

	{
		c := Covariance(igroupa, igroupb)
		expected := 14.5263157894737
		gsl_test_rel(c, expected, rel, "covariance")
	}

	{
		pv := PVariance(igroupa, igroupb)
		expected := 18.8421052631579
		gsl_test_rel(pv, expected, rel, "pvariance")
	}

	{
		t := TTest(igroupa, igroupb)
		expected := -1.45701922702927
		gsl_test_rel(t, expected, rel, "ttest")
	}

	{
		maxf, max_index := Max(igroupa)
		max := int64(maxf)
		expected := int64(22)
		expected_index := 9

		gsl_test(max != expected,
			"Max (%d observed vs %d expected)", max, expected)

		gsl_test(max_index != expected_index,
			"Max index (%d observed vs %d expected)",
			max_index, expected_index)
	}

	{
		minf, min_index := Min(igroupa)
		min := int(minf)
		expected := 8
		expected_index := 12

		gsl_test(min != expected,
			"min (%d observed vs %d expected)", min, expected)

		gsl_test(min_index != expected_index,
			"Min index (%d observed vs %d expected)",
			min_index, expected_index)
	}

	{
		minf, min_index, maxf, max_index := Minmax(igroupa)
		min := int(minf)
		max := int(maxf)
		expected_max := 22
		expected_min := 8
		expected_max_index := 9
		expected_min_index := 12

		gsl_test(max != expected_max,
			"minmax max (%d observed vs %d expected)",
			max, expected_max)
		gsl_test(min != expected_min,
			"minmax min (%d observed vs %d expected)",
			min, expected_min)

		gsl_test(max_index != expected_max_index,
			"Minmax index max (%u observed vs %u expected)",
			max_index, expected_max_index)
		gsl_test(min_index != expected_min_index,
			"Minmax index min (%u observed vs %u expected)",
			min_index, expected_min_index)
	}

	sorted_v := make(Int64, stridea*ina)
	sorted := NewStrider(&sorted_v, stridea)

	for i := 0; i < ina; i++ {
		sorted.SetValue(i, igroupa.Value(i))
	}

	sort.Sort(sorted)

	{
		median := MedianFromSortedData(sorted)
		expected := float64(18)
		gsl_test_rel(median, expected, rel,
			"median_from_sorted_data (even)")
	}

	{
		zeroth := QuantileFromSortedData(sorted, 0.0)
		expected := float64(8)
		gsl_test_rel(zeroth, expected, rel,
			"quantile_from_sorted_data (0)")
	}

	{
		top := QuantileFromSortedData(sorted, 1.0)
		expected := float64(22)
		gsl_test_rel(top, expected, rel,
			"quantile_from_sorted_data (100)")
	}

	{
		median := QuantileFromSortedData(sorted, 0.5)
		expected := float64(18)
		gsl_test_rel(median, expected, rel,
			"quantile_from_sorted_data (50, even)")
	}
}
