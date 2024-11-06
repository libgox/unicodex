package letter

import "testing"

func TestRandEnglishUpperLetters_Length(t *testing.T) {
	count := 10
	length := 10
	result := RandEnglishUpper(count)
	if len(result) != length {
		t.Errorf("Expected length %d, but got %d", length, len(result))
	}
}

func TestRandEnglishLowerLetters_Length(t *testing.T) {
	count := 10
	length := 10
	result := RandEnglishLower(count)
	if len(result) != length {
		t.Errorf("Expected length %d, but got %d", length, len(result))
	}
}

func TestRandEnglishLetters_Length(t *testing.T) {
	count := 10
	length := 10
	result := RandEnglish(count)
	if len(result) != length {
		t.Errorf("Expected length %d, but got %d", length, len(result))
	}
}
