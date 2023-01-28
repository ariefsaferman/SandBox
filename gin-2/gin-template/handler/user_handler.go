package handler

import (
	"net/http"

	"git.garena.com/sea-labs-id/batch-05/gin-template/entity"
	"git.garena.com/sea-labs-id/batch-05/gin-template/util"
	"github.com/gin-gonic/gin"
)

func (h *Handler) Login(c *gin.Context) {
	var user *entity.UserLoginRequest
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    "BAD_REQUEST",
			"message": util.ErrBadRequest.Error(),
		})
		return
	}

	res, err := h.userUsecase.Login(user)
	if err != nil {
		// should handle user not found and wrong password
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    "BAD_REQUEST",
			"message": util.ErrBadRequest.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"payload": res,
	})

}
