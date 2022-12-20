package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/CSSIT22/modlifesdrive/db"
	"github.com/CSSIT22/modlifesdrive/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

//type File {
//
//}

func main() {

	//apikey, isok := os.LookupEnv("API_KEY")
	if val, isok := os.LookupEnv("GO_ENV"); !isok || val != "production" {
		if err := godotenv.Load(); err != nil {
			log.Fatal("Fail to load ENV")
		}
	}

	if err := db.SetUpDb(); err != nil {
		log.Fatal(err)
	}

	app := fiber.New(fiber.Config{
		BodyLimit: 100 * 1024 * 1,
	})

	app.Use(func(c *fiber.Ctx) error {
		authorization := c.Get("Authorization")

		fmt.Println(authorization)
		fmt.Println(strings.Replace(authorization, "Bearer ", "", 1))
		if !strings.HasPrefix(authorization, "Bearer ") || strings.Replace(authorization, "Bearer ", "", 1) != os.Getenv("API_KEY") {
			return c.Status(401).SendString("Unauthorized")
		}
		return c.Next()
	})

	app.Post("/", routes.Upload)

	app.Get("/:id", routes.GetFile)

	app.Delete("/", routes.DeleteFile)

	app.Listen(":8001")

}
