package do

import "time"

type SponsorTypeDO struct {
	Id        *uint64    `json:"id"`
	Name      string     `json:"name"`
	Url       string     `json:"url"`
	Include   bool       `json:"include"`
	Link      bool       `json:"link"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}
