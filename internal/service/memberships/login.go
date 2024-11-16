package memberships

import (
	"errors"
	"github.com/fuadvi/music-catalog/internal/models/memberships"
	"github.com/fuadvi/music-catalog/pkg/jwt"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func (s *Service) Login(request memberships.LoginRequest) (string, error) {
	userDetail, err := s.repository.GetUser(request.Email, "", 0)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Error().Err(err).Msg("Failed to get user")
		return "", err
	}

	if userDetail == nil {
		return "", errors.New("user not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userDetail.Password), []byte(request.Password)); err != nil {
		return "", errors.New("email or password not match")
	}

	accessToken, err := jwt.CreateToken(int64(userDetail.ID), userDetail.Username, s.cfg.Service.SecretJwt)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create token")
		return "", err
	}

	return accessToken, nil
}
