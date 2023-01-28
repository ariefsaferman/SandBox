package randomizer

import "math/rand"

var allowedCharacters = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

type Randomizer interface {
	RandomStr(n int) string
}

type randomizer struct {
}

func NewRandomizer(seed int64) Randomizer {
	return &randomizer{}
}

func (r *randomizer) RandomStr(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = allowedCharacters[rand.Intn(len(allowedCharacters))]
	}

	return string(b)
}
