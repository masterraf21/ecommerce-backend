package auth

import "golang.org/x/crypto/bcrypt"

// GeneratePassword will hash plaintext
func GeneratePassword(plain string) (password string, err error) {
	pwd := []byte(plain)
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		return
	}

	password = string(hash)

	return
}

// ComparePassword will compare password
func ComparePassword(hashed string, plain string) (bool, error) {
	byteHash := []byte(hashed)
	bytePlain := []byte(plain)

	err := bcrypt.CompareHashAndPassword(byteHash, bytePlain)
	if err != nil {
		return false, err
	}

	return true, nil
}
