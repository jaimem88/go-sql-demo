package demo

import (
	"database/sql"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/jaimemartinez88/go-sql-demo/pkg/types"
	uuid "github.com/satori/go.uuid"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

var numRunes = []rune("0123456789")

const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

func randStr(n int) string {
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

func randNum(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = numRunes[rand.Intn(len(numRunes))]
	}
	return string(b)
}
func randBool() bool {
	return rand.Float32() < 0.5
}
func randUser() *types.User {

	u := &types.User{
		Name:  randStr(5) + " " + randStr(6),
		Email: fmt.Sprintf("%s@%s.com", randStr(5), randStr(5)),
		Admin: randBool(),
	}
	if randBool() {
		u.Mobile = sql.NullString{String: randNum(10), Valid: true}
	}
	if randBool() {
		v, _ := strconv.Atoi(randNum(2))
		u.Age = sql.NullInt64{Int64: int64(v), Valid: true}
	}
	return u
}
func randAddr(userID uuid.UUID) *types.Address {
	return &types.Address{
		UserID:        userID,
		StreetAddress: randStr(12),
		Suburb:        randStr(6),
		State:         randStr(3),
		Postcode:      randNum(4),
		Country:       randStr(2),
	}
}
