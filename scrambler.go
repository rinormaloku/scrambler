package scrambler

import (
	"math/rand"
	"strings"
	"time"
)

func Scramble(text string) string {
	slice := strings.Split(text, "")
	rand.Seed(time.Now().UnixNano())
	for i := len(slice) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		slice[i], slice[j] = slice[j], slice[i]
	}

	return strings.Join(slice, "")
}
