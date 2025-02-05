package internal

import "regexp"

// countTokensApprox uses a regular expression to split the input text into token‐like pieces.
// The regex below is inspired by GPT‑2’s tokenization rules and will match common contractions,
// words, numbers, and punctuation.
func GetTokenCount(text string) int {
	// This regex matches common English contractions and splits words and punctuation.
	// It is not perfect, but gives a rough approximation.
	re := regexp.MustCompile(`('s|'t|'re|'ve|'m|'ll|'d| ?\w+| ?[^\w\s]+)`)
	tokens := re.FindAllString(text, -1)
	return len(tokens)
}
