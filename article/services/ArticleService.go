package services

import (
	"go-jwt/article/models"
	"go-jwt/article/repositories"
)

func ListArticle() (dataRet []models.Article, err error) {
	err = repositories.Fetch(&dataRet)
	return dataRet, err
}

func DetailArticle(id int32) (dataRet models.Article, err error) {
	err = repositories.Detail(&dataRet, id)
	return dataRet, err
}

func CreateArticle(article models.Article) (dataRet models.Article, err error) {
	err = repositories.Create(&article)
	return article, err
}

func UpdateArticle(article models.Article) (dataRet models.Article, err error) {
	err = repositories.Update(&article)
	return article, err
}

func DeleteArticle(article models.Article) (dataRet models.Article, err error) {
	err = repositories.Delete(&article)
	return article, err
}
