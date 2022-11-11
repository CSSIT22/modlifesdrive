package routes

import (
	"github.com/CSSIT22/modlifesdrive/db"
	"github.com/gofiber/fiber/v2"
	"os"
)

func GetFile(c *fiber.Ctx) error {
	id := c.Params("id")
	fileinfo := db.FileInfo{Id: id}
	if result := db.DB.First(&fileinfo); result.Error != nil {
		return result.Error
	}
	ext := fileinfo.Extention
	if ext != "" {
		ext = "." + ext
	}
	dir, err := os.Getwd()
	if err != nil {
		return err
	}
	return c.SendFile(dir + "/files/" + fileinfo.Id + ext)
	//return c.JSON(fileinfo)

}
