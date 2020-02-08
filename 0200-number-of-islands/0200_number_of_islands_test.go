package leetcode

import "testing"

func TestNumerOfIslands(t *testing.T) {
	var tests = []struct {
		grids        [][]byte
		numOfIslands int
	}{
		{[][]byte{
			[]byte{'1', '0', '0', '0', '0', '1'},
			[]byte{'0', '1', '1', '0', '0', '0'},
			[]byte{'0', '0', '0', '1', '0', '0'},
			[]byte{'1', '0', '0', '1', '0', '1'},
			[]byte{'0', '0', '0', '0', '0', '0'},
			[]byte{'0', '1', '1', '1', '0', '0'},
			[]byte{'0', '0', '0', '1', '1', '0'},
			[]byte{'1', '0', '0', '0', '0', '1'},
		}, 9},
		{[][]byte{
			[]byte{'1', '1', '0', '0', '0'},
			[]byte{'1', '1', '0', '0', '0'},
			[]byte{'0', '0', '1', '0', '0'},
			[]byte{'0', '0', '0', '1', '1'},
		}, 3},
	}
	for _, test := range tests {
		result := numIslands(test.grids)
		if test.numOfIslands != result {
			t.Errorf("numIslands(%+v) return error result %v, %v is expected",
				test.grids, result, test.numOfIslands)
		}
	}
}
