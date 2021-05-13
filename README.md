
# golang-wire-clean-architecture
Base template for golang project using wire(dependency injection), gin and clean architecture.

Dependencies:
wire: https://github.com/google/wire
gin: https://github.com/gin-gonic/gin
gomock: https://github.com/golang/mock

## build project
go build

## wire, create/update dependency injections:
go generate

## create mock files
mockgen -source=product/repository/product.go -destination=product/mock/product_mock_auto_generated.go

## run tests for all folders and subfolders
go test -v ./...