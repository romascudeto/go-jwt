package article

import (
	"go-jwt/article/controllers"
	"go-jwt/middleware"

	"github.com/gin-gonic/gin"
)

func Routes(route *gin.Engine) {
	apiArticle := route.Group("/api").Use(middleware.Auth)
	{
		apiArticle.GET("/article", controllers.GetListArticle)
		apiArticle.GET("/article/:id", controllers.GetArticleByID)
		apiArticle.POST("/article", controllers.CreateArticle)
		apiArticle.PUT("/article", controllers.UpdateArticle)
		apiArticle.DELETE("/article", controllers.DeleteArticle)
	}
}
