package usecase

import (
	"context"
	"testing"
	mock_repository "wire-clean-architecture/mock"
	"wire-clean-architecture/product/model"

	"github.com/golang/mock/gomock"
)

func TestGetProductUsecase(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	resultProduct := model.Product{Code: "TEST_PRODUCT"}

	productRepositoryMock := mock_repository.NewMockIProductRepository(ctrl)
	productRepositoryMock.EXPECT().FindByID(gomock.Eq(uint(10))).Return(resultProduct)
	usecase := ProvideGetProductUsecase(productRepositoryMock)
	result := usecase.Execute(context.TODO(), 10)

	if result != resultProduct {
		t.Errorf("Product doesn't match expected result")
	}
}
