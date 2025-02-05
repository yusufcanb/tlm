package internal

const (
	// sampleSize is the number of bytes to check before making a decision
	sampleSize = 512
	// threshold is the percentage of non-text bytes that qualify a file as binary
	binaryThreshold = 0.30
	// nullThreshold is the percentage of null bytes that qualify a file as binary
	nullThreshold = 0.50
)

// IsBinary determines if the given byte slice represents binary content.
//
// This function inspects the first 512 bytes of the content for null bytes or non-printable
// characters, which are common indicators of binary files. It assumes UTF-8 or similar encoding,
// where control characters outside the printable ASCII range indicate binary data.
func isBinary(data []byte) bool {
	if len(data) == 0 {
		return false
	}

	// Limit the check to the first sampleSize bytes
	size := len(data)
	if size > sampleSize {
		size = sampleSize
	}

	nonText := 0
	nullCount := 0

	// Check each byte in the sample
	for i := 0; i < size; i++ {
		b := data[i]

		// Count null bytes
		if b == 0 {
			nullCount++
		}

		// Check for non-printable characters, excluding common whitespace
		if !isPrintable(b) && !isCommonWhitespace(b) {
			nonText++
		}
	}

	nullPercentage := float64(nullCount) / float64(size)
	nonTextPercentage := float64(nonText) / float64(size)

	return nullPercentage > nullThreshold || nonTextPercentage > binaryThreshold
}

func isPrintable(b byte) bool {
	return b >= 32 && b <= 126
}

func isCommonWhitespace(b byte) bool {
	// Common whitespace characters: space, tab, newline, carriage return
	return b == 32 || b == 9 || b == 10 || b == 13
}
