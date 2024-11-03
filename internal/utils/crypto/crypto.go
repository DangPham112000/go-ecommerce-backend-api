package crypto

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
)

func GetHash(key string) string {
	hash := sha256.New()
	hash.Write([]byte(key))
	hashBytes := hash.Sum(nil)

	return hex.EncodeToString(hashBytes)
}

func GenerateSalt(length int) (string, error) {
	salt := make([]byte, length)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(salt), nil
}

func HashPassword(password string, salt string) string {
	passSalt := password + salt
	hashPass := sha256.Sum256(([]byte(passSalt)))
	return hex.EncodeToString(hashPass[:])
}

func MatchingPassword(storeHash string, password string, salt string) bool {
	hashedPass := HashPassword(password, salt)
	return hashedPass == storeHash
}
