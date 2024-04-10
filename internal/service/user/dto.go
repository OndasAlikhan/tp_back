package user

type RegisterDTO struct {
	Email    string
	Password string
}

type LoginDTO struct {
	Email    string
	Password string
}

type SaveDTO struct {
	Email        string
	PasswordHash string
}

type FindByEmailDTO struct {
	Email string
}

type FindByUuidDTO struct {
	Uuid string
}

type ExistsByEmailDTO struct {
	Email string
}
