package datastore

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/originbenntou/2929BE/account/domain/model"
	"github.com/originbenntou/2929BE/account/domain/repository"
	"github.com/originbenntou/2929BE/shared/logger"
	"github.com/originbenntou/2929BE/shared/mysql"
)

const tableName = "session"

type sessionRepository struct {
	db mysql.DBManager
}

func NewSessionRepository(db mysql.DBManager) repository.SessionRepository {
	return &sessionRepository{db}
}

func (r sessionRepository) FindValidTokenByUserId(ctx context.Context, uid uint64) (token string, err error) {
	defer func() {
		if err != nil {
			logger.Common.Error(err.Error())
		}
	}()

	// valid session is in 24 hour
	q := "SELECT * FROM session WHERE user_id = :user_id AND DATE_ADD(updated_at, INTERVAL 1 DAY) > NOW()"

	rows, err := r.db.NamedQueryContext(ctx, q, map[string]interface{}{"user_id": uid})
	if err != nil {
		return "", err
	}

	// FIXME: 特定のカラムだけ取り出せない...
	m := &model.Session{}
	var list []*model.Session
	for rows.Next() {
		if err := rows.StructScan(m); err != nil {
			return "", err
		}
		list = append(list, m)
	}

	// no match record is ok, return empty
	if len(list) == 0 {
		return "", nil
	}

	// more than one record is error
	if len(list) > 1 {
		return "", errors.New("found token more than 1 by: " + string(uid))
	}

	// one match record
	return list[0].Token, nil
}

func (r sessionRepository) CreateSession(ctx context.Context, req *model.Session) (err error) {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			logger.Common.Error(err.Error())
		}

		if txErr := tx.CloseTransaction(err); txErr != nil {
			logger.Common.Error(txErr.Error())
		}
	}()

	q := "INSERT INTO session (token, user_id, company_id) VALUES (:token, :user_id, :company_id)"

	result, err := tx.NamedExecContext(ctx, q, req)
	if err != nil {
		return err
	}

	affect, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if affect != 1 {
		msg := fmt.Sprintf("total affected: %d", affect)
		return errors.New(msg)
	}

	return nil
}

func (r sessionRepository) UpdateSession(ctx context.Context, uid uint64) (err error) {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			logger.Common.Error(err.Error())
		}

		if txErr := tx.CloseTransaction(err); txErr != nil {
			logger.Common.Error(txErr.Error())
		}
	}()

	q := "UPDATE session SET updated_at = :updated_at WHERE user_id = :user_id AND DATE_ADD(updated_at, INTERVAL 1 DAY) > NOW()"

	result, err := tx.NamedExecContext(ctx, q, map[string]interface{}{
		"updated_at": time.Now(),
		"user_id":    uid,
	})
	if err != nil {
		return err
	}

	affect, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if affect != 1 {
		msg := fmt.Sprintf("total affected: %d", affect)
		return errors.New(msg)
	}

	return nil
}