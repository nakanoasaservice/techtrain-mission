package helper

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/speps/go-hashids"
	"log"
)

var hashID *hashids.HashID

func Init() {
	hd := hashids.NewData()
	hd.Salt = "this is my salt"
	var err error
	hashID, err = hashids.NewWithData(hd)
	if err != nil {
		log.Fatal(err)
	}
}

func GenerateToken(userID uint) (string, error) {
	encodedID, err := EncodeID(userID)
	if err != nil {
		return "", err
	}

	claims := jwt.StandardClaims{
		Subject: encodedID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte("0e8iuGo13NoLRfqpbpyWDyfMAAyBiqbPBN1U/6jDfI7K/9xdx36zAw=="))
}

func DecodeToken(signedString string) (*jwt.Token, *jwt.StandardClaims, error) {
	claims := &jwt.StandardClaims{}
	token, err := jwt.ParseWithClaims(signedString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return 0, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("0e8iuGo13NoLRfqpbpyWDyfMAAyBiqbPBN1U/6jDfI7K/9xdx36zAw=="), nil
	})

	return token, claims, err
}

func GetUserIDFromToken(token string) (uint, error) {
	_, claims, err := DecodeToken(token)
	if err != nil {
		return 0, err
	}
	return DecodeID(claims.Subject)
}

func EncodeID(id uint) (string, error) {
	return hashID.Encode([]int{int(id)})
}

func DecodeID(encodedID string) (uint, error) {
	ids, err := hashID.DecodeWithError(encodedID)
	if err != nil {
		return 0, err
	}
	if len(ids) != 1 {
		return 0, fmt.Errorf("invalid id: %v", ids)
	}
	return uint(ids[0]), err
}
