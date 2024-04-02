package sdk

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

// ParseInt convert string to int
func ParseInt(text string, defaultValue int) int {
	if text == "" {
		return defaultValue
	}

	num, err := strconv.Atoi(text)
	if err != nil {
		return defaultValue
	}
	return num
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func VerifyPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func HashDevice(ipAddress, userAgent string) string {
	hashKey := fmt.Sprintf("%s-%s", userAgent, ipAddress)

	hasher := md5.New()
	hasher.Write([]byte(hashKey))
	return hex.EncodeToString(hasher.Sum(nil))
}

func HashKey(slice []string) string {
	// hashKey := fmt.Sprintf("%s-%s", userAgent, ipAddress)
	hashKey := strings.Join(slice, "-")

	hasher := md5.New()
	hasher.Write([]byte(hashKey))
	return hex.EncodeToString(hasher.Sum(nil))
}
