package repository

import "context"

type PlanRepository interface {
	FindCapacityByCompanyId(ctx context.Context, cid uint64) (uint64, error)
}
