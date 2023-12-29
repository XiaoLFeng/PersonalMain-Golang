package do

import "time"

type SponsorDO struct {
	Id        *uint64   `json:"id"`
	Name      string    `json:"name"`
	UserId    *uint64   `json:"user_id"`
	Url       *string   `json:"url"`
	Type      uint8     `json:"type"`
	Amount    uint64    `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
	Check     uint8     `json:"check"`
}
