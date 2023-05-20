package dto

type OrderPaginationRes struct {
	CurrentPage     int         `json:"currentPage"`
	TotalPage       int         `json:"totalPage"`
	TotalData       int         `json:"totalData"`
	Limit           int         `json:"limit"`
	OrderRes        []*OrderRes `json:"orderRes"`
	SumOfTotalPrice float64     `json:"sumOfTotalPrice"`
}

func (_ *OrderPaginationRes) From(orders []*OrderRes, cp, tp, td, l int, s float64) *OrderPaginationRes {
	return &OrderPaginationRes{
		CurrentPage:     cp,
		TotalPage:       tp,
		TotalData:       td,
		Limit:           l,
		OrderRes:        orders,
		SumOfTotalPrice: s,
	}
}
