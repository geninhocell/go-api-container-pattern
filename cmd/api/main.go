package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/geninhocell/expertscrudgo/src/handlers"
	"github.com/geninhocell/expertscrudgo/src/repositories"
	"github.com/geninhocell/expertscrudgo/src/services"
	"github.com/geninhocell/expertscrudgo/src/structs"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	server := fiber.New()

	dialector := postgres.Open(fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		"kesavan.db.elephantsql.com",
		"plwtfwrr",
		"CT5yy5CV-m5y9q9q-voBqMN84BGkZTdW",
		"plwtfwrr",
		"5432",
	))

	db, err := gorm.Open(dialector)

	db.AutoMigrate(&structs.Installment{})

	repositoryContainer := repositories.GetRepositories(db)

	servicesContainer := services.GetServices(repositoryContainer)

	handlers.NewHandlerContainer(server, servicesContainer)

	if err != nil {
		log.Fatalf("Banco não conectou. Err=%v", err.Error())
	}

	server.Get("/health", func(f *fiber.Ctx) error {
		return f.Status(http.StatusOK).JSON("pong :)")
	})

	if err := server.Listen(":3001"); err != nil {
		log.Fatalf("Aplicação não subiu. Err=%v", err.Error())
	}
}
