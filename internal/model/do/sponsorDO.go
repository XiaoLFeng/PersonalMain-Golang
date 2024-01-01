package do

import "time"

type SponsorDO struct {
	Id        *uint64   `json:"id"`
	Name      string    `json:"name"`
	UserId    *uint64   `json:"user_id"`
	Url       *string   `json:"url"`
	Type      uint8     `json:"type"`
	Money     uint64    `json:"money"`
	CreatedAt time.Time `json:"created_at"`
	Check     bool      `json:"check"`
}
