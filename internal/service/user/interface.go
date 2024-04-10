package user

import (
	"tp_back/internal/domain"
)

type Service interface {
	Register(dto RegisterDTO) (*domain.User, error)
	Login(dto LoginDTO) (string, error)
}

type Repository interface {
	Save(dto SaveDTO) (string, error)
	FindByEmail(dto FindByEmailDTO) (*domain.User, error)
	FindByUuid(dto FindByUuidDTO) (*domain.User, error)
	ExistsByEmail(dto ExistsByEmailDTO) bool
}
