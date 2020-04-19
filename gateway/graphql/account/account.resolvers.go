package account

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"github.com/originbenntou/2929BE/gateway/graphql/account/generated"
	"github.com/originbenntou/2929BE/gateway/graphql/account/model"
	pbAccount "github.com/originbenntou/2929BE/proto/account/go"
)

func (r *mutationResolver) RegisterUser(ctx context.Context, user model.User) (bool, error) {
	pbUser, err := r.accountClient.RegisterUser(ctx, &pbAccount.RegisterUserRequest{
		Email:     user.Email,
		Password:  user.Password,
		Name:      user.Name,
		CompanyId: uint64(user.CompanyID),
	})
	if err != nil {
		return false, err
		//logger.Common.Error(err.Error())
	}

	if pbUser == nil {
		return false, nil
	}

	return pbUser.UserId > 0, nil
}

func (r *mutationResolver) UpdateUser(ctx context.Context, user model.User) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) VerifyUser(ctx context.Context, email string, password string) (string, error) {
	pbToken, err := r.accountClient.VerifyUser(ctx, &pbAccount.VerifyUserRequest{
		Email:    email,
		Password: password,
	})
	if err != nil {
		return "", err
	}

	return pbToken.Token, nil
}

func (r *queryResolver) RecoveryUser(ctx context.Context, email string) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
