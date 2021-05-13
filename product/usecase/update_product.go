package usecase

import (
	"context"
	"wire-clean-architecture/product/model"
	"wire-clean-architecture/product/repository"
)

type UpdateProductUsecase struct {
	ProductRepository repository.IProductRepository
	GetProductUsecase GetProductUsecase
}

func ProvideUpdateProductUsecase(repository repository.IProductRepository, getProductUsecase GetProductUsecase) UpdateProductUsecase {
	return UpdateProductUsecase{ProductRepository: repository, GetProductUsecase: getProductUsecase}
}

func (h UpdateProductUsecase) Execute(ctx context.Context, id uint, productDTO model.ProductDTO) (result bool) {
	product := h.GetProductUsecase.Execute(ctx, id)
	if product == (model.Product{}) {
		return false
	}

	product.Code = productDTO.Code
	product.Price = productDTO.Price
	h.ProductRepository.Save(product)

	return true
}
