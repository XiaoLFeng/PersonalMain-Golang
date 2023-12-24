package do

import (
	"time"
)

type TokenDO struct {
	Id        *int64    `json:"id"`
	UserId    *int64    `json:"user_id"`
	Token     string    `json:"token"`
	ExpiredAt time.Time `json:"expired_at"`
	CreatedAt time.Time `json:"created_at"`
}
