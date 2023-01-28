package handler

import (
	"errors"
	"net/http"
	"strconv"

	"git.garena.com/sea-labs-id/batch-05/gin-template/entity"
	"git.garena.com/sea-labs-id/batch-05/gin-template/util"
	"github.com/gin-gonic/gin"
)

func (h *Handler) AddBorrowBook(c *gin.Context) {
	var reqBorrowBook *entity.BorrowRequest
	err := c.ShouldBindJSON(&reqBorrowBook)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    "BAD_REQUEST",
			"message": util.ErrBadRequest.Error(),
		})
		return
	}

	rec := reqBorrowBook.ToBorrowingRecord()
	user, _ := util.ParseToken(c)

	if user.Id != rec.UserId {
		c.JSON(http.StatusForbidden, gin.H{
			"code":    "FORBIDDEN",
			"message": "FORBIDDEN ACCESS",
		})
		return
	}

	borrowBook, err := h.borrowingUsecase.AddBorrowBook(rec)
	if err != nil {
		if errors.Is(err, util.ErrInternalServer) {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":    "INTERNAL_SERVER_ERROR",
				"message": err.Error(),
			})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    "BAD_REQUEST",
			"message": errors.New("book is empty").Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, borrowBook)
}

func (h *Handler) ReturnBorrowBook(c *gin.Context) {
	id := c.Param("id")
	intId, _ := strconv.Atoi(id)
	user, _ := util.ParseToken(c)

	userBorrow, err := h.borrowingUsecase.GetRecordById(intId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    "BAD_REQUEST",
			"message": errors.New("bad request").Error(),
		})
		return
	}

	if user.Id != userBorrow.UserId {
		c.JSON(http.StatusForbidden, gin.H{
			"code":    "FORBIDDEN",
			"message": errors.New("forbidden access").Error(),
		})
		return
	}

	rec, err := h.borrowingUsecase.ReturnBorrowBook(intId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    "INTERNAL_SERVER_ERROR",
			"message": util.ErrInternalServer.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, rec)
}
