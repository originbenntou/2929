package model

import "time"

type User struct {
	Id        uint64    `db:"id"`
	Email     string    `db:"email"`
	PassHash  []byte    `db:"password"`
	Name      string    `db:"name"`
	CompanyId uint64    `db:"company_id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func (m *User) GetEmail() string {
	return m.Email
}
