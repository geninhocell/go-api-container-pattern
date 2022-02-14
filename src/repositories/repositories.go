package repositories

import (
	"github.com/geninhocell/expertscrudgo/src/interfaces"
	"github.com/geninhocell/expertscrudgo/src/repositories/installment"
	"gorm.io/gorm"
)

type RepositoryContainer struct {
	InstallmentRepository interfaces.InstallmentRepository
}

func GetRepositories(db *gorm.DB) RepositoryContainer {
	return RepositoryContainer{
		InstallmentRepository: installment.NewInstallmentRepository(db),
	}
}
