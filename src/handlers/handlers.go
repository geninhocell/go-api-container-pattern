package handlers

import (
	"github.com/geninhocell/expertscrudgo/src/handlers/installment"
	"github.com/geninhocell/expertscrudgo/src/services"
	"github.com/gofiber/fiber/v2"
)

func NewHandlerContainer(router fiber.Router, serviceContainer services.ServiceContainer) {
	installment.NewInstallmentHandler(router, serviceContainer).SetRoutes()
}
