package author

import (
	"go-jwt/author/controllers"

	"github.com/gin-gonic/gin"
)

func Routes(route *gin.Engine) {
	apiAuthor := route.Group("/api")
	{
		apiAuthor.GET("/author", controllers.GetListAuthor)
		apiAuthor.GET("/author/:id", controllers.GetAuthorByID)
		apiAuthor.POST("/author", controllers.CreateAuthor)
		apiAuthor.PUT("/author", controllers.UpdateAuthor)
		apiAuthor.DELETE("/author", controllers.DeleteAuthor)
		apiAuthor.POST("/author/login", controllers.LoginAuthor)
	}
}
