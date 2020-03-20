package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/originbenntou/2929BE/user/graph/generated"
	"github.com/originbenntou/2929BE/user/graph/model"
)

func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Text(ctx context.Context, textCondition *model.TextCondition) (*model.Text, error) {
	id := textCondition.TextID
	d := "common!"
	text := &model.Text{TextID: id, TextData: d}
	return text, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
