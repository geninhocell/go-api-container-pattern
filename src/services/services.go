package services

import (
	"github.com/geninhocell/expertscrudgo/src/interfaces"
	"github.com/geninhocell/expertscrudgo/src/repositories"
	"github.com/geninhocell/expertscrudgo/src/services/installment"
)

type ServiceContainer struct {
	InstallmentService interfaces.InstallmentService
}

func GetServices(repositoryContainer repositories.RepositoryContainer) ServiceContainer {
	return ServiceContainer{
		InstallmentService: installment.NewInstallmentService(repositoryContainer),
	}
}
