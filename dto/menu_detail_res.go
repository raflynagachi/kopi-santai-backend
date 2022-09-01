package dto

type MenuDetailRes struct {
	MenuRes    *MenuRes         `json:"menu"`
	MenuOption []*MenuOptionRes `json:"menuOptions"`
}

func (_ *MenuDetailRes) From(m *MenuRes, mo []*MenuOptionRes) *MenuDetailRes {
	return &MenuDetailRes{
		MenuRes:    m,
		MenuOption: mo,
	}
}
