package main

import (
	"database/sql"
	"log"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"

	"github.com/ksploitx/user-age-api/config"
	"github.com/ksploitx/user-age-api/internal/handler"
	"github.com/ksploitx/user-age-api/internal/logger"
	"github.com/ksploitx/user-age-api/internal/repository"
	"github.com/ksploitx/user-age-api/internal/routes"
	"github.com/ksploitx/user-age-api/internal/service"
)

func main() {
	cfg := config.Load()
	logger.Init()

	db, err := sql.Open("postgres", cfg.DSN())
	if err != nil {
		log.Fatal("failed to connect to db:", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal("db not reachable:", err)
	}

	// Wire layers
	repo := repository.NewUserRepository(db)
	svc := service.NewUserService(repo)
	h := handler.NewUserHandler(svc)

	app := fiber.New()

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok"})
	})

	routes.Register(app, h)

	log.Fatal(app.Listen(":" + cfg.AppPort))
}
