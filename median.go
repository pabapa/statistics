package statistics

/* statistics/median.go
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
