package bank

import (
	"database/sql"
	"desafio-pixelpay/src/useCases/bank/balanceUseCase"
	"github.com/gin-gonic/gin"
)

func Balance(c *gin.Context, db *sql.DB) {
	personId := c.Param("personId")

	balance, err := balanceUseCase.AccountBalance(personId, db)

	if err != nil {
		c.JSON(422, gin.H{
			"message": "Internal server error",
		})
		return
	}

	c.JSON(200, gin.H{
		"balance": balance,
	})
}
