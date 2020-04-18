package support

import (
	"context"

	pbAccount "github.com/originbenntou/2929BE/proto/account/go"
)

type contextKeyUser struct{}

func GetUserFromContext(ctx context.Context) *pbAccount.User {
	u := ctx.Value(contextKeyUser{})
	pUser, ok := u.(*pbAccount.User)
	if !ok {
		return nil
	}
	return pUser
}

func AddUserToContext(ctx context.Context, user *pbAccount.User) context.Context {
	return context.WithValue(ctx, contextKeyUser{}, user)
}
