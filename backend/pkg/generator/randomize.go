package generator

import (
	"math/rand"
	"time"
)

const letterBytes = "ABCDEFGHIJKLMNOPQRSTUVWXYZ01234567890"
const numberBytes = "1234567890"
var letterBytesShort = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz01234567890")
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

// RandomStringGenerator generates a random sequence of string
func RandomStringGenerator(n int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, n)
	// A rand.Int63() generates 63 random bits, enough for letterIdxMax letters!
	for i, cache, remain := n-1, rand.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = rand.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	
	return string(b)
}

// ShortLinkIdGenerator generates a short link as id
func ShortLinkIdGenerator(n int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, n)
	// A rand.Int63() generates 63 random bits, enough for letterIdxMax letters!
	for i := range b {
		b[i] = letterBytesShort[rand.Intn(len(letterBytesShort))]
	}
	
	return string(b)
}

// GenerateAPIKeys generates API Keys
func GenerateAPIKeys(n int) ([]byte, error) {
	if n == 0 {
		n = 15
	}
	
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	
	return b, nil
}