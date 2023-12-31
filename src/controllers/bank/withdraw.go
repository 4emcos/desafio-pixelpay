package bank

import (
	"database/sql"
	"desafio-pixelpay/src/lib"
	"desafio-pixelpay/src/model"
	"desafio-pixelpay/src/useCases/bank/withdrawUseCase"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func Withdraw(c *gin.Context, db *sql.DB) {
	withdrawModel := model.Withdraw{}

	if err := c.ShouldBind(&withdrawModel); err != nil {
		for _, fieldErr := range err.(validator.ValidationErrors) {
			c.JSON(422, gin.H{
				"message": lib.FieldError{Err: fieldErr}.String(),
			})
			return
		}
	}

	withdrawValue, status, err := withdrawUseCase.WithdrawValue(withdrawModel, db)

	if err != nil {
		c.JSON(status, gin.H{
			"message": "Internal server error",
		})
		return
	}

	if status == 422 {
		c.JSON(status, gin.H{
			"message": "Insufficient balance",
		})
		return
	}

	c.JSON(status, gin.H{
		"message": "Withdraw successfully",
		"balance": withdrawValue,
	})
}
