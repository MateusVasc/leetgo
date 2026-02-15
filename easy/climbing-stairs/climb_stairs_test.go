package climbingstairs

import (
	"reflect"
	"testing"
)

func TestClimbStairs(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{
			name: "Exemplo 1",
			n:    2,
			want: 2,
		},
		{
			name: "Exemplo 2",
			n:    3,
			want: 3,
		},
		{
			name: "Exemplo 3",
			n:    4,
			want: 5,
		},
		{
			name: "Exemplo 4",
			n:    5,
			want: 8,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := climbStairs(tt.n)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("climbStairs(%d) got %d wanted %d", tt.n, got, tt.want)
			}
		})
	}
}

// leetcode.com/problems/climbing-stairs/
