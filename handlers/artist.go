package handlers

import (
	"go-boilerplate/models"
	"golang-restful-starter-kit/app"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type (
	// artistService specifies the interface for the artist service needed by artistResource.
	artistService interface {
		Get(rs app.RequestScope, id int) (*models.Artist, error)
		Query(rs app.RequestScope, offset, limit int) ([]models.Artist, error)
		Count(rs app.RequestScope) (int, error)
		Create(rs app.RequestScope, model *models.Artist) (*models.Artist, error)
		Update(rs app.RequestScope, id int, model *models.Artist) (*models.Artist, error)
		Delete(rs app.RequestScope, id int) (*models.Artist, error)
	}

	// artistResource defines the handlers for the CRUD APIs.
	artistResource struct {
		service artistService
	}
)

// ServeArtist sets up the routing of artist endpoints and the corresponding handlers.
func ServeArtistResource(r *gin.Engine, service artistService) {
	at := &artistResource{service}
	r.GET("/artists/:artistID", at.get)
	r.GET("/artists", at.query)
	r.POST("/artists", at.create)
	r.PUT("/artists/:artistID", at.update)
	r.DELETE("/artists/:artistID", at.delete)
}

func (r *artistResource) get(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID do artista é inválido"})
		return
	}

	response, err := r.service.Get(c, id)
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
	paginatedList := getPaginatedListFromRequest(c, count)
	items, err := r.service.Query(app.GetRequestScope(c), paginatedList.Offset(), paginatedList.Limit())
	if err != nil {
		return err
	}
	paginatedList.Items = items
	return c.Write(paginatedList)
}

func (r *artistResource) create(c *gin.Context) {
	var model models.Artist
	if err := c.Read(&model); err != nil {
		return err
	}
	response, err := r.service.Create(app.GetRequestScope(c), &model)
	if err != nil {
		return err
	}

	return c.Write(response)
}

func (r *artistResource) update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	rs := app.GetRequestScope(c)

	model, err := r.service.Get(rs, id)
	if err != nil {
		return err
	}

	if err := c.Read(model); err != nil {
		return err
	}

	response, err := r.service.Update(rs, id, model)
	if err != nil {
		return err
	}

	return c.Write(response)
}

func (r *artistResource) delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	response, err := r.service.Delete(app.GetRequestScope(c), id)
	if err != nil {
		return err
	}

	return c.Write(response)
}
