package model

import "time"

type Company struct {
	Id        uint64    `db:"id"`
	Name      string    `db:"name"`
	PlanId    uint64    `db:"plan_id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func (m *User) GetName() string {
	return m.Name
}
