package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/laisolvx/crud-golang/src/configuration/rest_err"
	"github.com/laisolvx/crud-golang/src/controller/model/request"
)

func CreateUser(c *gin.Context) {

	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restErr := rest_err.NewBadResquetError(
			fmt.Sprintf("The are some incorrect filds, error=%s\n", err.Error()))

		c.JSON(restErr.Code, restErr)
		return
	}

	fmt.Println(userRequest)
}
