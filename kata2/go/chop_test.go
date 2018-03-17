package chop

import "testing"

func TestChop(t *testing.T) {
	testCases := []struct{
		want int
		value int
		array []int
	} {
		{-1, 3, nil},
		{-1, 3, []int{}},
		{-1, 3, []int{1}},
		{0, 1, []int{1}},
	}

	for _, tc := range testCases {
		t.Run("",func (t *testing.T){
			got := chop(tc.value, tc.array)
			if got != tc.want {
				t.Errorf("Got %d; want %d.", got, tc.want)
			}
		})
	}
}

// assert_equal(-1, chop(3, []))
//   assert_equal(-1, chop(3, [1]))
//   assert_equal(0,  chop(1, [1]))
//   #
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
