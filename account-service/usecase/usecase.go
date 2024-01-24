package usecase

import (
	"account-service/config"
	"account-service/domain"

	"github.com/sirupsen/logrus"
)

type Usecases struct {
	User UserInterface
}

func Init(cfg *config.Value, logger *logrus.Logger, dom *domain.Domains) *Usecases {
	return &Usecases{
		User: initUser(cfg, dom.User),
	}
}
