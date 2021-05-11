package services

import (
	"errors"
	"go-jwt/author/helpers"
	"go-jwt/author/models"
	"go-jwt/author/repositories"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func ListAuthor() (dataRet []models.Author, err error) {
	err = repositories.Fetch(&dataRet)
	return dataRet, err
}

func DetailAuthor(id int32) (dataRet models.Author, err error) {
	err = repositories.Detail(&dataRet, id)
	return dataRet, err
}

func CreateAuthor(author models.Author) (dataRet models.Author, err error) {
	author.Password = helpers.HashPassword(author.Password)
	err = repositories.Create(&author)
	return author, err
}

func UpdateAuthor(author models.Author) (dataRet models.Author, err error) {
	author.Password = helpers.HashPassword(author.Password)
	err = repositories.Update(&author)
	return author, err
}

func DeleteAuthor(author models.Author) (dataRet models.Author, err error) {
	err = repositories.Delete(&author)
	return author, err
}

func LoginAuthor(author models.Author) (token string, err error) {
	var dataRet models.Author
	err = repositories.DetailByEmail(&dataRet, author)
	if err != nil {
		return "", err
	}
	match := helpers.CheckPasswordHash(author.Password, dataRet.Password)
	if !match {
		err = errors.New("invalid email / password")
		return "", err
	}
	tokenClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":        dataRet.ID,
		"email":     dataRet.Email,
		"nbf":       time.Now().Unix(),
		"expiredAt": time.Now().AddDate(0, 0, 1).Unix(),
	})
	mySigningKey := []byte("secret")

	tokenString, _ := tokenClaim.SignedString(mySigningKey)
	return tokenString, nil
}
