package util

import (
	"math/rand"
	"strings"
	"time"
)

const str = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandonInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandonString(n int) string {
	var sb strings.Builder
	k := len(str)
	for i := 0; i < n; i++ {
		c := str[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomOwner() string {
	return RandonString(6)
}

func RandomBalance() int64 {
	return RandonInt(0, 10000)
}

func RandomCurrency() string {
	currency := []string{"EUR", "USD", "GPD"}
	k := len(currency)
	return currency[rand.Intn(k)]
}
