package file

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

func Upload(c *fiber.Ctx) error {
	path := "../../"
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		fmt.Println(file.Name(), file.IsDir())
	}

	year := time.Now().Year()
	month := time.Now().Month()

	fileID := c.FormValue("id")
	fileType := c.FormValue("type")
	if file, err := c.FormFile("file"); err == nil {
		subPath := strconv.Itoa(year) + "/" + strconv.Itoa(int(month))
		path = path + subPath
		from := c.OriginalURL()
		fmt.Println(year, month, fileID, fileType, from, path)
		fmt.Println(file.Filename)
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  true,
		"message": "Upload Page",
	})
}
