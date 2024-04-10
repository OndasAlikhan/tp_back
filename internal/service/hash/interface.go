package hash

type Service interface {
	Hash(dto HashDTO) (string, error)
	Check(dto CheckDTO) bool
	Token(dto TokenDTO) (string, error)
}
