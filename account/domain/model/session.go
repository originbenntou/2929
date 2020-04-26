package model

import "time"

type Session struct {
	Id        uint64    `db:"id"`
	Token     string    `db:"token"`
	UserId    uint64    `db:"user_id"`
	CompanyId uint64    `db:"company_id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
