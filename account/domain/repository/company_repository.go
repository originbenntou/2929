package repository

import (
	"context"
	"github.com/originbenntou/2929BE/account/domain/model"
)

type CompanyRepository interface {
	FindCompanyById(context.Context, uint64) (*model.Company, error)
}
