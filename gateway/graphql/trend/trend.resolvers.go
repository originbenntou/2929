package trend

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/originbenntou/2929BE/gateway/graphql/trend/generated"
	"github.com/originbenntou/2929BE/gateway/graphql/trend/model"
)

func (r *queryResolver) TrendSearch(ctx context.Context, keyword string) (int, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) TrendHistory(ctx context.Context) ([]*model.History, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) TrendSuggest(ctx context.Context, suggestID int) ([]*model.Suggest, error) {
	panic(fmt.Errorf("not implemented"))
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
