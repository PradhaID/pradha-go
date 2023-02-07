package file

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gosimple/slug"
)

func Upload(c *fiber.Ctx) error {
	path := os.Getenv("UPLOAD_PATH")

	fileID := c.FormValue("id")

	folder := ""
	if _, err := os.Stat(path + "domain/" + c.FormValue("app")); !os.IsNotExist(err) {
		folder = path + "domain/" + c.FormValue("app") + "/public/uploads/"
	} else {
		log.Println(path+"domain/"+c.FormValue("app"), "Not Exists")
	}

	if _, err := os.Stat(path + "subdomain/" + c.FormValue("app")); !os.IsNotExist(err) {
		folder = path + "subdomain/" + c.FormValue("app") + "/public/uploads/"
	} else {
		log.Println(path+"subdomain/"+c.FormValue("app"), "Not Exists")
	}

	if os.Getenv("MODE") == "local" {
		folder = path + c.FormValue("app") + "/public/uploads/"
	}

	if folder != "" && c.FormValue("app") != "" {
		if file, err := c.FormFile("file"); err == nil {
			year := time.Now().Year()
			month := time.Now().Month()
			m := fmt.Sprintf("%02d", int(month))

			folder = folder + strconv.Itoa(year) + "/" + m + "/"

			// check or create folder
			err := os.MkdirAll(folder, os.ModePerm)
			if err != nil {
				log.Println("Create folder error: ", err)
				return c.Status(401).JSON(fiber.Map{
					"status":  false,
					"message": "Failed to create folder",
				})
			}

			fileName := slug.Make(file.Filename[:len(file.Filename)-len(filepath.Ext(file.Filename))]) + filepath.Ext(file.Filename)
			// check duplicate file name
			i := 2
			for {
				if _, err := os.Stat(folder + fileName); os.IsNotExist(err) {
					c.SaveFile(file, fmt.Sprintf("./%s", folder+fileName))
					break
				}
				fileName = slug.Make(file.Filename[:len(file.Filename)-len(filepath.Ext(file.Filename))]) + "_" + strconv.Itoa(i) + filepath.Ext(file.Filename)
				i = i + 1
			}
			log.Println("File has been successfully uploaded to ", folder+fileName)

			return c.Status(201).JSON(fiber.Map{
				"success":  1,
				"status":   true,
				"message":  "File has been successfully uploaded",
				"path":     "/public/uploads/" + strconv.Itoa(year) + "/" + m + "/" + fileName,
				"filename": fileName,
				"id":       fileID,
				"file":     map[string]string{"url": "/uploads/" + strconv.Itoa(year) + "/" + m + "/" + fileName},
			})

		} else {
			return c.Status(401).JSON(fiber.Map{
				"status":  false,
				"message": "No file uploaded",
			})
		}
	} else {
		log.Println("Folder target not exists")
		if f, err := os.Open(path); err == nil {
			if files, err := f.Readdir(0); err == nil {
				for _, v := range files {
					log.Println(v.Name(), v.IsDir())
				}
			}
		}
		return c.Status(401).JSON(fiber.Map{
			"status":  false,
			"message": "Target not registered",
		})
	}
}
