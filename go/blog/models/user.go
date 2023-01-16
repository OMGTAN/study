package models

import (
	"database/sql"
	"time"
)

type User struct {
	Uid      int          `json:"uid"`
	UserName string       `json:"userName"`
	Passwd   string       `json:"passwd"`
	Avatar   string       `json:"avatar"`
	CreateAt time.Time    `json:"create_at"`
	UpdateAt sql.NullTime `json:"update_at"`
}

type UserInfo struct {
	Uid      int    `json:"uid"`
	UserName string `json:"userName"`
	Avatar   string `json:"avatar"`
}
