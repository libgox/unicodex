package letter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsEnglishUpper(t *testing.T) {
	assert.True(t, IsEnglishUpper('A'))
	assert.True(t, IsEnglishUpper('Z'))
	assert.False(t, IsEnglishUpper('a'))
	assert.False(t, IsEnglishUpper('z'))
	assert.False(t, IsEnglishUpper('Γ'))
}

func TestIsEnglishLower(t *testing.T) {
	assert.True(t, IsEnglishLower('a'))
	assert.True(t, IsEnglishLower('z'))
	assert.False(t, IsEnglishLower('A'))
	assert.False(t, IsEnglishLower('Z'))
	assert.False(t, IsEnglishLower('γ'))
}

func TestIsEnglish(t *testing.T) {
	assert.True(t, IsEnglish('A'))
	assert.True(t, IsEnglish('z'))
	assert.True(t, IsEnglish('a'))
	assert.False(t, IsEnglish('Γ'))
	assert.False(t, IsEnglish('γ'))
	assert.False(t, IsEnglish('1'))
}

func TestIsGreekUpper(t *testing.T) {
	assert.True(t, IsGreekUpper('Γ'))
	assert.True(t, IsGreekUpper('Ω'))
	assert.False(t, IsGreekUpper('α'))
	assert.False(t, IsGreekUpper('A'))
	assert.False(t, IsGreekUpper('z'))
}

func TestIsGreekLower(t *testing.T) {
	assert.True(t, IsGreekLower('α'))
	assert.True(t, IsGreekLower('ω'))
	assert.False(t, IsGreekLower('Γ')) // Greek uppercase, should be false
	assert.False(t, IsGreekLower('A'))
	assert.False(t, IsGreekLower('z'))
}

func TestIsGreek(t *testing.T) {
	assert.True(t, IsGreek('Γ'))
	assert.True(t, IsGreek('Ω'))
	assert.True(t, IsGreek('α'))
	assert.True(t, IsGreek('ω'))
	assert.False(t, IsGreek('A'))
	assert.False(t, IsGreek('z'))
	assert.False(t, IsGreek('1'))
}
