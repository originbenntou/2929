package repository

import (
	"context"
	"github.com/originbenntou/2929BE/account/domain/model"
)

type SessionRepository interface {
	FindExistTokenByUserId(context.Context, uint64) (string, error)
	CreateSession(context.Context, *model.Session) error
	UpdateSession(context.Context, uint64) error
	CountValidSessionByCompanyId(context.Context, uint64) (uint64, error)
}
