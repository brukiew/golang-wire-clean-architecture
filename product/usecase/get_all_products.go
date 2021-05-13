package usecase

import (
	"context"
	"wire-clean-architecture/product/model"
	"wire-clean-architecture/product/repository"
)

type GetAllProductsUsecase struct {
	ProductRepository repository.IProductRepository
}

func ProvideGetAllProductsUsecase(repository repository.IProductRepository) GetAllProductsUsecase {
	return GetAllProductsUsecase{ProductRepository: repository}
}

func (h GetAllProductsUsecase) Execute(ctx context.Context) (tr []model.Product) {
	return h.ProductRepository.FindAll()
}
