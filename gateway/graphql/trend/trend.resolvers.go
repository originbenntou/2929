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
	return 1, nil
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

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
