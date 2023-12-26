package do

import (
	"time"
)

type TokenDO struct {
	Id        *uint64   `json:"id"`
	UserId    *uint64   `json:"user_id"`
	Token     string    `json:"token"`
	ExpiredAt time.Time `json:"expired_at"`
	CreatedAt time.Time `json:"created_at"`
}
