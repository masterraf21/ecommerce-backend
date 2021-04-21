package models

// AuthUsecase method for auth
type AuthUsecase interface {
	CreateToken(role string, id uint32) (string, error)
}
