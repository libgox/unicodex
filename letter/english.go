package letter

const EnglishUpperMin = 'A'
const EnglishUpperMax = 'Z'
const EnglishLowerMin = 'a'
const EnglishLowerMax = 'z'

const EnglishUpperCount = 26
const EnglishLowerCount = 26
const EnglishCount = 52

const EnglishUpperLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const EnglishLowerLetters = "abcdefghijklmnopqrstuvwxyz"
const EnglishLetters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

var EnglishUpperLetterRunes = []rune{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}
var EnglishLowerLetterRunes = []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
var EnglishLetterRunes = []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}

func RandEnglishUpper(count int) string {
	return RandLetters(EnglishUpperLetterRunes, count)
}

func RandEnglishLower(count int) string {
	return RandLetters(EnglishLowerLetterRunes, count)
}

func RandEnglish(count int) string {
	return RandLetters(EnglishLetterRunes, count)
}
