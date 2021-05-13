# golang-wire-clean-architecture

run tests for all folders and subfolders: 
    go test -v ./...
build project: 
    run build
wire, create/update dependency injections: 
    go generate
create mock files: 
    mockgen -source=product/repository/product.go -destination=product/mock/product_mock_auto_generated.go