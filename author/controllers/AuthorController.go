package controllers

import (
	"go-jwt/author/services"
	"go-jwt/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//GetAuthor ... Get all author
func GetListAuthor(c *gin.Context) {
	data, err := services.ListAuthor()
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		c.JSON(200, gin.H{
			"status": "success",
			"data":   data,
		})
	}
}

//GetAuthorByID ... Get all author by id
func GetAuthorByID(c *gin.Context) {
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)
	data, err := services.DetailAuthor(int32(idInt))
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(200, gin.H{
			"status": "success",
			"data":   data,
		})
	}
}

//CreateAuthor ... Create author
func CreateAuthor(c *gin.Context) {
	var dataAuthor models.Author
	c.BindJSON(&dataAuthor)
	createdData, err := services.CreateAuthor(dataAuthor)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		c.JSON(200, gin.H{
			"status": "success",
			"data":   createdData,
		})
	}
}

//UpdateAuthor ... Update author
func UpdateAuthor(c *gin.Context) {
	var dataAuthor models.Author
	c.BindJSON(&dataAuthor)
	updatedData, err := services.UpdateAuthor(dataAuthor)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(200, gin.H{
			"status": "success",
			"data":   updatedData,
		})
	}
}

//DeleteAuthor ... Delete author
func DeleteAuthor(c *gin.Context) {
	var dataAuthor models.Author
	c.BindJSON(&dataAuthor)
	deletedData, err := services.DeleteAuthor(dataAuthor)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(200, gin.H{
			"status": "success",
			"data":   deletedData,
		})
	}
}

//LoginAuthor ... Login author
func LoginAuthor(c *gin.Context) {
	var dataAuthor models.Author
	c.BindJSON(&dataAuthor)
	token, err := services.LoginAuthor(dataAuthor)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "failed",
			"data":   "invalid email / password",
		})
	} else {
		c.JSON(200, gin.H{
			"status": "success",
			"data":   map[string]interface{}{"token": token},
		})
	}
}
