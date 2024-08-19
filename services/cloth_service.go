package services

import (
	"sagara-msib-test/models"
)

type ClothService struct {
	repo models.ClothRepository
}

func NewClothService(repo models.ClothRepository) *ClothService {
	return &ClothService{repo}
}

func (s *ClothService) CreateCloth(cloth models.Cloth) error {
	return s.repo.CreateCloth(cloth)
}

func (s *ClothService) GetCloth(clothID uint) (models.Cloth, error) {
	return s.repo.GetCloth(clothID)
}

func (s *ClothService) GetClothes(color, size string, page, pageSize int) ([]models.Cloth, error) {
	offset := (page - 1) * pageSize
	return s.repo.GetClothes(color, size, offset, pageSize)
}

func (s *ClothService) UpdateCloth(cloth models.Cloth) error {
	return s.repo.UpdateCloth(cloth)
}

func (s *ClothService) DeleteCloth(clothID uint) error {
	return s.repo.DeleteCloth(clothID)
}

func (s *ClothService) FindOutOfStockClothes(color, size string) ([]models.Cloth, error) {
	return s.repo.FindOutOfStockClothes(color, size)
}

func (s *ClothService) FindLowStockClothes(color, size string) ([]models.Cloth, error) {
	return s.repo.FindLowStockClothes(color, size)
}


