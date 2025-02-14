package service

import (
	"Gomall/app/cart/biz/mock"
	cart "Gomall/rpc_gen/kitex_gen/cart"
	"Gomall/rpc_gen/kitex_gen/product"
	"context"
	"testing"
	mo "github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/assert"
)
func TestAddItem_Run(t *testing.T) {
    // 创建一个 mock 的 ProductClient
    mockProductClient := new(mock.MockProductClient)

    // 模拟 GetProduct 方法的返回值
    mockProductClient.On("GetProduct", mo.Anything, &product.GetProductReq{Id: 123}).Return(&product.GetProductResp{
        Product: &product.Product{
            Id:          123,
            Name:        "Test Product",
            Description: "This is a test product.",
            Picture:     "test_product_picture.jpg",
            Price:       100.0,
            Categories:  []string{"Category1", "Category2"},
        },
    }, nil)

	// 创建一个新的 AddItemService 实例
	addItemService := NewAddItemService(context.Background())

	// 模拟请求
	req := &cart.AddItemReq{
		UserId: 1,
		Item: &cart.CartItem{
			ProductId: 123,
			Quantity:  2,
		},
	}

	// 执行 Run 方法
	resp, err := addItemService.Run(req)

	// 验证没有错误，并且返回的响应是我们预期的
	assert.Nil(t, err)
	assert.NotNil(t, resp)

    // 确保 mock 的方法被调用
    mockProductClient.AssertExpectations(t)
}
// func TestAddItem_Run(t *testing.T) {
// 	ctx := context.Background()
// 	s := NewAddItemService(ctx)
// 	// init req and assert value

// 	req := &cart.AddItemReq{}
// 	resp, err := s.Run(req)
// 	t.Logf("err: %v", err)
// 	t.Logf("resp: %v", resp)

// 	// todo: edit your unit test

// }

