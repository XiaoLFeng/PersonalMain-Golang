package do

import "time"

type BlogDO struct {
	Id           *uint64    `json:"id"`
	BlogTitle    string     `json:"blog_title"`
	BlogUrl      string     `json:"blog_url"`
	BlogDesc     string     `json:"blog_desc"`
	BlogEmail    *string    `json:"blog_email"`
	BlogAvatar   string     `json:"blog_avatar"`
	BlogRssJudge bool       `json:"blog_rss_judge"`
	BlogRss      *string    `json:"blog_rss"`
	BlogHost     *string    `json:"blog_host"`
	BlogLocation uint       `json:"blog_location"`
	BlogAddType  uint8      `json:"blog_add_type"`
	BlogColor    uint       `json:"blog_color"`
	BlogUserId   *uint64    `json:"blog_user_id"`
	BlogRemark   *string    `json:"blog_remark"`
	CreatedAt    *time.Time `json:"created_at"`
	UpdatedAt    *time.Time `json:"updated_at"`
}
