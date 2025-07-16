package twosum

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{
			name:   "Exemplo básico",
			nums:   []int{2, 7, 11, 15},
			target: 9,
			want:   []int{0, 1}, // Índices de 2 e 7
		},
		{
			name:   "Números com valores repetidos",
			nums:   []int{3, 2, 4},
			target: 6,
			want:   []int{1, 2}, // Índices de 2 e 4
		},
		{
			name:   "Números negativos",
			nums:   []int{-1, -2, -3, -4, -5},
			target: -8,
			want:   []int{2, 4}, // Índices de -3 e -5
		},
		{
			name:   "Array pequeno",
			nums:   []int{5, 5},
			target: 10,
			want:   []int{0, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := twoSum(tt.nums, tt.target)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum(%v, %d) = %v; want %v", tt.nums, tt.target, got, tt.want)
			}
		})
	}
}
