package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"pradha-go/initializers"
	"pradha-go/router"
	"runtime/pprof"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/encryptcookie"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/template/html"
)

func init() {
	initializers.LoadEnv()
}

func deferFunc() int {
	// defer function
	defer func() {
		fmt.Println("Closing MongoDB connection...")
		initializers.Client.Disconnect(context.Background())
		pprof.StopCPUProfile()
	}()

	return 1
}

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for sig := range c {
			log.Printf("captured %v, stopping profiler and exiting..", sig)
			pprof.StopCPUProfile()
			os.Exit(deferFunc())
		}
	}()

	engine := html.New("./views/", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Use(encryptcookie.New(encryptcookie.Config{
		Key: os.Getenv("APP_SECRET"), Except: []string{"csrf_"}}))
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins:     os.Getenv("ALLOWED_CLIENT"),
		AllowCredentials: true,
		AllowMethods:     "GET,POST,PUT,DELETE",
	}))
	//app.Use(csrf.New())
	app.Use(requestid.New())
	app.Use(recover.New())

	router.SetupRoutes(app)

	log.Fatal(app.Listen(":6060"))
	os.Exit(deferFunc())
}
