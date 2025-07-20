package palindromenumber

import (
	"reflect"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name string
		num  int
		want bool
	}{
		{
			name: "Exemplo palíndromo",
			num:  121,
			want: true,
		},
		{
			name: "Exemplo não palíndromo",
			num:  -121,
			want: false,
		},
		{
			name: "Exemplo não palíndromo 2",
			num:  10,
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isPalindrome(tt.num)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("isPalindrome(%v) = %v; want %v", tt.num, got, tt.want)
			}
		})
	}
}
