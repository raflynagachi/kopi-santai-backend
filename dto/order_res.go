package dto

import (
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/model"
	"time"
)

type OrderRes struct {
	UserID        uint
	CouponID      uint
	OrderedDate   time.Time
	TotalPrice    float64
	IsActive      bool
	Coupon        *CouponRes
	Delivery      *DeliveryRes
	PaymentOption *PaymentOptionRes
	OrderItems    []*OrderItemRes
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
		UserID:        o.UserID,
		CouponID:      couponID,
		OrderedDate:   o.OrderedDate,
		TotalPrice:    o.TotalPrice,
		IsActive:      o.IsActive,
		Coupon:        coupon,
		Delivery:      delivery,
		PaymentOption: paymentOpt,
		OrderItems:    orderItemsRes,
	}
}
