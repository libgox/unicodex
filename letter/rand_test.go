package letter

import "testing"

func TestRandLetters_Count(t *testing.T) {
	letters := []rune("abcdefghijklmnopqrstuvwxyz")
	count := 10
	result := RandLetters(letters, count)
	if len(result) != count {
		t.Errorf("Expected count %d, but got %d", count, len(result))
	}
}

func TestRandLetters_ValidCharacters(t *testing.T) {
	letters := []rune("abcdefghijklmnopqrstuvwxyz")
	count := 10
	result := RandLetters(letters, count)
	for _, r := range result {
		if !containsRune(letters, r) {
			t.Errorf("Unexpected character '%c' in result", r)
		}
	}
}

func containsRune(slice []rune, r rune) bool {
	for _, sr := range slice {
		if sr == r {
			return true
		}
	}
	return false
}
