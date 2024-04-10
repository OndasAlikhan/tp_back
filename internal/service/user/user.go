package user

import (
	"log/slog"

	"tp_back/internal/domain"
	"tp_back/internal/exception"
	"tp_back/internal/lib/sl"
	"tp_back/internal/service/hash"
)

type service struct {
	log        *slog.Logger
	repository Repository
	hService   hash.Service
}

func New(log *slog.Logger, repository Repository, hService hash.Service) Service {
	return &service{log: log, repository: repository, hService: hService}
}

func (s *service) Register(dto RegisterDTO) (*domain.User, error) {
	const op = "service.user.Register"

	log := s.log.With(
		slog.String("op", op),
		slog.Any("dto", dto),
	)

	log.Info("checking exists user by email")
	if s.repository.ExistsByEmail(ExistsByEmailDTO{Email: dto.Email}) {
		log.Error("user by email " + dto.Email + " already exists")
		return nil, exception.ErrAlreadyExists
	}
	log.Info("checked exists user by email")

	log.Info("hashing user password")
	pHash, err := s.hService.Hash(hash.HashDTO{Password: dto.Password})
	if err != nil {
		log.Error("failed to hash user password: ", sl.Error(err))
		return nil, err
	}
	log.Info("hashed user password", slog.String("hashed password", pHash))

	log.Info("saving user")
	uuid, err := s.repository.Save(SaveDTO{Email: dto.Email, PasswordHash: pHash})
	if err != nil {
		log.Error("failed to save user: ", sl.Error(err))
		return nil, err
	}
	log.Info("saved user", slog.String("uuid", uuid))

	log.Info("finding user by uuid")
	user, err := s.repository.FindByUuid(FindByUuidDTO{Uuid: uuid})
	if err != nil {
		log.Error("failed to found user by uuid: ", sl.Error(err))
		return nil, err
	}
	log.Info("found user by uuid", slog.Any("user", user))

	return user, nil
}

func (s *service) Login(dto LoginDTO) (string, error) {
	const op = "service.user.Login"

	log := s.log.With(
		slog.String("op", op),
		slog.Any("dto", dto),
	)

	log.Info("finding user by email")
	user, err := s.repository.FindByEmail(FindByEmailDTO{Email: dto.Email})
	if err != nil {
		log.Error("failed to found user by email: ", sl.Error(err))
		return "", err
	}
	log.Info("found user by email")

	log.Info("checking user password")
	if !s.hService.Check(hash.CheckDTO{Password: dto.Password, Hash: user.Password}) {
		log.Error("invalid credential: ", slog.Any("dto", dto))
		return "", exception.ErrInvalidCred
	}
	log.Info("checked user password")

	log.Info("generating token")
	token, err := s.hService.Token(hash.TokenDTO{Uuid: user.Uuid, Email: user.Email})
	if err != nil {
		log.Error("failed to generate token: ", sl.Error(err))
		return "", err
	}
	log.Info("generated token")

	return token, nil
}
