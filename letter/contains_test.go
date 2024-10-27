package letter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContainsEnglishUpper(t *testing.T) {
	assert.True(t, ContainsEnglishUpper("Hello"))
	assert.False(t, ContainsEnglishUpper("hello"))
	assert.False(t, ContainsEnglishUpper("1234"))
	assert.False(t, ContainsEnglishUpper("Γειά"))
}

func TestContainsEnglishLower(t *testing.T) {
	assert.True(t, ContainsEnglishLower("hello"))
	assert.False(t, ContainsEnglishLower("HELLO"))
	assert.False(t, ContainsEnglishLower("1234"))
	assert.False(t, ContainsEnglishLower("Γειά"))
}

func TestContainsEnglish(t *testing.T) {
	assert.True(t, ContainsEnglish("Hello"))
	assert.True(t, ContainsEnglish("hello"))
	assert.False(t, ContainsEnglish("1234"))
	assert.False(t, ContainsEnglish("Γειά"))
}

func TestContainsGreekUpper(t *testing.T) {
	assert.True(t, ContainsGreekUpper("Γειά"))
	assert.False(t, ContainsGreekUpper("γειά"))
	assert.False(t, ContainsGreekUpper("Hello"))
	assert.False(t, ContainsGreekUpper("1234"))
}

func TestContainsGreekLower(t *testing.T) {
	assert.True(t, ContainsGreekLower("γειά"))
	assert.False(t, ContainsGreekLower("ΓΕΙΑ"))
	assert.False(t, ContainsGreekLower("Hello"))
	assert.False(t, ContainsGreekLower("1234"))
}

func TestContainsGreek(t *testing.T) {
	assert.True(t, ContainsGreek("Γειά"))
	assert.True(t, ContainsGreek("γειά"))
	assert.False(t, ContainsGreek("Hello"))
	assert.False(t, ContainsGreek("1234"))
}
