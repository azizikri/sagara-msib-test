package models

import (
	"gorm.io/gorm"
)

type Cloth struct {
	gorm.Model
	Color string  `gorm:"type:varchar(100);not null" json:"color" binding:"required,min=1,max=100"`
	Size  string  `gorm:"type:varchar(1); not null" json:"size" binding:"required,oneof='S' 'M' 'L'"`
	Price float64 `gorm:"type:numeric(12,2); not null" json:"price" binding:"required,gt=0"`
	Stock uint    `gorm:"type:uint;" json:"stock"`
}

type ClothRepository interface {
	CreateCloth(cloth Cloth) error
	GetCloth(clothID uint) (Cloth, error)
	GetClothes(color, size string, offset, limit int) ([]Cloth, error)
	UpdateCloth(cloth Cloth) error
	DeleteCloth(clothID uint) error
	FindOutOfStockClothes(color, size string) ([]Cloth, error)
	FindLowStockClothes(color, size string) ([]Cloth, error)
}
