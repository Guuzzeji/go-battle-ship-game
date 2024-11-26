package routes

import (
	"math/rand"
	"time"
)

// Used for random string generation
var seedRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

// RandomString generates a random string of the specified length using a set of
// alphanumeric characters (both uppercase and lowercase letters, and digits).
func randomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}
	return string(result)
}
