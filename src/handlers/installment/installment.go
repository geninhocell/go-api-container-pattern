package installment

import (
	"net/http"

	"github.com/geninhocell/expertscrudgo/src/interfaces"
	"github.com/geninhocell/expertscrudgo/src/services"
	"github.com/geninhocell/expertscrudgo/src/structs"
	"github.com/gofiber/fiber/v2"
)

type InstallmentHandler struct {
	router             fiber.Router
	InstallmentService interfaces.InstallmentService
}

func NewInstallmentHandler(
	router fiber.Router,
	serviceContainer services.ServiceContainer,
) InstallmentHandler {
	return InstallmentHandler{
		router:             router,
		InstallmentService: serviceContainer.InstallmentService,
	}
}

func (ih InstallmentHandler) SetRoutes() {
	group := ih.router.Group("/installment")

	group.Post("/", ih.CreateInstallment)
}

func (ih InstallmentHandler) CreateInstallment(f *fiber.Ctx) error {
	var body structs.Installment

	err := f.BodyParser(&body)

	if err != nil {
		return f.Status(http.StatusBadRequest).JSON(structs.Response{
			Data: err.Error(),
			Tag:  "BAD_REQUEST",
		})
	}

	err = ih.InstallmentService.Create(body)

	if err != nil {
		return f.Status(http.StatusInternalServerError).JSON(structs.Response{
			Data: err.Error(),
			Tag:  "INTERNAL_SERVER_ERROR",
		})
	}

	return f.Status(http.StatusCreated).JSON(structs.Response{
		Data: "Installment created",
	})
}
