package routers

import (
	"database/sql"
	"desafio-pixelpay/src/controllers/user"
	"github.com/gin-gonic/gin"
)

func InitRoutes(db *sql.DB) *gin.Engine {
	router := gin.Default()

	userGroup := router.Group("/user")
	{
		userGroup.POST("/create", func(c *gin.Context) {
			user.Create(c, db)
		})
	}

	bankGroup := router.Group("/bank")
	{
		bankGroup.GET("/balance")
		bankGroup.POST("/deposit")
		bankGroup.POST("/withdraw")
		bankGroup.POST("/transaction")
	}

	return router
}
