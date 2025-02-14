package mock

import (
    "context"
    "github.com/stretchr/testify/mock"
    "Gomall/rpc_gen/kitex_gen/product"// 假设这是 product 包
)

// MockProductClient 用于模拟 ProductClient
type MockProductClient struct {
    mock.Mock
}

// GetProduct 是 MockProductClient 的模拟方法
func (m *MockProductClient) GetProduct(ctx context.Context, req *product.GetProductReq) (*product.GetProductResp, error) {
    args := m.Called(ctx, req)
    
    // 返回模拟的 GetProductResp，其中包含一个模拟的 Product
    return &product.GetProductResp{
        Product: &product.Product{
            Id:          123,
            Name:        "Test Product",
            Description: "This is a test product.",
            Picture:     "test_product_picture.jpg",
            Price:       100.0,
            Categories:  []string{"Category1", "Category2"},
        },
    }, args.Error(1)
}
