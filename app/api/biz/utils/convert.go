package utils

import (
	cart "Gomall/app/api/hertz_gen/api/cart"
	order "Gomall/app/api/hertz_gen/api/order"
	cart_rpc "Gomall/rpc_gen/kitex_gen/cart"
	order_rpc "Gomall/rpc_gen/kitex_gen/order"
)

// Address 类型转换
func ConvertAddress(addr *order.Address) *order_rpc.Address {
	if addr == nil {
		return nil
	}
	return &order_rpc.Address{
		StreetAddress: addr.StreetAddress,
		City:          addr.City,
		Country:       addr.Country,
		State:         addr.State,
		ZipCode:       addr.ZipCode,
	}
}

// Address 类型转换（从 RPC 响应转换回 API）
func ConvertAddressBack(addr *order_rpc.Address) *order.Address {
	if addr == nil {
		return nil
	}
	return &order.Address{
		StreetAddress: addr.StreetAddress,
		City:          addr.City,
		Country:       addr.Country,
		State:         addr.State,
		ZipCode:       addr.ZipCode,
	}
}

// OrderItem 类型转换
func ConvertOrderItems(items []*order.OrderItem) []*order_rpc.OrderItem {
	if items == nil {
		return nil
	}
	var rpcItems []*order_rpc.OrderItem
	for _, item := range items {
		rpcItems = append(rpcItems, &order_rpc.OrderItem{
			Item: ConvertCartItem(item.Item), // CartItem 类型转换
			Cost: item.Cost,
		})
	}
	return rpcItems
}

// OrderItem 结果转换（从 RPC 响应转换回 API）
func ConvertOrderItemsBack(items []*order_rpc.OrderItem) []*order.OrderItem {
	if items == nil {
		return nil
	}
	var orderItems []*order.OrderItem
	for _, item := range items {
		orderItems = append(orderItems, &order.OrderItem{
			Item: ConvertCartItemBack(item.Item), // CartItem 反向转换
			Cost: item.Cost,
		})
	}
	return orderItems
}

// CartItem 类型转换
func ConvertCartItem(item *cart.CartItem) *cart_rpc.CartItem {
	if item == nil {
		return nil
	}
	return &cart_rpc.CartItem{
		ProductId: item.ProductId,
		Quantity:  item.Quantity,
	}
}

// CartItem 结果转换（从 RPC 响应转换回 API）
func ConvertCartItemBack(item *cart_rpc.CartItem) *cart.CartItem {
	if item == nil {
		return nil
	}
	return &cart.CartItem{
		ProductId: item.ProductId,
		Quantity:  item.Quantity,
	}
}
