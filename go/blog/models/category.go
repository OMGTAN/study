package models

type Category struct {
	Cid      int
	Name     string
	CreateAt string
	UpdateAt string
}

type CategoryResp struct {
	HomeResp
	CategoryName string
}
