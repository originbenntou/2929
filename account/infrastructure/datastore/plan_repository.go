package datastore

import (
	"context"
	"errors"

	"github.com/originbenntou/2929BE/account/domain/repository"
	"github.com/originbenntou/2929BE/shared/logger"
	"github.com/originbenntou/2929BE/shared/mysql"
)

type planRepository struct {
	db mysql.DBManager
}

func NewPlanRepository(db mysql.DBManager) repository.PlanRepository {
	return &planRepository{db}
}

func (r planRepository) FindCapacityByCompanyId(ctx context.Context, cid uint64) (cap uint64, err error) {
	defer func() {
		if err != nil {
			logger.Common.Error(err.Error())
		}
	}()

	q := "SELECT capacity FROM plan WHERE id = :company_id"

	rows, err := r.db.NamedQueryContext(ctx, q, map[string]interface{}{"company_id": cid})
	if err != nil {
		return 0, err
	}

	var list []uint64
	for rows.Next() {
		if err := rows.Scan(&cap); err != nil {
			return 0, err
		}
		list = append(list, cap)
	}

	// no match record is ok, return empty
	if len(list) == 0 {
		return 0, nil
	}

	// more than one record is error
	if len(list) > 1 {
		return 0, errors.New("found capacity more than 1 by: " + string(cid))
	}

	// one match record
	return list[0], nil
}
