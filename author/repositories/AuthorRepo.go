package repositories

import (
	"go-jwt/author/models"
	Config "go-jwt/config"

	_ "github.com/go-sql-driver/mysql"
)

func Fetch(authors *[]models.Author) (err error) {

	if err = Config.DB.Find(authors).Error; err != nil {
		return err
	}
	return nil
}

func Detail(author *models.Author, id int32) (err error) {

	if err = Config.DB.Where("id = ?", id).First(author).Error; err != nil {
		return err
	}
	return nil
}

func Create(author *models.Author) (err error) {

	if err = Config.DB.Create(author).Error; err != nil {
		return err
	}
	return nil
}

func Update(author *models.Author) (err error) {
	if err = Config.DB.Model(&author).Updates(author).Error; err != nil {
		return err
	}
	return nil
}

func Delete(author *models.Author) (err error) {
	if err = Config.DB.Where("id = ?", author.ID).First(author).Error; err != nil {
		return err
	}
	if err = Config.DB.Delete(author).Error; err != nil {
		return err
	}
	return nil
}

func DetailByEmail(author *models.Author, authorReq models.Author) (err error) {

	if err = Config.DB.Where("email = ?", authorReq.Email).First(author).Error; err != nil {
		return err
	}
	return nil
}
