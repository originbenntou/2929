package container

import (
	"github.com/originbenntou/2929BE/account/infrastructure/persistence/mysql"
)

func (c Container) GetAccountMySQL() mysql.UserMySQL {
	return mysql.NewUserMySQL()
}
