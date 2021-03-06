package controllers

import (
	"strconv"

	"github.com/dackmagic115/go-fiber-kickstart/database"
	"github.com/dackmagic115/go-fiber-kickstart/middleware"
	"github.com/dackmagic115/go-fiber-kickstart/models"
	"github.com/dackmagic115/go-fiber-kickstart/util"
	"github.com/gofiber/fiber/v2"
)

func AllUsers(c *fiber.Ctx) error {
	if err := middleware.IsAuthorized(c, "users"); err != nil {
		return nil
	}

	page, _ := strconv.Atoi(c.Query("page", "1"))

	return c.JSON(util.Paginate(database.DB, &models.User{}, page))
}

func CreateUser(c *fiber.Ctx) error {
	if err := middleware.IsAuthorized(c, "users"); err != nil {
		return nil
	}

	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return err
	}

	user.SetPassword("1234")

	database.DB.Create(&user)

	return c.JSON(user)
}

func GetUser(c *fiber.Ctx) error {
	if err := middleware.IsAuthorized(c, "users"); err != nil {
		return nil
	}

	id, _ := strconv.Atoi(c.Params("id"))

	user := models.User{
		Id: uint(id),
	}

	database.DB.Preload("Role").Find(&user)

	return c.JSON(user)
}

func UpdateUser(c *fiber.Ctx) error {
	if err := middleware.IsAuthorized(c, "users"); err != nil {
		return nil
	}

	id, _ := strconv.Atoi(c.Params("id"))

	user := models.User{
		Id: uint(id),
	}

	if err := c.BodyParser(&user); err != nil {
		return err
	}

	database.DB.Model(&user).Updates(user)

	return c.JSON(user)
}

func DeletedUser(c *fiber.Ctx) error {
	if err := middleware.IsAuthorized(c, "users"); err != nil {
		return nil
	}

	id, _ := strconv.Atoi(c.Params("id"))

	user := models.User{
		Id: uint(id),
	}

	database.DB.Delete(&user)

	return nil
}
