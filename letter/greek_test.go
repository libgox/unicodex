package letter

import "testing"

func TestRandGreekUpperLetters_Length(t *testing.T) {
	count := 10
	length := 20
	result := RandGreekUpper(count)
	if len(result) != length {
		t.Errorf("Expected length %d, but got %d", length, len(result))
	}
}

func TestRandGreekLowerLetters_Length(t *testing.T) {
	count := 10
	length := 20
	result := RandGreekLower(count)
	if len(result) != length {
		t.Errorf("Expected length %d, but got %d", length, len(result))
	}
}

func TestRandGreekLetters_Length(t *testing.T) {
	count := 10
	length := 20
	result := RandGreek(count)
	if len(result) != length {
		t.Errorf("Expected length %d, but got %d", length, len(result))
	}
}
