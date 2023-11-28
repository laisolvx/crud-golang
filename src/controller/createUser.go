package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/laisolvx/crud-golang/src/configuration/rest_err"
)

func CreateUser(c *gin.Context) {
	err := rest_err.NewBadResquetError("Você chamou a rota de forma errada")
	c.JSON(err.Code, err)
}
