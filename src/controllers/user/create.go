package user

import (
	"database/sql"
	"desafio-pixelpay/src/lib"
	"desafio-pixelpay/src/model"
	"desafio-pixelpay/src/useCases/createUserUseCase"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func Create(c *gin.Context, db *sql.DB) {
	userModel := model.User{}

	if err := c.ShouldBind(&userModel); err != nil {
		for _, fieldErr := range err.(validator.ValidationErrors) {
			c.JSON(422, gin.H{
				"message": lib.FieldError{Err: fieldErr}.String(),
			})
			return // exit on first error
		}
	}

	fmt.Println(userModel.Document + " " + userModel.Email + " " + userModel.Name)

	status, err := createUserUseCase.CreateUser(userModel, db)

	if err != nil {
		fmt.Println(err)

		c.JSON(status, gin.H{
			"message": "Internal server error",
		})

		return
	}

	if status == 422 {
		c.JSON(status, gin.H{
			"message": "Person already exists",
		})
		return
	}

	c.JSON(status, gin.H{
		"message": "Person created successfully",
	})

}
