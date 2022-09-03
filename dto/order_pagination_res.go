package dto

type OrderPaginationRes struct {
	CurrentPage int         `json:"currentPage"`
	TotalPage   int         `json:"totalPage"`
	TotalData   int         `json:"totalData"`
	Limit       int         `json:"limit"`
	OrderRes    []*OrderRes `json:"orderRes"`
}

func (_ *OrderPaginationRes) From(orders []*OrderRes, cp, tp, td, l int) *OrderPaginationRes {
	return &OrderPaginationRes{
		CurrentPage: cp,
		TotalPage:   tp,
		TotalData:   td,
		Limit:       l,
		OrderRes:    orders,
	}
}
