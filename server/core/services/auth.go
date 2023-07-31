package services

import "github.com/hashicorp/go-hclog"

type authService struct {
	l hclog.Logger
}

func NewAuthService(l hclog.Logger) AuthService {
	return &authService{l}
}
