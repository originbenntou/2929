package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/originbenntou/2929BE/gateway/graph/generated"
	"github.com/originbenntou/2929BE/gateway/graph/model"
	"golang.org/x/crypto/bcrypt"
)

func (r *mutationResolver) CreateUser(ctx context.Context, user model.User) (bool, error) {
	if user.Email == "" || len(user.Password) < 0 {
		return false, errors.New("bad request")
	}

	passHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return false, errors.New("internal server request")
	}

	db, err := sql.Open("mysql", "2929:2929@tcp(2929mysql:3306)/account")
	if err != nil {
		return false, err
	}
	defer db.Close()

	insert, err := db.Prepare("INSERT INTO user(email, password, name, company_id, created_at, updated_at) VALUES(?, ?, ?, ?, ?, ?)")
	if err != nil {
		return false, err
	}
	defer insert.Close()

	result, err := insert.Exec(user.Email, passHash, "", "0", time.Now().Format("2006-1-2 15:04:05"), time.Now().Format("2006-1-2 15:04:05"))
	if err != nil {
		return false, err
	}

	_, err = result.LastInsertId()
	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *queryResolver) VerifyUser(ctx context.Context, user model.User) (string, error) {
	return "", nil

	db, err := sql.Open("mysql", "2929:2929@tcp(2929mysql:3306)/account")
	if err != nil {
		return "", err
	}
	defer db.Close()

	rows, err := db.Query("SELECT password FROM user WHERE email = ?", user.Email)
	if err != nil {
		return "", err
	}
	defer rows.Close()

	var m model.User
	for rows.Next() {
		if err := rows.Scan(&m.Password); err != nil {
			return "", err
		}
	}

	if err = rows.Err(); err != nil {
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(m.Password), []byte(user.Password)); err != nil {
		return "", err
	}

	return "fire!", nil
}

func (r *queryResolver) TrendSearch(ctx context.Context, keyword string) (int, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) TrendHistory(ctx context.Context) ([]*model.History, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) TrendSuggest(ctx context.Context, suggestID int) ([]*model.Suggest, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *queryResolver) FindTrend(ctx context.Context, word string) ([]*model.Suggest, error) {
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
