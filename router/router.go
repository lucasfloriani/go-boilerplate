package router

import (
	"go-boilerplate/daos"
	"go-boilerplate/handlers"
	"go-boilerplate/middleware"
	"go-boilerplate/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-contrib/location"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// Setup creates a the gin engine struct and add custom options to it
func Setup(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.Use(
		middleware.DBtoContext(db),
		cors.Default(),
		location.Default(),
		gzip.Gzip(gzip.BestCompression),
	)
	InitializeRoutes(r)
	return r
}

// InitializeRoutes add routes to the system
func InitializeRoutes(r *gin.Engine) {
	authMiddleware := middleware.GetJWTAuth()

	v1 := r.Group("/v1")
	v1.POST("/login", authMiddleware.LoginHandler)
	v1.Use(authMiddleware.MiddlewareFunc())

	artistDAO := daos.NewArtistDAO()
	handlers.ServeArtistResource(v1, services.NewArtistService(artistDAO))
}
