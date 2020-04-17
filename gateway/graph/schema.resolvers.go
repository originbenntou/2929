package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/originbenntou/2929BE/gateway/graph/generated"
	"github.com/originbenntou/2929BE/gateway/graph/model"
	pbAccount "github.com/originbenntou/2929BE/proto/account/go"
)

func (r *mutationResolver) RegisterUser(ctx context.Context, user model.User) (bool, error) {
	pbUser, err := r.userClient.RegisterUser(ctx, &pbAccount.RegisterUserRequest{
		Email:     user.Email,
		Password:  user.Password,
		Name:      "",
		CompanyId: 0,
	})
	if err != nil {
		//logger.Common.Error(err.Error())
	}

	return pbUser.UserId > 0, nil
}

func (r *mutationResolver) UpdateUser(ctx context.Context, user model.User) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteUser(ctx context.Context, user model.User) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) LoginUser(ctx context.Context, email string, password string) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) TrendSearch(ctx context.Context, keyword string) (int, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) TrendHistory(ctx context.Context) ([]*model.History, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) TrendSuggest(ctx context.Context, suggestID int) ([]*model.Suggest, error) {
	return []*model.Suggest{
		{
			Keyword: "リモート",
			ChildSuggests: []*model.ChildSuggest{
				{
					Word: "Zoom",
					Growth: &model.Growth{
						Short:  model.ArrowDown,
						Midium: model.ArrowFlat,
						Long:   model.ArrowUp,
					},
					Graphs: []*model.Graph{
						{
							Date:  "20200326",
							Value: 100,
						},
						{
							Date:  "20200327",
							Value: 99,
						},
						{
							Date:  "20200328",
							Value: 98,
						},
					},
				},
				{
					Word: "Skype",
					Growth: &model.Growth{
						Short:  model.ArrowDown,
						Midium: model.ArrowFlat,
						Long:   model.ArrowUp,
					},
					Graphs: []*model.Graph{
						{
							Date:  "20200326",
							Value: 98,
						},
						{
							Date:  "20200327",
							Value: 99,
						},
						{
							Date:  "20200328",
							Value: 100,
						},
					},
				},
				{
					Word: "GoogleHang",
					Growth: &model.Growth{
						Short:  model.ArrowDown,
						Midium: model.ArrowFlat,
						Long:   model.ArrowUp,
					},
					Graphs: []*model.Graph{
						{
							Date:  "20200326",
							Value: 100,
						},
						{
							Date:  "20200327",
							Value: 100,
						},
						{
							Date:  "20200328",
							Value: 100,
						},
					},
				},
			},
		},
		{
			Keyword: "フリーランス",
			ChildSuggests: []*model.ChildSuggest{
				{
					Word: "保証",
					Growth: &model.Growth{
						Short:  model.ArrowDown,
						Midium: model.ArrowFlat,
						Long:   model.ArrowUp,
					},
					Graphs: []*model.Graph{
						{
							Date:  "20200326",
							Value: 100,
						},
						{
							Date:  "20200327",
							Value: 99,
						},
						{
							Date:  "20200328",
							Value: 98,
						},
					},
				},
				{
					Word: "税金",
					Growth: &model.Growth{
						Short:  model.ArrowDown,
						Midium: model.ArrowFlat,
						Long:   model.ArrowUp,
					},
					Graphs: []*model.Graph{
						{
							Date:  "20200326",
							Value: 98,
						},
						{
							Date:  "20200327",
							Value: 99,
						},
						{
							Date:  "20200328",
							Value: 100,
						},
					},
				},
				{
					Word: "チャンス",
					Growth: &model.Growth{
						Short:  model.ArrowDown,
						Midium: model.ArrowFlat,
						Long:   model.ArrowUp,
					},
					Graphs: []*model.Graph{
						{
							Date:  "20200326",
							Value: 100,
						},
						{
							Date:  "20200327",
							Value: 100,
						},
						{
							Date:  "20200328",
							Value: 100,
						},
					},
				},
			},
		},
		{
			Keyword: "予防",
			ChildSuggests: []*model.ChildSuggest{
				{
					Word: "手洗い",
					Growth: &model.Growth{
						Short:  model.ArrowDown,
						Midium: model.ArrowFlat,
						Long:   model.ArrowUp,
					},
					Graphs: []*model.Graph{
						{
							Date:  "20200326",
							Value: 100,
						},
						{
							Date:  "20200327",
							Value: 99,
						},
						{
							Date:  "20200328",
							Value: 98,
						},
					},
				},
				{
					Word: "濃厚接触",
					Growth: &model.Growth{
						Short:  model.ArrowDown,
						Midium: model.ArrowFlat,
						Long:   model.ArrowUp,
					},
					Graphs: []*model.Graph{
						{
							Date:  "20200326",
							Value: 98,
						},
						{
							Date:  "20200327",
							Value: 99,
						},
						{
							Date:  "20200328",
							Value: 100,
						},
					},
				},
				{
					Word: "クラスタ",
					Growth: &model.Growth{
						Short:  model.ArrowDown,
						Midium: model.ArrowFlat,
						Long:   model.ArrowUp,
					},
					Graphs: []*model.Graph{
						{
							Date:  "20200326",
							Value: 100,
						},
						{
							Date:  "20200327",
							Value: 100,
						},
						{
							Date:  "20200328",
							Value: 100,
						},
					},
				},
			},
		},
	}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

func (r *queryResolver) VerifyUser(ctx context.Context, email string, password string) (string, error) {
	return "", nil
}
