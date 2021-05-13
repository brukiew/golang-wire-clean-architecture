package helper

import "wire-clean-architecture/product/model"

func ToProduct(productDTO model.ProductDTO) model.Product {
	return model.Product{Code: productDTO.Code, Price: productDTO.Price}
}

func ToProductDTO(product model.Product) model.ProductDTO {
	return model.ProductDTO{ID: product.ID, Code: product.Code, Price: product.Price}
}

func ToProductDTOs(products []model.Product) []model.ProductDTO {
	productdtos := make([]model.ProductDTO, len(products))

	for i, itm := range products {
		productdtos[i] = ToProductDTO(itm)
	}

	return productdtos
}
