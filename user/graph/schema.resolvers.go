package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"time"

	"github.com/originbenntou/2929BE/user/graph/generated"
	"github.com/originbenntou/2929BE/user/graph/model"
)

func (r *mutationResolver) CreateText(ctx context.Context, textCreateInput model.TextCreateInput) (*model.Text, error) {
	var text *model.Text

	fmt.Println(ctx)

	text = &model.Text{
		TextID:    "20001",
		TextData:  textCreateInput.TextData,
		Length:    1,
		Bot:       true,
		Version:   1.1,
		Rank:      model.RankCopper,
		CreatedAt: time.Now().Format("2006-1-2 15:04:05"),
		UpdatedAt: time.Now().Format("2006-1-2 15:04:05"),
	}

	return text, nil
}

func (r *mutationResolver) UpdateText(ctx context.Context, textUpdateInput model.TextUpdateInput) (*model.Text, error) {
	var text *model.Text

	text = &model.Text{
		TextID:    textUpdateInput.TextID,
		TextData:  textUpdateInput.TextData,
		Length:    1,
		Bot:       true,
		Version:   1.1,
		Rank:      model.RankCopper,
		CreatedAt: time.Now().Format("2006-1-2 15:04:05"),
		UpdatedAt: time.Now().Format("2006-1-2 15:04:05"),
	}

	return text, nil
}

func (r *mutationResolver) DeleteText(ctx context.Context, textDeleteInput model.TextDeleteInput) (*model.MutationResult, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Text(ctx context.Context, textCondition model.TextCondition) ([]*model.Text, error) {
	var text []*model.Text

	ids := textCondition.TextID
	for _, id := range ids {
		text = append(text, &model.Text{
			TextID:    *id,
			TextData:  "NICE RESPONSE!",
			Length:    len(ids),
			Bot:       true,
			Version:   1.0,
			Rank:      model.RankGold,
			CreatedAt: time.Now().Format("2006-1-2 15:04:05"),
			UpdatedAt: time.Now().Format("2006-1-2 15:04:05"),
		})
	}

	return text, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
