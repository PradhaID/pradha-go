package auth

import (
	"log"
	"pradha-go/library"
	models "pradha-go/models/system"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	Store = session.New(session.Config{
		Expiration:     24 * time.Hour,
		CookieHTTPOnly: true,
		KeyGenerator:   utils.UUID,
	})
)

func SignCheck(c *fiber.Ctx) error {
	s, err := Store.Get(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "not authorized",
		})
	}

	auth := s.Get("sid")
	if auth != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "authenticated",
			"sid":     s.Get("sid"),
			"user": bson.M{
				"uid":      s.Get("uid"),
				"gid":      s.Get("gid"),
				"name":     s.Get("name"),
				"username": s.Get("username"),
				"email":    s.Get("email"),
			},
		})
	} else {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "not authorized",
		})
	}
}

func SignIn(c *fiber.Ctx) error {
	req := struct {
		User     string `json:"user"`
		Password string `json:"password"`
		Remember string `json:"remember"`
	}{}
	if err := c.BodyParser(&req); err != nil {
		log.Println(err)
	}

	// Validate user here
	if user, err := models.FindOneUser(bson.M{"$or": bson.A{bson.M{"username": req.User}, bson.M{"email": req.User}}}, bson.D{}); err != nil {
		log.Fatal("Error accured in Query user login")
	} else {
		if user["username"] != nil {
			if library.CheckPasswordHash(req.Password, user["password"].(string)) {
				s, err := Store.Get(c)
				if err != nil {
					return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
						"message": "something went wrong: " + err.Error(),
					})
				}
				sid := s.ID()
				s.Set("sid", sid)
				s.Set("uid", user["_id"].(primitive.ObjectID).Hex())
				s.Set("gid", user["group"].(primitive.ObjectID).Hex())
				s.Set("name", user["name"].(string))
				s.Set("username", user["username"].(string))
				s.Set("email", user["email"].(string))

				err = s.Save()
				if err != nil {
					return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
						"message": "something went wrong: " + err.Error(),
					})
				}

				return c.Status(fiber.StatusOK).JSON(fiber.Map{
					"sid": sid,
					"user": bson.M{
						"uid":      user["_id"].(primitive.ObjectID).Hex(),
						"gid":      user["group"].(primitive.ObjectID).Hex(),
						"name":     user["name"].(string),
						"username": user["username"].(string),
						"email":    user["email"].(string),
					},
				})

			}
			return c.Status(401).JSON(fiber.Map{
				"status":  false,
				"message": "Invalid Password",
			})
		}
	}
	return c.Status(401).JSON(fiber.Map{
		"status":  false,
		"message": "Invalid Account",
	})

}

func SignOut(c *fiber.Ctx) error {
	s, err := Store.Get(c)
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "logged out (no session)",
		})
	}
	err = s.Destroy()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "something went wrong: " + err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "logged out",
	})
}
