package main

import "testing"

func TestCardinality(t *testing.T) {
	_vals := map[string][]int{
		"data1": []int{1, 2, 3, 4},
		"data2": []int{5, 6, 7, 8},
		"data3": []int{9, 10, 11, 12},
	}

	type testDataS struct {
		expected     int
		exptectedErr bool
		key          string
		h            histogram
	}
	testData := []testDataS{
		{
			expected:     4,
			exptectedErr: false,
			key:          "data1",
			h: histogram{
				vals: _vals,
			},
		},
		{
			expected:     0,
			exptectedErr: true,
			key:          "data5",
			h: histogram{
				vals: _vals,
			},
		},
	}

	for _, test := range testData {
		h := NewHistogram(10, test.h.vals)
		got, err := h.Cardinality(test.key)
		if err != nil && !test.exptectedErr {
			t.Fatalf("error: %+v", err)
		}

		if got != test.expected {
			t.Fatalf("expected %+v, got %+v", test.expected, got)
		}
	}
}

func TestPush(t *testing.T) {
	_vals := map[string][]int{
		"data1": []int{1, 2, 3, 4},
		"data3": []int{},
	}
	h := NewHistogram(4, _vals)

	{
		got := h.Push("data1", 4)
		if *got != 4 {
			t.Fatalf("invalid result 1")
		}
	}
	{
		got := h.Push("data3", 1)
		if got != nil {
			t.Fatal("invalid result 3")
		}
	}
}
