package interfaces

import "github.com/geninhocell/expertscrudgo/src/structs"

type InstallmentService interface {
	Create(installment structs.Installment) error
}

type InstallmentRepository interface {
	Create(installment structs.Installment) error
}
