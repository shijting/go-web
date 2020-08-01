package common

import (
	"crypto/md5"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

var salt = "slkfjslfj123--=40+-88"

func Md5(v string) string {
	h := md5.New()
	h.Write([]byte(salt))
	return fmt.Sprintf("%s", h.Sum([]byte(v)))
}
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GetHashingCost(hashedPassword []byte) (cost int, err error) {
	cost, err = bcrypt.Cost(hashedPassword) // 为了简单忽略错误处理
	return
}
