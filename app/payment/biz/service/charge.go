package service

import (
	"Gomall/app/payment/biz/dal/mysql"
	"Gomall/app/payment/biz/model"
	payment "Gomall/rpc_gen/kitex_gen/payment"
	"context"
	"fmt"
	"runtime/debug"
	"strconv"
	"time"

	"github.com/cloudwego/kitex/pkg/kerrors"
	creditcard "github.com/durango/go-credit-card"
	"github.com/google/uuid"
)

type ChargeService struct {
	ctx context.Context
} // NewChargeService new ChargeService
func NewChargeService(ctx context.Context) *ChargeService {
	return &ChargeService{ctx: ctx}
}

// Run create note info
func (s *ChargeService) Run(req *payment.ChargeReq) (resp *payment.ChargeResp, err error) {
	// Finish your business logic.
	fmt.Printf("Received req: %+v\n", req)
	if req.CreditCard == nil {
		return nil, kerrors.NewBizStatusError(400, "credit card info is missing")
	}
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Panic: %v\nStack: %s\n", r, debug.Stack())
			err = kerrors.NewGRPCBizStatusError(500, "internal server error")
			resp = nil
		}
	}()
	resp = &payment.ChargeResp{}
	card := creditcard.Card{
		Number: req.CreditCard.CreditCardNumber,
		Cvv:    strconv.Itoa(int(req.CreditCard.CreditCardCvv)),
		Month:  strconv.Itoa(int(req.CreditCard.CreditCardExpirationMonth)),
		Year:   strconv.Itoa(int(req.CreditCard.CreditCardExpirationYear)),
	}

	err = card.Validate(true)
	if err != nil {
		return nil, kerrors.NewBizStatusError(400, err.Error())
	}

	transaction_id, err := uuid.NewRandom()
	if err != nil {
		return nil, kerrors.NewGRPCBizStatusError(500, err.Error())
	}

	paymentLog := &model.PaymentLog{
		UserId:        req.UserId,
		OrderId:       req.OrderId,
		TransactionId: transaction_id.String(),
		Amount:        req.Amount,
		PayAt:         time.Now(),
	}

	err = model.CreatePaymentLog(mysql.DB, s.ctx, paymentLog)
	if err != nil {
		return nil, kerrors.NewGRPCBizStatusError(500, err.Error())
	}

	return &payment.ChargeResp{
		TransactionId: transaction_id.String(),
	}, nil
}
