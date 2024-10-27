package number

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContainsDigit(t *testing.T) {
	// Test cases with digits
	assert.True(t, ContainsDigit("abc123"))
	assert.True(t, ContainsDigit("Hello4World"))
	assert.True(t, ContainsDigit("0"))
	assert.False(t, ContainsDigit("٤"))
	assert.False(t, ContainsDigit("７"))

	// Test cases without digits
	assert.False(t, ContainsDigit("abcdef"))
	assert.False(t, ContainsDigit("Hello World!"))
	assert.False(t, ContainsDigit("!@#$%^&*"))
	assert.False(t, ContainsDigit(""))

	// Mixed cases
	assert.True(t, ContainsDigit("Test1")) // Single digit
	assert.False(t, ContainsDigit("Test")) // No digit
	assert.False(t, ContainsDigit("１２３"))
	assert.False(t, ContainsDigit("Ⅷ")) // Roman numeral VIII (not a decimal digit)
}
