package chop

import "testing"

func TestChop(t *testing.T) {
	testFunctions := []struct {
		name string
		fn   func(int, []int) int
	}{
		{
			"recursive",
			chopRec,
		},
		{
			"iterative",
			chopIter,
		},
	}

	testCases := []struct {
		name  string
		want  int
		value int
		array []int
	}{
		{"1_1", -1, 3, nil},
		{"1_2", -1, 3, []int{}},
		{"1_3", -1, 3, []int{1}},
		{"1_4", 0, 1, []int{1}},

		{"2_1", 0, 1, []int{1, 3, 5}},
		{"2_2", 1, 3, []int{1, 3, 5}},
		{"2_3", 2, 5, []int{1, 3, 5}},
		{"2_4", -1, 0, []int{1, 3, 5}},
		{"2_5", -1, 2, []int{1, 3, 5}},
		{"2_6", -1, 4, []int{1, 3, 5}},
		{"2_7", -1, 6, []int{1, 3, 5}},

		{"3_1", 0, 1, []int{1, 3, 5, 7}},
		{"3_2", 1, 3, []int{1, 3, 5, 7}},
		{"3_3", 2, 5, []int{1, 3, 5, 7}},
		{"3_4", 3, 7, []int{1, 3, 5, 7}},
		{"3_5", -1, 0, []int{1, 3, 5, 7}},
		{"3_6", -1, 2, []int{1, 3, 5, 7}},
		{"3_7", -1, 4, []int{1, 3, 5, 7}},
		{"3_8", -1, 6, []int{1, 3, 5, 7}},
		{"3_9", -1, 8, []int{1, 3, 5, 7}},
	}

	for _, tFn := range testFunctions {
		for _, tc := range testCases {
			t.Run(tFn.name+"-"+tc.name, func(t *testing.T) {
				got := tFn.fn(tc.value, tc.array)
				if got != tc.want {
					t.Errorf("Got %d; want %d.", got, tc.want)
				}
			})
		}
	}
}
