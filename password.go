package tools

import "golang.org/x/crypto/bcrypt"

// HashAndSaltPassword hashes and salts a password
func HashAndSaltPassword(pwd string) (string, error) {
	password := []byte(pwd)
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	hashedPassword := string(hash)

	return hashedPassword, nil
}

// CompareHashAndPassword compares a plain password with an hashed one
func CompareHashAndPassword(hashedPassword string, plainPassword string) error {
	byteHash := []byte(hashedPassword)
	bytePlain := []byte(plainPassword)

	err := bcrypt.CompareHashAndPassword(byteHash, bytePlain)
	if err != nil {
		return err
	}

	return nil
}
