package internal

import "testing"

func TestIsBinary(t *testing.T) {
	tests := []struct {
		name     string
		data     []byte
		expected bool
	}{
		{
			name:     "Empty data",
			data:     []byte{},
			expected: false,
		},
		{
			name:     "Text data",
			data:     []byte("This is a sample text."),
			expected: false,
		},
		{
			name:     "Binary data with null bytes",
			data:     []byte{0, 1, 2, 3, 0, 5, 6, 0, 8, 9, 0},
			expected: true,
		},
		{
			name:     "Binary data with non-printable characters",
			data:     []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			expected: true,
		},
		{
			name:     "Mixed data with text and binary",
			data:     append([]byte("This is a sample text."), 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15),
			expected: true,
		},
		{
			name:     "Text data with common whitespace",
			data:     []byte("This is a sample text with spaces, tabs\t, and newlines\n."),
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isBinary(tt.data)
			if result != tt.expected {
				t.Errorf("isBinary(%v) = %v; expected %v", tt.data, result, tt.expected)
			}
		})
	}
}
