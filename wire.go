//+build wireinject

package main

import (
	"wire-clean-architecture/product/controller"
	"wire-clean-architecture/product/repository"
	"wire-clean-architecture/product/usecase"

	"github.com/google/wire"
	"github.com/jinzhu/gorm"
)

func InitProductController(db *gorm.DB) controller.ProductController {
	wire.Build(
		repository.ProvideProductRepostiory,
		controller.ProvideProductController,
		usecase.ProvideGetAllProductsUsecase,
		usecase.ProvideGetProductUsecase,
		usecase.ProvideDeleteProductUsecase,
		usecase.ProvideUpdateProductUsecase,
		usecase.ProvideCreateProductUsecase,
	)

	return controller.ProductController{}
}
