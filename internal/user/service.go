package user

import (
	"context"
	"github.com/Salavei/golang_advanced_restapi/pkg/logging"
)

type Service struct {
	storage Storage
	logger  *logging.Logger
}

func (s *Service) Create(ctx context.Context, dto CreateUserDTO) (u User, err error) {
	// TODO next one
	return
}
