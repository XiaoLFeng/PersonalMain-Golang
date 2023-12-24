package auth

import (
	"PersonalMain/api"
)

type ControllerV1 struct{}

func NewAuthV1() api.IAuthV1 {
	return &ControllerV1{}
}
