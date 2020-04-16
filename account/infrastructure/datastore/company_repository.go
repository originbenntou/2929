package datastore

import (
	"context"
	"errors"
	"fmt"
	"github.com/originbenntou/2929BE/account/domain/model"
	"github.com/originbenntou/2929BE/account/domain/repository"
	"github.com/originbenntou/2929BE/shared/logger"
	"github.com/originbenntou/2929BE/shared/mysql"
)

type companyRepository struct {
	db mysql.DBManager
}

func NewCompanyRepository(db mysql.DBManager) repository.CompanyRepository {
	return &companyRepository{db}
}

func (r companyRepository) FindCompanyById(ctx context.Context, id uint64) (m *model.Company, err error) {
	defer func() {
		if err != nil {
			logger.Common.Error(err.Error())
		}
	}()

	q := "SELECT * FROM company WHERE id = :id"

	rows, err := r.db.NamedQueryContext(ctx, q, map[string]interface{}{"id": id})
	if err != nil {
		return nil, err
	}

	// no match record is ok
	if !rows.Next() {
		return nil, nil
	}

	m = &model.Company{}
	c := 0
	for rows.Next() {
		if err := rows.StructScan(m); err != nil {
			return nil, err
		}
		c++
	}

	if c > 1 {
		msg := fmt.Sprintf("found company more than 1 by: %d", id)
		return nil, errors.New(msg)
	}

	// one match record
	return m, nil
}
