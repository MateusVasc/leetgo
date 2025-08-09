package simplearraysum

import "testing"

func TestSimpleArraySum(t *testing.T) {
	tests := []struct {
		name     string
		arr      []int
		expected int
	}{
		{
			"Empty array",
			[]int{},
			0,
		},
		{
			"Normal array",
			[]int{1, 2, 3, 4, 5},
			15,
		},
		{
			"Array with negative values",
			[]int{1, 2, 3, 4, 5, -3, -4},
			8,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SimpleArraySum(tt.arr)

			if got != tt.expected {
				t.Errorf("Got: %d, expected: %d", got, tt.expected)
			}
		})
	}
}
