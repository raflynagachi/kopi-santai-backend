package dto

import (
	"time"

	"github.com/raflynagachi/kopi-santai-backend/model"
)

type OrderRes struct {
	ID            uint              `json:"id"`
	UserID        uint              `json:"userID"`
	CouponID      uint              `json:"couponID"`
	OrderedDate   time.Time         `json:"orderedDate"`
	TotalPrice    float64           `json:"totalPrice"`
	Coupon        *CouponRes        `json:"coupon"`
	Delivery      *DeliveryRes      `json:"delivery"`
	PaymentOption *PaymentOptionRes `json:"paymentOption"`
	OrderItems    []*OrderItemRes   `json:"orderItems"`
}

func (_ *OrderRes) From(o *model.Order) *OrderRes {
	var coupon *CouponRes
	var couponID uint
	if o.Coupon != nil {
		coupon = new(CouponRes).FromCoupon(o.Coupon)
		couponID = *o.CouponID
	}
	delivery := new(DeliveryRes).FromDelivery(o.Delivery)
	paymentOpt := new(PaymentOptionRes).FromPaymentOption(o.PaymentOption)

	var orderItemsRes []*OrderItemRes
	for _, item := range o.OrderItems {
		menuRes := new(MenuRes).FromMenu(item.Menu)
		orderItemsRes = append(orderItemsRes, new(OrderItemRes).From(item, menuRes))
	}

	return &OrderRes{
		ID:            o.ID,
		UserID:        o.UserID,
		CouponID:      couponID,
		OrderedDate:   o.OrderedDate,
		TotalPrice:    o.TotalPrice,
		Coupon:        coupon,
		Delivery:      delivery,
		PaymentOption: paymentOpt,
		OrderItems:    orderItemsRes,
	}
}
