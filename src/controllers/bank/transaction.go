package bank

import (
	"database/sql"
	"desafio-pixelpay/src/lib"
	"desafio-pixelpay/src/model"
	"desafio-pixelpay/src/useCases/bank/transactionUseCase"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func Transaction(c *gin.Context, db *sql.DB) {
	transactionModel := model.Transaction{}

	if err := c.ShouldBind(&transactionModel); err != nil {
		for _, fieldErr := range err.(validator.ValidationErrors) {
			c.JSON(422, gin.H{
				"message": lib.FieldError{Err: fieldErr}.String(),
			})
			return
		}
	}

	transactionStatus, err, message := transactionUseCase.TransactionValue(transactionModel, db)

	if err != nil {
		c.JSON(transactionStatus, gin.H{
			"message": message,
		})
		return
	}

	if transactionStatus != 200 {
		c.JSON(transactionStatus, gin.H{
			"message": message,
		})
		return
	}

	c.JSON(transactionStatus, gin.H{
		"message": "Transaction successfully",
	})
}
