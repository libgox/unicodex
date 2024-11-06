package letter

import "math/rand"

func RandLetters(letters []rune, count int) string {
	result := make([]rune, count)
	for i := 0; i < count; i++ {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}
