package usecase

import (
	"context"
	"wire-clean-architecture/product/model"
	"wire-clean-architecture/product/repository"
)

type CreateProductUsecase struct {
	ProductRepository repository.IProductRepository
}

func ProvideCreateProductUsecase(repository repository.IProductRepository, getProductUsecase GetProductUsecase) CreateProductUsecase {
	return CreateProductUsecase{ProductRepository: repository}
}

func (h CreateProductUsecase) Execute(ctx context.Context, product model.Product) (result model.Product) {

	createdProduct := h.ProductRepository.Save(product)

	return createdProduct
}
