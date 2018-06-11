package services

import (
	"go-boilerplate/models"

	"github.com/gin-gonic/gin"
)

// artistDAO specifies the interface of the artist DAO needed by ArtistService.
type artistDAO interface {
	// Get returns the artist with the specified artist ID.
	Get(c *gin.Context, id int) (*models.Artist, error)
	// Count returns the number of artists.
	Count(c *gin.Context) (int, error)
	// Query returns the list of artists with the given offset and limit.
	Query(c *gin.Context, offset, limit int) ([]models.Artist, error)
	// Create saves a new artist in the storage.
	Create(c *gin.Context, artist *models.Artist) error
	// Update updates the artist with given ID in the storage.
	Update(c *gin.Context, id int, artist *models.Artist) error
	// Delete removes the artist with given ID from the storage.
	Delete(c *gin.Context, id int) error
}

// ArtistService provides services related with artists.
type ArtistService struct {
	dao artistDAO
}

// NewArtistService creates a new ArtistService with the given artist DAO.
func NewArtistService(dao artistDAO) *ArtistService {
	return &ArtistService{dao}
}

// Get returns the artist with the specified the artist ID.
func (s *ArtistService) Get(c *gin.Context, id int) (*models.Artist, error) {
	return s.dao.Get(c, id)
}

// Create creates a new artist.
func (s *ArtistService) Create(c *gin.Context, model *models.Artist) (*models.Artist, error) {
	if err := model.Validate(); err != nil {
		return nil, err
	}
	if err := s.dao.Create(c, model); err != nil {
		return nil, err
	}
	return s.dao.Get(c, model.ID)
}

// Update updates the artist with the specified ID.
func (s *ArtistService) Update(c *gin.Context, id int, model *models.Artist) (*models.Artist, error) {
	if err := model.Validate(); err != nil {
		return nil, err
	}
	if err := s.dao.Update(c, id, model); err != nil {
		return nil, err
	}
	return s.dao.Get(c, id)
}

// Delete deletes the artist with the specified ID.
func (s *ArtistService) Delete(c *gin.Context, id int) (*models.Artist, error) {
	artist, err := s.dao.Get(c, id)
	if err != nil {
		return nil, err
	}
	err = s.dao.Delete(c, id)
	return artist, err
}

// Count returns the number of artists.
func (s *ArtistService) Count(c *gin.Context) (int, error) {
	return s.dao.Count(c)
}

// Query returns the artists with the specified offset and limit.
func (s *ArtistService) Query(c *gin.Context, offset, limit int) ([]models.Artist, error) {
	return s.dao.Query(c, offset, limit)
}
