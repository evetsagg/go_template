package database

import (
	"go_template/src/model"

	"gorm.io/gorm"
)

type ProductDao struct {
	DB     *gorm.DB
	logger LoggingI
}

func NewProductDao(logger LoggingI, DB *gorm.DB) *ProductDao {
	return &ProductDao{DB: DB, logger: logger}
}

// https://gorm.io/docs/create.html
func (p *ProductDao) Create(product *model.Product) *model.Product {
	result := p.DB.Create(&product)
	if result.Error != nil {
		p.logger.Error(result.Error)
	}
	return product
}
func (p *ProductDao) Update(product *model.Product) *model.Product {
	result := p.DB.Model(&product).Updates(&product)
	if result.Error != nil {
		p.logger.Error(result.Error)
	}
	return product
}
func (p *ProductDao) Delete(product *model.Product) {
	result := p.DB.Delete(&product)
	if result.Error != nil {
		p.logger.Error(result.Error)
	}
}
func (p *ProductDao) FindAll() *[]model.Product {
	var products []model.Product
	result := p.DB.Find(&products)
	if result.Error != nil {
		p.logger.Error(result.Error)
	}
	return &products
}
func (p *ProductDao) FindById(product *model.Product) *model.Product {
	result := p.DB.First(&product)
	if result.Error != nil {
		p.logger.Error(result.Error)
	}
	return product
}
