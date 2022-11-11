package routes

import (
	"fmt"
	"github.com/CSSIT22/modlifesdrive/db"
	"github.com/gofiber/fiber/v2"
	"github.com/lucsky/cuid"
	"os"
	"strings"
)

type UploadRespond struct {
	Id   string
	Name string
}

func extention(name string) string {
	ext := strings.Split(name, ".")
	if len(ext) == 1 {
		return ""
	}
	return ext[len(ext)-1]
}

func Upload(c *fiber.Ctx) error {
	var respond []UploadRespond
	if form, err := c.MultipartForm(); err == nil {
		files := form.File["upload"]
		var filenames []string
		dir, err := os.Getwd()
		if err != nil {
			return err
		}
		for _, file := range files {
			//fmt.Println(dir)
			f := db.FileInfo{Id: cuid.New(), Name: file.Filename, Size: file.Size, Extention: extention(file.Filename)}
			if result := db.DB.Create(&f); result.Error != nil {
				return result.Error
			}
			extention := f.Extention
			if extention != "" {
				extention = "." + extention
			}
			if err := c.SaveFile(file, fmt.Sprintf("%s/files/%s", dir, f.Id+extention)); err != nil {
				return err
			}

			respond = append(respond, UploadRespond{Id: f.Id, Name: f.Name})

			filenames = append(filenames, file.Filename)
		}

	}

	return c.JSON(respond)

}
