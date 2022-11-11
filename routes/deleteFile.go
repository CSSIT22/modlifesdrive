package routes

import (
	"fmt"
	"github.com/CSSIT22/modlifesdrive/db"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm/clause"
	"os"
)

type deleteRequest struct {
	Id string `json:"id"`
}

func DeleteFile(c *fiber.Ctx) error {
	req := new(deleteRequest)
	if err := c.BodyParser(req); err != nil {
		return err
	}
	var deletedFile db.FileInfo
	if result := db.DB.Clauses(clause.Returning{}).Where("id=?", req.Id).Delete(&deletedFile); result.Error != nil {
		return result.Error
	}
	fmt.Println("Deleted")
	fmt.Println(deletedFile.Name)
	ext := deletedFile.Extention
	if ext != "" {
		ext = "." + ext
	}
	if err := os.Remove("./files/" + deletedFile.Id + ext); err != nil {
		return err
	}
	return c.JSON(deletedFile)
}
