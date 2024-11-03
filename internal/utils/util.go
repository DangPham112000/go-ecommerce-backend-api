package utils

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

func GetUserKey(hashKey string) string {
	return fmt.Sprintf("u:%s:otp", hashKey)
}

func GenerateCliTokenUUID(userId int) string {
	newUUID := uuid.New()
	// Convert UUID to string and remove `-` symbol
	uuidString := strings.ReplaceAll(newUUID.String(), "-", "")
	return strconv.Itoa(userId) + "clitoken" + uuidString
}
