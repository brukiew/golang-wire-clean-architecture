package repository

import (
	model "wire-clean-architecture/product/model"

	"github.com/jinzhu/gorm"
)

type IProductRepository interface {
	FindAll() []model.Product
	FindByID(id uint) model.Product
	Save(product model.Product) model.Product
	Delete(product model.Product)
}

type ProductRepository struct {
	DB *gorm.DB
}

func ProvideProductRepostiory(db *gorm.DB) IProductRepository {
	instance := &ProductRepository{DB: db}
	return instance
}

func (p *ProductRepository) FindAll() []model.Product {
	var products []model.Product
	p.DB.Find(&products)

	return products
}

func (p *ProductRepository) FindByID(id uint) model.Product {
	var product model.Product
	p.DB.First(&product, id)

	return product
}

func (p *ProductRepository) Save(product model.Product) model.Product {
	p.DB.Save(&product)

	return product
}

func (p *ProductRepository) Delete(product model.Product) {
	p.DB.Delete(&product)
}
