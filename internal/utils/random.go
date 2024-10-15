package utils

import (
	"math/rand"
	"strings"
	"time"
)


const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"


func RandomInt(min int64, max int64) int64 {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	return min + rng.Int63n(max -min+1)
}


func RandomString(length int) string {
	var sb strings.Builder
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		sb.WriteByte(alphabet[rng.Intn(len(alphabet))])
	}
	return sb.String()
}

func RandomOwner() string {
	return RandomString(6)
}

func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

func RandomCurrency() string {
	currencies := []string{USD, EUR, CAD}
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	n := len(currencies)
	return currencies[rng.Intn(n)]
}

func RandomEmail() string {
	return RandomString(6) + "@" + "email" + ".com"
}