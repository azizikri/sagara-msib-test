package repositories

import (
	"errors"
	"sagara-msib-test/models"

	"gorm.io/gorm"
)

type ClothRepository struct {
	DB *gorm.DB
}

func NewClothRepository(dB *gorm.DB) models.ClothRepository {
	return &ClothRepository{DB: dB}
}

func (r *ClothRepository) CreateCloth(cloth models.Cloth) error {
	var existingCloth models.Cloth

	result := r.DB.Where("color = ? AND size = ?", cloth.Color, cloth.Size).First(&existingCloth)
	if result.Error == nil {
		return errors.New("cloth with the same color and size already exists")
	} else if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return result.Error
	}

	// Create the new cloth record
	return r.DB.Create(&cloth).Error
}

func (r *ClothRepository) GetCloth(id uint) (models.Cloth, error) {
	var cloth models.Cloth
	err := r.DB.First(&cloth, id).Error
	return cloth, err
}

func (r *ClothRepository) GetClothes(color, size string, offset, limit int) ([]models.Cloth, error) {
	var clothes []models.Cloth
	query := r.DB

	if color != "" {
		query = query.Where("color = ?", color)
	}

	if size != "" {
		query = query.Where("size = ?", size)
	}

	err := query.Offset(offset).Limit(limit).Order("updated_at desc, created_at desc").Find(&clothes).Error
	return clothes, err
}

func (r *ClothRepository) UpdateCloth(cloth models.Cloth) error {
	var duplicateCloth models.Cloth

	result := r.DB.Where("color = ? AND size = ? AND id != ?", cloth.Color, cloth.Size, cloth.ID).First(&duplicateCloth)
	if result.Error == nil {
		return errors.New("cloth with the same color and size already exists")
	} else if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return result.Error
	}

	return r.DB.Save(&cloth).Error
}

func (r *ClothRepository) DeleteCloth(id uint) error {
	return r.DB.Delete(&models.Cloth{}, id).Error
}

func (r *ClothRepository) FindLowStockClothes(color, size string) ([]models.Cloth, error) {
	var clothes []models.Cloth
	query := r.DB

	if color != "" {
		query = query.Where("color = ?", color)
	}

	if size != "" {
		query = query.Where("size = ?", size)
	}

	err := query.Where("stock < ?", 5).Order("created_at desc").Find(&clothes).Error
	return clothes, err
}

func (r *ClothRepository) FindOutOfStockClothes(color, size string) ([]models.Cloth, error) {
	var clothes []models.Cloth
	query := r.DB

	if color != "" {
		query = query.Where("color = ?", color)
	}

	if size != "" {
		query = query.Where("size = ?", size)
	}

	err := query.Where(map[string]interface{}{"stock": 0}).Find(&clothes).Error
	return clothes, err
}
