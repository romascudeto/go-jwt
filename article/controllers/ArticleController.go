package controllers

import (
	"go-jwt/article/models"
	"go-jwt/article/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//GetArticle ... Get all article
func GetListArticle(c *gin.Context) {
	data, err := services.ListArticle()
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		c.JSON(200, gin.H{
			"status": "success",
			"data":   data,
		})
	}
}

//GetArticleByID ... Get all article by id
func GetArticleByID(c *gin.Context) {
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)
	data, err := services.DetailArticle(int32(idInt))
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(200, gin.H{
			"status": "success",
			"data":   data,
		})
	}
}

//CreateArticle ... Create article
func CreateArticle(c *gin.Context) {
	var dataArticle models.Article
	c.BindJSON(&dataArticle)
	createdData, err := services.CreateArticle(dataArticle)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		c.JSON(200, gin.H{
			"status": "success",
			"data":   createdData,
		})
	}
}

//UpdateArticle ... Update article
func UpdateArticle(c *gin.Context) {
	var dataArticle models.Article
	c.BindJSON(&dataArticle)
	updatedData, err := services.UpdateArticle(dataArticle)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(200, gin.H{
			"status": "success",
			"data":   updatedData,
		})
	}
}

//DeleteArticle ... Delete article
func DeleteArticle(c *gin.Context) {
	var dataArticle models.Article
	c.BindJSON(&dataArticle)
	deletedData, err := services.DeleteArticle(dataArticle)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(200, gin.H{
			"status": "success",
			"data":   deletedData,
		})
	}
}
