package model

import "time"

type User struct {
	Id        uint64
	Email     string
	Password  []byte
	Name      string
	CompanyId uint64
	CreatedAt time.Time
	UpdatedAt time.Time
}
