package repositories

import (
	"go-jwt/article/models"
	Config "go-jwt/config"

	_ "github.com/go-sql-driver/mysql"
)

func Fetch(articles *[]models.Article) (err error) {

	if err = Config.DB.Find(articles).Error; err != nil {
		return err
	}
	return nil
}

func Detail(article *models.Article, id int32) (err error) {

	if err = Config.DB.Where("id = ?", id).First(article).Error; err != nil {
		return err
	}
	return nil
}

func Create(article *models.Article) (err error) {

	if err = Config.DB.Create(article).Error; err != nil {
		return err
	}
	return nil
}

func Update(article *models.Article) (err error) {
	if err = Config.DB.Model(&article).Updates(article).Error; err != nil {
		return err
	}
	return nil
}

func Delete(article *models.Article) (err error) {
	if err = Config.DB.Where("id = ?", article.ID).First(article).Error; err != nil {
		return err
	}
	if err = Config.DB.Delete(article).Error; err != nil {
		return err
	}
	return nil
}
