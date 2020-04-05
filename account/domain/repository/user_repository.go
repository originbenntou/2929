package repository

import (
	"context"
	"github.com/originbenntou/2929BE/account/domain/model"
)

type UserRepository interface {
	FindUserByEmail(context.Context, string) (*model.User, error)
	CreateUser(context.Context, *model.User) (uint64, error)
}
