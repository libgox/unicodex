package number

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsDigit(t *testing.T) {
	// Test true cases for common digits
	assert.True(t, IsDigit('0'))
	assert.True(t, IsDigit('5'))
	assert.True(t, IsDigit('9'))

	// Test false cases for non-digit characters
	assert.False(t, IsDigit('a'))
	assert.False(t, IsDigit('Z'))
	assert.False(t, IsDigit('!'))

	assert.False(t, IsDigit('٤')) // Arabic-Indic digit 4
	assert.False(t, IsDigit('७')) // Devanagari digit 7

	// Test false case for Unicode numeric symbols (not categorized as Nd)
	assert.False(t, IsDigit('Ⅳ')) // Roman numeral IV (not a decimal digit)

	// Test edge cases for characters near digit ranges
	assert.False(t, IsDigit(' '))
	assert.False(t, IsDigit('-'))
}
