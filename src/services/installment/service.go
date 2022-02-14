package installment

import (
	"errors"
	"time"

	"github.com/geninhocell/expertscrudgo/src/interfaces"
	"github.com/geninhocell/expertscrudgo/src/repositories"
	"github.com/geninhocell/expertscrudgo/src/structs"
)

type InstallmentService struct {
	InstallmentRepository interfaces.InstallmentRepository
}

func NewInstallmentService(
	repositoryContainer repositories.RepositoryContainer,
) InstallmentService {
	return InstallmentService{
		InstallmentRepository: repositoryContainer.InstallmentRepository,
	}
}

func (is InstallmentService) Create(installment structs.Installment) error {
	if installment.DueDay < time.Now().Day() {
		return errors.New("Parcela vencida!")
	}

	return is.InstallmentRepository.Create(installment)
}
