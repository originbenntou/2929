package container

import (
	"github.com/originbenntou/2929BE/account/domain/repository"
	repo "github.com/originbenntou/2929BE/account/infrastructure/datastore"
	"github.com/originbenntou/2929BE/shared/mysql"
)

func (c Container) GetUserRepository(db mysql.DBManager) repository.UserRepository {
	return repo.NewUserRepository(db)
}

func (c Container) GetCompanyRepository(db mysql.DBManager) repository.CompanyRepository {
	return repo.NewCompanyRepository(db)
}
