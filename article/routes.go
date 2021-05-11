package article

import (
	"go-jwt/article/controllers"

	"github.com/gin-gonic/gin"
)

func Routes(route *gin.Engine) {
	apiArticle := route.Group("/api")
	{
		apiArticle.GET("/article", controllers.GetListArticle)
		apiArticle.GET("/article/:id", controllers.GetArticleByID)
		apiArticle.POST("/article", controllers.CreateArticle)
		apiArticle.PUT("/article", controllers.UpdateArticle)
		apiArticle.DELETE("/article", controllers.DeleteArticle)
	}
}
