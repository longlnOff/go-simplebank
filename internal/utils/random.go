package utils

import (
	"math/rand"
	"strings"
)


const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"


func RanomInt(min int64, max int64) int64 {
	return min + rand.Int63n(max -min+1)
}


func RandomString(length int) string {
	var sb strings.Builder
	for i := 0; i < length; i++ {
		sb.WriteByte(alphabet[rand.Intn(len(alphabet))])
	}
	return sb.String()
}

func RandomOwner() string {
	return RandomString(6)
}

func RandomMoney() int64 {
	return RanomInt(0, 1000)
}

func RandomCurrency() string {
	currencies := []string{"USD", "EUR", "CAD"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}