package util

import (
	"math/rand"
	"strings"
	"time"
)

var localRand *rand.Rand

func init() {
	seed := time.Now().UnixNano()
	localRand = rand.New(rand.NewSource(seed))
}

func RandomInt(min, max int64) int64 {
	return localRand.Int63n(max-min) + min
}

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

func RandomName() string {
	return RandomString(10)
}
