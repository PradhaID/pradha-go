package router

import (
	"pradha-go/handlers"
	"pradha-go/handlers/accounting"
	"pradha-go/handlers/auth"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

	// public handlers
	app.Get("/", handlers.HomePage)
	app.Get("/about", handlers.AboutPage)
	app.Get("/contact", handlers.ContactPage)

	// auth handlers
	sign := app.Group("/auth")
	sign.Post("/signin", auth.SignIn)
	sign.Post("/signout", auth.SignOut)
	sign.Get("/check", auth.SignCheck)

	// accounting handlers
	app.Get("/accounting/coa/*", accounting.Coa)
	app.Get("/accounting/coa/add", accounting.AddCoa)
	app.Get("/accounting/coa/edit/:id", accounting.EditCoa)
	app.Get("/accounting/coa/detail/:id", accounting.DetailCoa)
	app.Get("/accounting/account", accounting.Account)
	app.Get("/accounting/account/add", accounting.AddAccount)
	app.Get("/accounting/account/edit/:id", accounting.EditAccount)
	app.Get("/accounting/account/detail/:id", accounting.DetailAccount)
}
