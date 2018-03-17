package chop

import "testing"

func TestChop(t *testing.T) {
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
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := chop(tc.value, tc.array)
			if got != tc.want {
				t.Errorf("Got %d; want %d.", got, tc.want)
			}
		})
	}
}

//   assert_equal(0,  chop(1, [1, 3, 5]))
//   assert_equal(1,  chop(3, [1, 3, 5]))
//   assert_equal(2,  chop(5, [1, 3, 5]))
//   assert_equal(-1, chop(0, [1, 3, 5]))
//   assert_equal(-1, chop(2, [1, 3, 5]))
//   assert_equal(-1, chop(4, [1, 3, 5]))
//   assert_equal(-1, chop(6, [1, 3, 5]))
//   #
//   assert_equal(0,  chop(1, [1, 3, 5, 7]))
//   assert_equal(1,  chop(3, [1, 3, 5, 7]))
//   assert_equal(2,  chop(5, [1, 3, 5, 7]))
//   assert_equal(3,  chop(7, [1, 3, 5, 7]))
//   assert_equal(-1, chop(0, [1, 3, 5, 7]))
//   assert_equal(-1, chop(2, [1, 3, 5, 7]))
//   assert_equal(-1, chop(4, [1, 3, 5, 7]))
//   assert_equal(-1, chop(6, [1, 3, 5, 7]))
//   assert_equal(-1, chop(8, [1, 3, 5, 7]))
// end
