package controllers

import (
	"github.com/gin-gonic/gin"
	"server/logic"
)

func History(c *gin.Context) {
	response := logic.History(c)
	ResponseSuccess(c, response)
	return
}
