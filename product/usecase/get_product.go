package usecase

import (
	"context"
	"wire-clean-architecture/product/model"
	"wire-clean-architecture/product/repository"
)

type GetProductUsecase struct {
	ProductRepository repository.IProductRepository
}

func ProvideGetProductUsecase(repository repository.IProductRepository) GetProductUsecase {
	return GetProductUsecase{ProductRepository: repository}
}

func (h GetProductUsecase) Execute(ctx context.Context, id uint) (tr model.Product) {
	return h.ProductRepository.FindByID(id)
}
