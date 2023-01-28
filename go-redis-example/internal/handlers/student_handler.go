package handlers

import (
	"net/http"

	"git.garena.com/aldino.rahman/go-redis-example/internal/dto/request"
	"git.garena.com/aldino.rahman/go-redis-example/internal/dto/response"
	"github.com/gin-gonic/gin"
)


func (h *Handler) GetStudentbyId(c *gin.Context)  {
	id:= c.Param("id")
	student,err := h.studentUseCase.GetStudentbyId(id)
	if err != nil{
		c.AbortWithStatusJSON(http.StatusInternalServerError,&response.Standard{
			Code: http.StatusInternalServerError,
			Message:err.Error(),
		})
	}

	c.JSON(http.StatusOK,&response.Standard{
		Code: http.StatusOK,
		Message: http.StatusText(http.StatusOK),
		Data: student,
	})

}

func (h *Handler) AddStudent(c *gin.Context)  {
	body:= request.StudentDTO{}
	c.ShouldBind(&body)
	err := h.studentUseCase.AddStudent(&body)
	if err != nil{
		c.AbortWithStatusJSON(http.StatusInternalServerError,&response.Standard{
			Code: http.StatusInternalServerError,
			Message:err.Error(),
		})
	}
	c.JSON(http.StatusCreated,&response.Standard{
		Code: http.StatusCreated,
		Message: http.StatusText(http.StatusCreated),
		Data: nil,
	})

}