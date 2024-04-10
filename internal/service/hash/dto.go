package hash

type HashDTO struct {
	Password string
}

type CheckDTO struct {
	Password string
	Hash     string
}

type TokenDTO struct {
	Uuid  string
	Email string
}

type AuthDTO struct {
	Token string
}
