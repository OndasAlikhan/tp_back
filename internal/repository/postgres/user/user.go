package user

import (
	"database/sql"
	"errors"
	"log/slog"
	"time"

	"tp_back/internal/domain"
	"tp_back/internal/exception"
	"tp_back/internal/lib/sl"
	"tp_back/internal/resource/postgres"
	"tp_back/internal/service/user"
)

type Repository struct {
	log *slog.Logger
	res *postgres.Res
}

func New(log *slog.Logger, res *postgres.Res) *Repository {
	return &Repository{log: log, res: res}
}

func (r *Repository) Save(dto user.SaveDTO) (string, error) {
	const op = "repository.postgres.user.Save"

	log := r.log.With(
		slog.String("op", op),
		slog.Any("dto", dto),
	)

	var uuid string

	log.Info("saving user")
	query := `insert into users(email, password, created_at, updated_at) values ($1, $2, $3, $4) RETURNING uuid`
	err := r.res.DB.QueryRow(query, dto.Email, dto.PasswordHash, time.Now(), time.Now()).Scan(&uuid)
	if err != nil {
		log.Error("failed to save user: ", sl.Error(err))

		return "", err
	}
	log.Info("saved user uuid", slog.String("uuid", uuid))

	return uuid, nil
}

func (r *Repository) FindByEmail(dto user.FindByEmailDTO) (*domain.User, error) {
	const op = "repository.postgres.user.FindByEmail"

	log := r.log.With(
		slog.String("op", op),
		slog.Any("dto", dto),
	)

	var u domain.User

	log.Info("finding user by email")
	query := `select uuid, email, password, created_at, updated_at from users where email = $1`
	if err := r.res.DB.QueryRow(query, dto.Email).Scan(&u.Uuid, &u.Email, &u.Password, &u.CreatedAt, &u.UpdatedAt); err != nil {
		log.Error("failed to found user by email: ", sl.Error(err))
		if errors.Is(err, sql.ErrNoRows) {
			return nil, exception.ErrNotFound
		}
		return nil, err
	}
	log.Info("found user by email", slog.Any("user", u))

	return &u, nil
}

func (r *Repository) FindByUuid(dto user.FindByUuidDTO) (*domain.User, error) {
	const op = "repository.postgres.user.FindByUuid"

	log := r.log.With(
		slog.String("op", op),
		slog.Any("dto", dto),
	)

	var u domain.User

	log.Info("finding user by uuid")
	query := `select uuid, email, password, created_at, updated_at from users where uuid = $1`
	if err := r.res.DB.QueryRow(query, dto.Uuid).Scan(&u.Uuid, &u.Email, &u.Password, &u.CreatedAt, &u.UpdatedAt); err != nil {
		log.Error("failed to found user by uuid: ", sl.Error(err))
		if errors.Is(err, sql.ErrNoRows) {
			return nil, exception.ErrNotFound
		}
		return nil, err
	}
	log.Info("found user by uuid", slog.Any("user", u))

	return &u, nil
}

func (r *Repository) ExistsByEmail(dto user.ExistsByEmailDTO) bool {
	const op = "repository.postgres.user.ExistsByEmail"

	log := r.log.With(
		slog.String("op", op),
		slog.Any("dto", dto),
	)

	var exists bool

	log.Info("checking exists user by email")
	query := `select exists(select 1 from users where email = $1)`
	if err := r.res.DB.QueryRow(query, dto.Email).Scan(&exists); err != nil {
		log.Error("failed to check user by email: ", sl.Error(err))
		return false
	}
	log.Info("checked exists user by email", slog.Bool("exists", exists))

	return exists
}
