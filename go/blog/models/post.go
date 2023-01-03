package models

import (
	"blog/config"
	"html/template"
	"time"
)

type Post struct {
	Pid        int       `json:"pid"`
	Title      string    `json:"title"`
	Slug       string    `json:"slug"`
	Content    string    `json:"content"`
	Markdown   string    `json:"markdown"`
	CategoryId int       `json:"categoryId"`
	UserId     int       `json:"userId"`
	ViewCount  int       `json:"viewCount"`
	Type       int       `json:"type"` //0:普通文章 1：自定义文章
	CreateAt   time.Time `json:"createAt"`
	UpdateAt   time.Time `json:"updateAt"`
}

type PostMore struct {
	Pid          int           `json:"pid"`
	Title        string        `json:"title"`
	Slug         string        `json:"slug"`
	Content      template.HTML `json:"content"`
	CategoryId   int           `json:"categoryId"`
	CategoryName string        `json:"categoryName"`
	UserId       int           `json:"userId"`
	UserName     string        `json:"userName"`
	ViewCount    int           `json:"viewCount"`
	Type         int           `json:"type"` //0:普通文章 1：自定义文章
	CreateAt     string        `json:"createAt"`
	UpdateAt     string        `json:"updateAt"`
}
type PostReq struct {
	Pid        int    `json:"pid"`
	Title      string `json:"title"`
	Slug       string `json:"slug"`
	Content    string `json:"content"`
	Markdown   string `json:"markdown"`
	CategoryId int    `json:"categoryId"`
	UserId     int    `json:"userId"`
	Type       int    `json:"type"` //0:普通文章 1：自定义文章
}

type SearchResp struct {
	Pid   int    `orm:"pid" json:"pid"`
	Title string `orm:"title" json:"title"`
}

type PostRes struct {
	config.Viewer
	config.SystemConfig
	Article PostMore
}
