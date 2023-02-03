package library

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CheckAuth(c *fiber.Ctx, store session.Store) bool {
	s, err := store.Get(c)
	if err != nil {
		return false
	}
	auth := s.Get("sid")
	return auth != nil
}

func GetUID(c *fiber.Ctx, store session.Store) (primitive.ObjectID, error) {
	s, err := store.Get(c)
	objID, _ := primitive.ObjectIDFromHex("000000000000000000000000")
	if err != nil {
		panic(err)
	}
	if err != nil {
		return objID, err
	}

	uid := s.Get("uid")
	if uid != nil {
		return objID, err
	}
	objID, _ = primitive.ObjectIDFromHex(s.Get("uid").(string))
	return objID, nil
}
