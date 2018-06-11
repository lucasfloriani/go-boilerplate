package router

import (
	"go-boilerplate/daos"
	"go-boilerplate/handlers"
	"go-boilerplate/middleware"
	"golang-restful-starter-kit/services"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// Setup creates a the gin engine struct and add custom options to it
func Setup(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.Use(middleware.DBtoContext(db))
	InitializeRoutes(r)
	return r
}

// InitializeRoutes add routes to the system
func InitializeRoutes(r *gin.Engine) {
	artistDAO := daos.NewArtistDAO()
	handlers.ServeArtistResource(r, services.NewArtistService(artistDAO))
}
