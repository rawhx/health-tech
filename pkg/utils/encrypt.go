package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"os"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Encrypt(c *gin.Context, plaintext string) (string, bool) {
	key := os.Getenv("ENCRYPTION_KEY")
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		ResponseError(c, 500, "Gagal membuat cipher", err)
		return "", false
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		ResponseError(c, 500, "Gagal membuat mode GCM", err)
		return "", false
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		ResponseError(c, 500, "Gagal membuat nonce", err)
		return "", false
	}

	ciphertext := gcm.Seal(nonce, nonce, []byte(plaintext), nil)
	return base64.StdEncoding.EncodeToString(ciphertext), true
}

func Decrypt(c *gin.Context, encodedCipher string) (string, bool) {
	key := os.Getenv("ENCRYPTION_KEY")
	data, err := base64.StdEncoding.DecodeString(encodedCipher)
	if err != nil {
		ResponseError(c, 400, "Gagal decode base64", err)
		return "", false
	}

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		ResponseError(c, 500, "Gagal membuat cipher", err)
		return "", false
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		ResponseError(c, 500, "Gagal membuat mode GCM", err)
		return "", false
	}

	if len(data) < gcm.NonceSize() {
		ResponseError(c, 400, "Ciphertext tidak valid", fmt.Errorf("ukuran ciphertext kurang dari nonce"))
		return "", false
	}

	nonce := data[:gcm.NonceSize()]
	ciphertext := data[gcm.NonceSize():]

	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		ResponseError(c, 400, "Gagal mendekripsi data", err)
		return "", false
	}

	return string(plaintext), true
}

func HashPassword(password string) (string, bool) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", false
	}
	return string(hashed), true
}

func CheckPassword(password, hashed string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
	if err != nil {
		return false
	}
	return true
}