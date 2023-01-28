package handler

import (
	"net/http"
	"strings"

	"git.garena.com/sea-labs-id/batch-05/gin-template/entity"
	"git.garena.com/sea-labs-id/batch-05/gin-template/util"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetBooks(c *gin.Context) {
	books, err := h.bookUsecase.GetBooks()
	// if error happens in database
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    "INTERNAL_SERVER_ERROR",
			"message": err.Error(),
		})
		return
	}

	// var responBooks []*entity.BookResponse
	// for _, book := range books {
	// 	responBooks = append(responBooks, book.ToResponseBook())
	// }

	c.JSON(http.StatusOK, books)
}

func (h *Handler) AddBook(c *gin.Context) {
	var reqBook *entity.Book
	err := c.ShouldBindJSON(&reqBook)

	// if request qty < 0 or description < 10 characters
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    "BAD_REQUEST",
			"message": util.ErrBadRequest.Error(),
		})
		return
	}

	book, err := h.bookUsecase.AddBook(reqBook)
	if err != nil {
		// handle error if there is same title in database when add book
		if strings.Contains(err.Error(), "duplicate key") {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    "SAME_TITLE_ERROR",
				"message": util.ErrBookSameTitle.Error(),
			})
			return
		}

		// handle error if there is an error from server
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    "INTERNAL_SERVER_ERROR",
			"message": util.ErrInternalServer,
		})
		return
	}
	// responBook := book.ToResponseBook()
	c.JSON(http.StatusOK, book)
}
