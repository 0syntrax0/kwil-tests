package main

import "fmt"

// Background
//   - We have many categories of data.
//   - We receive new values occasionally, and we want to store the N latest
//	   values for each category.
//   - We want to compute statistics within and across categories.
//
// Let's define a type with methods to do this.  Call it `histogram`:

// histogram is a type that keeps various categories of data in slices with
// limited capacity. The methods should compute various statistic within
// categories and across categories.
type histogram struct {
	// hint/suggestion: map of slices
	// - key is a category (string)
	// - slice are the values for that category
	vals map[string][]int

	cap int // capacity of the values list in each category
}

// pt1. constructor and basic methods
//  - use any new data types, standard libraries
//  - define new methods for common operations as needed
//
// NOTE: some of the methods may build off of eachother or new unexported methods!

// NewHistogram creates a new histogram instance for a certain capacity for each
// category's slice. There is no limit on the number of categories, only the number
// of values recorded within each bin.
func NewHistogram(capPerBin int, vals map[string][]int) *histogram {
	return &histogram{
		cap:  capPerBin,
		vals: vals,
	}
}

// Cardinality computes the size of bin k.
func (t *histogram) Cardinality(k string) (int, error) {
	data, found := t.vals[k]
	if found {
		return len(data), nil
	}

	return 0, fmt.Errorf("key '%s' not found", k)
}

// Push prepends a value to the *front* of the slice, popping off and returning
// the last value *if* at capacity, otherwise return a nil pointer.
// FEEL FREE TO REFERENCE https://go.dev/wiki/SliceTricks
func (t *histogram) Push(k string, v int) *int { // oldest values must be popped out (FIFO)
	data, found := t.vals[k]
	if !found || len(data) == 0 {
		t.vals[k] = []int{v}
		return nil
	}

	lastValue := data[len(data)-1]

	if t.cap >= len(data) {
		newSlice := []int{v}
		t.vals[k] = nil
		t.vals[k] = append(newSlice, data[len(data)-1:]...)

		return &lastValue
	}

	return nil
}

func (t *histogram) Mean(k string) int {
	panic("not implemented") // TODO: Implement
}

func (t *histogram) Range(k string) int {
	panic("not implemented") // TODO: Implement
}

func (t *histogram) Max(k string) int {
	data, found := t.vals[k]
	if !found {
		return 0
	}

	// TODO: implement negative values
	maxVal := 0
	for _, val := range data {
		if val > maxVal {
			maxVal = val
		}
	}

	return maxVal
}

// pt2. cross-category aggregate methods:

func (t *histogram) BinSizes() []int {
	// compute the size of each bin
	totals := []int{}

	for _, data := range t.vals {
		totals = append(totals, len(data))
	}

	return totals
}

func (t *histogram) BinWeights() []int {
	// compute the total value (sum) of each bin

	totals := []int{}
	for _, data := range t.vals {
		total := 0
		for _, val := range data {
			total += val
		}

		totals = append(totals, total)
	}

	return totals
}

func (t *histogram) TotalSum() int {
	// compute the total sum of all values across all keys
	var total int
	for _, totalWeight := range t.BinWeights() {
		total += totalWeight
	}

	return total
}

func (t *histogram) TotalMean() int {
	// compute the mean of all values across all categories
	panic("not implemented") // TODO: Implement
}

func (t *histogram) GlobalMax() int {
	// compute the maximum value among all keys and returns the key and the maximum value
	panic("not implemented") // TODO: Implement
}

func (t *histogram) GlobalMin() int {
	// compute the minimum value among all keys and returns the key and the minimum value
	panic("not implemented") // TODO: Implement
}

// pt3. deep copy values and clone

func (t *histogram) Vals(k string) []int {
	// return a deep copy of e.g. t.vals[k]
	panic("not implemented") // TODO: Implement
}

func (t *histogram) Clone() *histogram {
	// deep copy the entire histogram instance
	panic("not implemented") // TODO: Implement
}

// pt4. arbitrary statistics with first class function

// Reduce performs a supplied aggregate operation on each bin, return all results
func (t *histogram) Reduce(k string, statFn func([]int) int) []int {
	return nil
}

// pt5. DISCUSSION:
// - thread safety - considerations for concurrent access?
// - generic implementation that would work for float64?
// - store largest/smallest values instead of FIFO? What data structures? What packages?
// - other improvements?

/* histogram_test.go -- EXPAND ON THIS! here's a start:
package categories
import (
	"testing"
)
var vals = []struct {
	k string
	v int
}{
	{
		"a",
		1,
	},
	{
		"a",
		2,
	},
	{
		"a",
		3,
	},
	{
		"a",
		4,
	},
	{
		"b",
		99,
	},
}
func Test_histogram_Cardinality(t *testing.T) {
	const N = 3
	h := NewHistogram(N)
	for _, kv := range vals {
		h.Push(kv.k, kv.v)
	}
	t.Log(h.Vals("a"))
	t.Log(h.Vals("b"))
}
// TODO: expand on tests and test values
*/
