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
	"github.com/originbenntou/2929BE/user/graph/generated"
	"github.com/originbenntou/2929BE/user/graph/model"
	"golang.org/x/crypto/bcrypt"
)

func (r *mutationResolver) Create(ctx context.Context, user model.User) (*model.Result, error) {
	if user.Email == "" || len(user.Password) < 0 {
		return &model.Result{Success: false}, errors.New("bad request")
	}

	passHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return &model.Result{Success: false}, errors.New("internal server request")
	}

	db, err := sql.Open("mysql", "2929:2929@tcp(2929mysql:3306)/account")
	if err != nil {
		return &model.Result{Success: false}, err
	}
	defer db.Close()

	insert, err := db.Prepare("INSERT INTO user(email, password, name, company_id, created_at, updated_at) VALUES(?, ?, ?, ?, ?, ?)")
	if err != nil {
		return &model.Result{Success: false}, err
	}
	defer insert.Close()

	result, err := insert.Exec(user.Email, passHash, "", "0", time.Now().Format("2006-1-2 15:04:05"), time.Now().Format("2006-1-2 15:04:05"))
	if err != nil {
		return &model.Result{Success: false}, err
	}

	_, err = result.LastInsertId()
	if err != nil {
		return &model.Result{Success: false}, err
	}

	return &model.Result{Success: true}, nil
}

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

func (r *queryResolver) Verify(ctx context.Context, user model.User) (*model.Result, error) {
	db, err := sql.Open("mysql", "2929:2929@tcp(2929mysql:3306)/account")
	if err != nil {
		return &model.Result{Success: false}, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT password FROM user WHERE email = ?", user.Email)
	if err != nil {
		return &model.Result{Success: false}, err
	}
	defer rows.Close()

	var m model.User
	for rows.Next() {
		if err := rows.Scan(&m.Password); err != nil {
			return &model.Result{Success: false}, err
		}
	}

	if err = rows.Err(); err != nil {
		return &model.Result{Success: false}, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(m.Password), []byte(user.Password)); err != nil {
		return &model.Result{Success: false}, err
	}

	return &model.Result{Success: true}, nil
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

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}
