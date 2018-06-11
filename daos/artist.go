package daos

import (
	dbpkg "go-boilerplate/db"
	"go-boilerplate/models"

	"github.com/gin-gonic/gin"
)

// ArtistDAO persists artist data in database
type ArtistDAO struct{}

// NewArtistDAO creates a new ArtistDAO
func NewArtistDAO() *ArtistDAO {
	return &ArtistDAO{}
}

// Get reads the artist with the specified ID from the database.
func (dao *ArtistDAO) Get(c *gin.Context, id uint) (*models.Artist, error) {
	artist := models.Artist{}
	db := dbpkg.Instance(c)
	err := db.First(&artist, id).Error
	return &artist, err
}

// Create saves a new artist record in the database.
// The Artist.Id field will be populated with an automatically generated ID upon successful saving.
func (dao *ArtistDAO) Create(c *gin.Context, artist *models.Artist) error {
	db := dbpkg.Instance(c)
	return db.Save(&artist).Error
}

// Update saves the changes to an artist in the database.
func (dao *ArtistDAO) Update(c *gin.Context, id uint, artist *models.Artist) error {
	if _, err := dao.Get(c, id); err != nil {
		return err
	}
	db := dbpkg.Instance(c)
	err := db.Save(&artist).Error
	return err
}

// Delete deletes an artist with the specified ID from the database.
func (dao *ArtistDAO) Delete(c *gin.Context, id uint) error {
	artist, err := dao.Get(c, id)
	if err != nil {
		return err
	}
	db := dbpkg.Instance(c)
	return db.Delete(&artist).Error
}

// Count returns the number of the artist records in the database.
func (dao *ArtistDAO) Count(c *gin.Context) (count int, err error) {
	db := dbpkg.Instance(c)
	err = db.Model(&models.Artist{}).Count(&count).Error
	return
}

// Query retrieves the artist records with the specified offset and limit from the database.
func (dao *ArtistDAO) Query(c *gin.Context, offset, limit int) ([]models.Artist, error) {
	artists := []models.Artist{}
	db := dbpkg.Instance(c)
	return artists, db.Offset(offset).Limit(limit).Order("id asc").Find(&artists).Error
}
