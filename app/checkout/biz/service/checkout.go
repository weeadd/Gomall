package service

import (
	"Gomall/app/checkout/infra/rpc"
	"context"
	"errors"
	"fmt"
	"strconv"

	cart_rpc "Gomall/rpc_gen/kitex_gen/cart"
	checkout_rpc "Gomall/rpc_gen/kitex_gen/checkout"
	order_rpc "Gomall/rpc_gen/kitex_gen/order"
	payment_rpc "Gomall/rpc_gen/kitex_gen/payment"
	product_rpc "Gomall/rpc_gen/kitex_gen/product"

	"github.com/cloudwego/kitex/pkg/klog"
)

type CheckoutService struct {
	ctx context.Context
} // NewCheckoutService new CheckoutService
func NewCheckoutService(ctx context.Context) *CheckoutService {
	return &CheckoutService{ctx: ctx}
}

// Run create note info
func (s *CheckoutService) Run(req *checkout_rpc.CheckoutReq) (resp *checkout_rpc.CheckoutResp, err error) {
	// get cart
	cartResult, err := rpc.CartClient.GetCart(s.ctx, &cart_rpc.GetCartReq{
		UserId: req.UserId,
	})

	if err != nil {
		klog.Error(err)
		err = fmt.Errorf("GetCart.err:%v", err)
		return
	}

	if cartResult == nil || len(cartResult.Items) == 0 {
		err = errors.New("cart is empty")
		return
	}

	var (
		oi    []*order_rpc.OrderItem
		total float32
	)

	for _, cartItem := range cartResult.Items {
		productResp, resultErr := rpc.ProductClient.GetProduct(s.ctx, &product_rpc.GetProductReq{Id: cartItem.ProductId})
		if resultErr != nil {
			klog.Error(resultErr)
			err = resultErr
			return
		}
		if productResp.Product == nil {
			continue
		}
		p := productResp.Product
		cost := p.Price * float32(cartItem.Quantity)
		total += cost
		oi = append(oi, &order_rpc.OrderItem{
			Item: &cart_rpc.CartItem{ProductId: cartItem.ProductId, Quantity: cartItem.Quantity},
			Cost: cost,
		})
	}

	// create order
	orderReq := &order_rpc.PlaceOrderReq{
		UserId:       req.UserId,
		UserCurrency: "USD",
		OrderItems:   oi,
		Email:        req.Email,
	}
	if req.Address != nil {
		addr := req.Address
		zipCodeInt, _ := strconv.Atoi(addr.ZipCode)
		orderReq.Address = &order_rpc.Address{
			StreetAddress: addr.StreetAddress,
			City:          addr.City,
			Country:       addr.Country,
			State:         addr.State,
			ZipCode:       int32(zipCodeInt),
		}
	}
	orderResult, err := rpc.OrderClient.PlaceOrder(s.ctx, orderReq)
	if err != nil {
		err = fmt.Errorf("PlaceOrder.err:%v", err)
		return
	}
	klog.Info("orderResult", orderResult)

	// empty cart
	emptyResult, err := rpc.CartClient.EmptyCart(s.ctx, &cart_rpc.EmptyCartReq{UserId: req.UserId})
	if err != nil {
		err = fmt.Errorf("EmptyCart.err:%v", err)
		return
	}
	klog.Info(emptyResult)

	// charge
	var orderId string
	if orderResult != nil {
		orderId = orderResult.Order.OrderId
	}
	payReq := &payment_rpc.ChargeReq{
		UserId:  req.UserId,
		OrderId: orderId,
		Amount:  total,
		CreditCard: &payment_rpc.CreditCardInfo{
			CreditCardNumber:          req.CreditCard.CreditCardNumber,
			CreditCardExpirationYear:  req.CreditCard.CreditCardExpirationYear,
			CreditCardExpirationMonth: req.CreditCard.CreditCardExpirationMonth,
			CreditCardCvv:             req.CreditCard.CreditCardCvv,
		},
	}
	paymentResult, err := rpc.PaymentClient.Charge(s.ctx, payReq)
	if err != nil {
		err = fmt.Errorf("Charge.err:%v", err)
		return
	}
	klog.Info(paymentResult)

	// change order state
	klog.Info(orderResult)
	_, err = rpc.OrderClient.MarkOrderPaid(s.ctx, &order_rpc.MarkOrderPaidReq{UserId: req.UserId, OrderId: orderId})
	if err != nil {
		klog.Error(err)
		return
	}

	resp = &checkout_rpc.CheckoutResp{
		OrderId:       orderId,
		TransactionId: paymentResult.TransactionId,
	}
	return resp, nil
}
