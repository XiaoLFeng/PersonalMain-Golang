package do

import "time"

type UserDO struct {
	Id          *uint64    `json:"id"`
	UserName    string     `json:"username"`
	DisplayName *string    `json:"display_name"`
	Email       string     `json:"email"`
	Qq          *string    `json:"qq"`
	Password    string     `json:"password"`
	OldPassword *string    `json:"old_password"`
	Permission  bool       `json:"permission"`
	EmailVerify bool       `json:"email_verify"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
}
