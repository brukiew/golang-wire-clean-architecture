package usecase

import (
	"context"
	"wire-clean-architecture/product/model"
	"wire-clean-architecture/product/repository"
)

type DeleteProductUsecase struct {
	ProductRepository repository.IProductRepository
	GetProductUsecase GetProductUsecase
}

func ProvideDeleteProductUsecase(repository repository.IProductRepository, getProductUsecase GetProductUsecase) DeleteProductUsecase {
	return DeleteProductUsecase{ProductRepository: repository, GetProductUsecase: getProductUsecase}
}

func (h DeleteProductUsecase) Execute(ctx context.Context, id uint) (result bool) {
	product := h.GetProductUsecase.Execute(ctx, id)
	if product == (model.Product{}) {
		return false
	}

	h.ProductRepository.Delete(product)
	return true
}
