package accounting

import (
	"fmt"
	"log"
	"os"
	"pradha-go/handlers/auth"
	"pradha-go/library"

	acc "pradha-go/models/accounting"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Coa(c *fiber.Ctx) error {
	if library.CheckAuth(c, *auth.Store) {
		var dataPerPage int64 = 0
		sort := bson.D{}
		if c.Query("dataPerPage") != "" {
			fmt.Sscan(c.Query("dataPerPage"), &dataPerPage)
		} else {
			fmt.Sscan(os.Getenv("DATA_PER_PAGE"), &dataPerPage)
		}
		if c.Query("sort") != "" {
			sort = bson.D{{Key: c.Query("sort"), Value: 1}}
		}
		data, err := acc.FindCoa(bson.M{}, sort, dataPerPage, 1)
		if err != nil {
			log.Fatal("Get Data Error")
		}
		return c.Status(200).JSON(fiber.Map{
			"status": true,
			"data":   data,
		})
	} else {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "not authorized",
		})
	}
}

func AddCoa(c *fiber.Ctx) error {
	if library.CheckAuth(c, *auth.Store) {
		error := bson.M{}
		req := struct {
			Parent      string `json:"parent"`
			Code        string `json:"code"`
			Name        string `json:"name"`
			Description string `json:"description"`
			Position    string `json:"position"`
		}{}
		if err := c.BodyParser(&req); err != nil {
			log.Println(err)
		}
		Parent, err := primitive.ObjectIDFromHex(string(req.Parent))
		if err != nil {
			error["parent"] = err
		}

		var code int32
		fmt.Sscan(req.Code, &code)
		if save, err := acc.CreateCoa(&acc.Coa{
			Parent:      Parent,
			Code:        code,
			Name:        req.Name,
			Description: req.Description,
			Position:    req.Position,
		}); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Error on save CoA",
			})
		} else {
			return c.Status(200).JSON(fiber.Map{
				"status": true,
				"result": save,
			})
		}
	} else {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "not authorized",
		})
	}
}

func EditCoa(c *fiber.Ctx) error {
	if library.CheckAuth(c, *auth.Store) {
		id := c.Params("id")
		return c.Status(200).JSON(fiber.Map{
			"status":  true,
			"message": "Edit Coa Page",
			"id":      id,
		})
	} else {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "not authorized",
		})
	}
}

func DetailCoa(c *fiber.Ctx) error {
	if library.CheckAuth(c, *auth.Store) {
		id := c.Params("id")
		return c.Status(200).JSON(fiber.Map{
			"status":  true,
			"message": "Detail Coa Page",
			"id":      id,
		})
	} else {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "not authorized",
		})
	}
}
