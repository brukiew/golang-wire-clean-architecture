package controller

import (
	"log"
	"net/http"
	"strconv"

	"wire-clean-architecture/product/helper"
	"wire-clean-architecture/product/model"
	"wire-clean-architecture/product/usecase"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	GetAllProductsUsecase usecase.GetAllProductsUsecase
	GetProductUsecase     usecase.GetProductUsecase
	DeleteProductUsecase  usecase.DeleteProductUsecase
	UpdateProductUsecase  usecase.UpdateProductUsecase
	CreateProductUsecase  usecase.CreateProductUsecase
}

func ProvideProductController(
	getAllProductsUsecase usecase.GetAllProductsUsecase,
	getProductUsecase usecase.GetProductUsecase,
	deleteProductUsecase usecase.DeleteProductUsecase,
	updateProductUsecase usecase.UpdateProductUsecase,
	createProductUsecase usecase.CreateProductUsecase,
) ProductController {
	return ProductController{
		GetAllProductsUsecase: getAllProductsUsecase,
		GetProductUsecase:     getProductUsecase,
		DeleteProductUsecase:  deleteProductUsecase,
		UpdateProductUsecase:  updateProductUsecase,
		CreateProductUsecase:  createProductUsecase,
	}
}

func (p *ProductController) FindAll(ctx *gin.Context) {
	products := p.GetAllProductsUsecase.Execute(ctx)
	ctx.JSON(http.StatusOK, gin.H{"products": helper.ToProductDTOs(products)})
}

func (p *ProductController) FindByID(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	product := p.GetProductUsecase.Execute(ctx, uint(id))
	ctx.JSON(http.StatusOK, gin.H{"product": helper.ToProductDTO(product)})
}

func (p *ProductController) Create(ctx *gin.Context) {
	var productDTO model.ProductDTO
	err := ctx.BindJSON(&productDTO)
	if err != nil {
		log.Fatalln(err)
		ctx.Status(http.StatusBadRequest)
		return
	}
	createdProduct := p.CreateProductUsecase.Execute(ctx, helper.ToProduct(productDTO))
	ctx.JSON(http.StatusOK, gin.H{"product": helper.ToProductDTO(createdProduct)})
}

func (p *ProductController) Update(ctx *gin.Context) {
	var productDTO model.ProductDTO
	err := ctx.BindJSON(&productDTO)
	if err != nil {
		log.Fatalln(err)
		ctx.Status(http.StatusBadRequest)
		return
	}

	id, _ := strconv.Atoi(ctx.Param("id"))
	productUpdated := p.UpdateProductUsecase.Execute(ctx, uint(id), productDTO)
	if !productUpdated {
		ctx.Status(http.StatusBadRequest)
		return
	}
	ctx.Status(http.StatusOK)
}

func (p *ProductController) Delete(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	productDeleted := p.DeleteProductUsecase.Execute(ctx, uint(id))
	if !productDeleted {
		ctx.Status(http.StatusBadRequest)
		return
	}
	ctx.Status(http.StatusOK)
}
