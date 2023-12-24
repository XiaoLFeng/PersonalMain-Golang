package Processing

import (
	"math/rand"
	"strconv"
	"time"
)

// CreateToken
//
// 生成 token
func CreateToken() string {
	token := strconv.FormatInt(time.Now().UnixMilli(), 10)
	if len(token) <= 13 {
		token = "0" + token
	}
	token += "-"
	for i := 0; i < 20; i++ {
		if rand.Intn(2) == 0 {
			token += strconv.Itoa(rand.Intn(10))
		} else {
			token += string(byte(rand.Intn(26) + 97))
		}
	}
	token += "-"
	for i := 0; i < 14; i++ {
		token += strconv.Itoa(rand.Intn(10))
	}
	return token
}
