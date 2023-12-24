package token

import (
	"PersonalMain/api"
)

type ControllerV1 struct{}

func NewTokenV1() api.ITokenV1 {
	return &ControllerV1{}
}
