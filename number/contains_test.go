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
	assert.True(t, ContainsDigit("٤")) // Arabic-Indic digit 4
	assert.True(t, ContainsDigit("７")) // Full-width digit 7 (used in East Asian scripts)

	// Test cases without digits
	assert.False(t, ContainsDigit("abcdef"))
	assert.False(t, ContainsDigit("Hello World!"))
	assert.False(t, ContainsDigit("!@#$%^&*"))
	assert.False(t, ContainsDigit(""))

	// Mixed cases
	assert.True(t, ContainsDigit("Test1")) // Single digit
	assert.False(t, ContainsDigit("Test")) // No digit
	assert.True(t, ContainsDigit("１２３"))   // Full-width digits
	assert.False(t, ContainsDigit("Ⅷ"))    // Roman numeral VIII (not a decimal digit)
}
