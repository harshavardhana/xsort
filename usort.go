package usort

import "sort"

// uniqueSort returns n, the number of distinct elements in data in
// sorted order.
func uniqueSort(data sort.Interface) (n int) {
	// Slice of length less than '2' return quickly.
	if n = data.Len(); n < 2 {
		return n
	}
	// Sort the data.
	sort.Sort(data)
	a, b := 0, 1
	// Loop through each entries.
	for b < n {
		// Less reports whether the element with
		// index a should sort before the element with index b.
		if data.Less(a, b) {
			a++
			// For indexes not equal swap the entries.
			if a != b {
				data.Swap(a, b)
			}
		}
		// Increment and move to the next element.
		b++
	}
	// Return distinct number of elements which are unique.
	return a + 1
}

// UniqueStrings sorts a slice of strings uniquely in increasing order
// and returns that slice.
func UniqueStrings(a []string) []string {
	n := uniqueSort(sort.StringSlice(a))
	return a[0:n]
}

// UniqueInts sorts a slice of ints uniquely in increasing order and
// returns that slice.
func UniqueInts(a []int) []int {
	n := uniqueSort(sort.IntSlice(a))
	return a[0:n]
}

// UniqueFloat64s sorts a slice of float64s uniquely in increasing order and
// returns that slice.
func UniqueFloat64s(a []float64) []float64 {
	n := uniqueSort(sort.Float64Slice(a))
	return a[0:n]
}
