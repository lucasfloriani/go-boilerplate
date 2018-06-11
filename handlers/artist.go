package handlers

import (
	"go-boilerplate/helper"
	"go-boilerplate/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type (
	// artistService specifies the interface for the artist service needed by artistResource.
	artistService interface {
		Get(c *gin.Context, id uint) (*models.Artist, error)
		Query(c *gin.Context, offset, limit int) ([]models.Artist, error)
		Count(c *gin.Context) (int, error)
		Create(c *gin.Context, model *models.Artist) (*models.Artist, error)
		Update(c *gin.Context, id uint, model *models.Artist) (*models.Artist, error)
		Delete(c *gin.Context, id uint) (*models.Artist, error)
	}

	// artistResource defines the handlers for the CRUD APIs.
	artistResource struct {
		service artistService
	}
)

// ServeArtistResource sets up the routing of artist endpoints and the corresponding handlers.
func ServeArtistResource(r *gin.RouterGroup, service artistService) {
	at := &artistResource{service}
	artistGroup := r.Group("/artists")
	{
		artistGroup.GET("/:artistID", at.get)
		artistGroup.GET("/", at.query)
		artistGroup.POST("/", at.create)
		artistGroup.PUT("/:artistID", at.update)
		artistGroup.DELETE("/:artistID", at.delete)
	}
}

func (r *artistResource) get(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("artistID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID do artista é inválido"})
		return
	}

	response, err := r.service.Get(c, uint(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusFound, gin.H{
		"id":      response.ID,
		"error":   "",
		"message": response,
	})
}

func (r *artistResource) query(c *gin.Context) {
	count, err := r.service.Count(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	paginatedList := helper.GetPaginatedListFromRequest(c, count)
	items, err := r.service.Query(c, paginatedList.Offset(), paginatedList.Limit())
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	paginatedList.Items = items
	c.JSON(http.StatusFound, gin.H{
		"id":      "",
		"error":   "",
		"message": paginatedList,
	})
}

func (r *artistResource) create(c *gin.Context) {
	var model models.Artist
	if err := c.Bind(&model); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	response, err := r.service.Create(c, &model)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":      response.ID,
		"error":   "",
		"message": response,
	})
}

func (r *artistResource) update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("artistID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	model, err := r.service.Get(c, uint(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	if err := c.Bind(model); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	response, err := r.service.Update(c, uint(id), model)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":      response.ID,
		"error":   "",
		"message": response,
	})
}

func (r *artistResource) delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("artistID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	response, err := r.service.Delete(c, uint(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":      response.ID,
		"error":   "",
		"message": response,
	})
}
