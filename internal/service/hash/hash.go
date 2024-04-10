package hash

import (
	"log/slog"
	"time"

	"tp_back/internal/config"

	"github.com/form3tech-oss/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type service struct {
	log    *slog.Logger
	jwtCfg config.Jwt
}

func New(log *slog.Logger, jwtCfg config.Jwt) Service {
	return &service{log: log, jwtCfg: jwtCfg}
}

func (s *service) Hash(dto HashDTO) (string, error) {
	const op = "service.hash.Hash"

	log := s.log.With(
		slog.String("op", op),
		slog.Any("dto", dto),
	)

	log.Info("generating hash from password")
	bytes, err := bcrypt.GenerateFromPassword([]byte(dto.Password), 14)
	if err != nil {
		log.Error("failed to generate hash from password")
		return "", err
	}
	pHash := string(bytes)
	log.Info("generated hash from password", slog.String("hashed password", pHash))

	return pHash, nil
}

func (s *service) Check(dto CheckDTO) bool {
	const op = "service.hash.Check"

	log := s.log.With(
		slog.String("op", op),
		slog.Any("dto", dto),
	)

	log.Info("comparing hash and password")
	if err := bcrypt.CompareHashAndPassword([]byte(dto.Hash), []byte(dto.Password)); err != nil {
		log.Error("failed to compare hash and password", slog.Any("err", err))
		return false
	}
	log.Info("compared hash and password")

	return true
}

func (s *service) Token(dto TokenDTO) (string, error) {
	const op = "service.hash.Token"

	log := s.log.With(
		slog.String("op", op),
		slog.Any("dto", dto),
	)

	log.Info("creating token")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uuid":  dto.Uuid,
		"email": dto.Email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenS, err := token.SignedString([]byte(s.jwtCfg.SecretKey))
	if err != nil {
		log.Error("failed to create token", slog.Any("err", err))
		return "", err
	}
	log.Info("created token")

	return tokenS, nil
}
