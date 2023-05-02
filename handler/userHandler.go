package handler

import (
	"fiberv2/database"
	"fiberv2/model/entity"
	"fiberv2/model/request"
	"fiberv2/model/response"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"log"
)

func UserHandlerGetAll(ctx *fiber.Ctx) error {
	var users []entity.User

	result := database.DB.Find(&users)
	if result.Error != nil {
		log.Println(result.Error)
	}
	return ctx.JSON(users)
}

func UserHandlerCreate(ctx *fiber.Ctx) error {
	user := new(request.UserCreateRequest)
	if err := ctx.BodyParser(user); err != nil {
		return err
	}

	validate := validator.New()
	errValidate := validate.Struct(user)
	if errValidate != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "Error",
			"data":    errValidate.Error(),
		})
	}
	newUser := entity.User{
		Name:    user.Name,
		Address: user.Address,
		Email:   user.Email,
		Phone:   user.Phone,
	}

	errCreateUser := database.DB.Debug().Create(&newUser).Error
	if errCreateUser != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "failed to create in database",
		})
	}
	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    newUser,
	})
}

func UserHandlerFindById(ctx *fiber.Ctx) error {
	userId := ctx.Params("id")

	var user entity.User

	err := database.DB.Debug().First(&user, "id = ?", userId).Error

	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "user not found",
		})
	}

	userResponse := response.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Address:   user.Address,
		Phone:     user.Phone,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return ctx.JSON(fiber.Map{
		"message": "user found",
		"data":    userResponse,
	})
}
