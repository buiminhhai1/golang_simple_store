package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomInt generates a random integer between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomString generates a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomEmail generates a random email
func RandomEmail() string {
	var sb strings.Builder
	sb.WriteString(RandomString(10))
	sb.WriteString("@gmail.com")
	return sb.String()
}

// RandomRole generates a random role between admin and customer
func RandomRole() string {
	roles := []string{"admin", "customer"}
	n := len(roles)
	return roles[rand.Intn(n)]
}

// RandomStatusBook ganerates a random status in array [out_of_stock, in_stock, running_low]
func RandomStatusBook() string {
	status := []string{"out_of_stock", "in_stock", "running_low"}
	n := len(status)
	return status[rand.Intn(n)]
}

// RandomOrderStatus generates a random status in array
func RandomOrderStatus() string {
	status := []string{"pending", "process", "done"}
	n := len(status)
	return status[rand.Intn(n)]
}
