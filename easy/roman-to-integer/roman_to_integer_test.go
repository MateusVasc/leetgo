package romantointeger

import "testing"

func TestRomanToInt(t *testing.T) {
	tests := []struct {
		name     string
		romanVal string
		want     int
	}{
		{
			name:     "Example 1",
			romanVal: "III",
			want:     3,
		},
		{
			name:     "Example 2",
			romanVal: "LVIII",
			want:     58,
		},
		{
			name:     "Example 3",
			romanVal: "MCMXCIV",
			want:     1994,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RomanToInt(tt.romanVal)

			if got != tt.want {
				t.Errorf("RomanToInt(%v) = %v; want %v", tt.romanVal, got, tt.want)
			}
		})
	}
}
