// Agrupación de rutas 
package routes

import (
	"go-microservices/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api")

	api.POST("/signup", controllers.Signup)
	api.POST("/organization", controllers.CreateOrganization)
}
