package memberships

import (
	"github.com/fuadvi/music-catalog/internal/configs"
	"github.com/fuadvi/music-catalog/internal/models/memberships"
)

//go:generate mockgen -source=service.go -destination=service_mock_test.go -package=memberships
type repository interface {
	CreateUser(model memberships.User) error
	GetUser(email, username string, id uint) (*memberships.User, error)
}

type Service struct {
	cfg        *configs.Config
	repository repository
}

func NewService(cfg *configs.Config, repository repository) *Service {
	return &Service{
		cfg:        cfg,
		repository: repository,
	}
}
