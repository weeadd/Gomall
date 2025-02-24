package service

import (
	order "Gomall/rpc_gen/kitex_gen/order"
	"context"
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"Gomall/app/order/biz/dal/mysql"
	"Gomall/app/order/biz/model"
)

// 初始化测试数据库
func setupTestDB(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}
	// 迁移表结构
	db.AutoMigrate(&model.Order{}, &model.OrderItem{})
	mysql.DB = db
}

func TestPlaceOrderService_Run_NormalCase(t *testing.T) {
	setupTestDB(t)

	ctx := context.Background()
	service := NewPlaceOrderService(ctx)

	// 构造有效请求
	req := &order.PlaceOrderReq{
		UserId:       1001,
		UserCurrency: "USD",
		Email:        "test@example.com",
		OrderItems:   nil,
		Address: &order.Address{
			Country:       "US",
			State:         "CA",
			City:          "SF",
			StreetAddress: "123 Main St",
		},
	}

	// 模拟服务运行
	resp, err := service.Run(req)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if resp.Order == nil || resp.Order.OrderId == "" {
		t.Error("Expected valid OrderID in response")
	}

	// 验证数据库记录
	var orderCount int64
	mysql.DB.Model(&model.Order{}).Count(&orderCount)
	if orderCount != 1 {
		t.Errorf("Expected 1 order, got %d", orderCount)
	}

	var orderItems []model.OrderItem
	mysql.DB.Find(&orderItems)
	if len(orderItems) != 1 {
		t.Errorf("Expected 1 order item, got %d", len(orderItems))
	}

	// 验证 OrderItem 的 ProductId 和 Quantity
	if orderItems[0].ProductId != 1 {
		t.Errorf("Expected ProductId 1, got %d", orderItems[0].ProductId)
	}
	if orderItems[0].Quantity != 2 {
		t.Errorf("Expected Quantity 2, got %d", orderItems[0].Quantity)
	}
}
