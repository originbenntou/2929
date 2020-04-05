package model

import "time"

type User struct {
	Id        uint64
	Email     string
	PassHash  []byte
	Name      string
	CompanyId uint64
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (m *User) GetEmail() string {
	return m.Email
}
