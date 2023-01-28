package response

import (
	"errors"
	"log"
	"net/http"
	"url-shortener/internal/dto/response"
	cerrors "url-shortener/internal/utils/errors"
	clog "url-shortener/internal/utils/log"

	"github.com/gin-gonic/gin"
)

func ResponseSuccessWithoutMessage(c *gin.Context, data interface{}, code ...int) {
	ResponseSuccess(c, data, "", code...)
}

func ResponseSuccessWithoutData(c *gin.Context, msg string, code ...int) {
	ResponseSuccess(c, nil, msg, code...)
}

func ResponseSuccess(c *gin.Context, data interface{}, msg string, code ...int) {
	resCode := http.StatusOK
	if len(code) > 0 {
		resCode = code[0]
	}

	responseJSON(c, data, msg, resCode)
}

func ResponseError(c *gin.Context, err error) {
	resCode := http.StatusInternalServerError
	var ew *cerrors.ErrorWrapper
	if errors.As(err, &ew) {
		switch ew.Code {
		case cerrors.CodeClientError:
			resCode = http.StatusBadRequest
		case cerrors.CodeNotFoundError:
			resCode = http.StatusNotFound
		case cerrors.CodeConflictError:
			resCode = http.StatusConflict
		}

		if resCode != http.StatusInternalServerError {
			if ew.Err != nil {
				responseJSON(c, nil, ew.Err.Error(), resCode)
				return
			}
			responseJSON(c, nil, ew.Message, resCode)
			return
		}

		// logging server error
		clog.LogAppErr(ew)

		responseJSON(c, nil, cerrors.MsgServerError, resCode)
		return
	}

	log.Println("error: ", err.Error())
	responseJSON(c, nil, cerrors.MsgServerError, resCode)
}

func responseJSON(c *gin.Context, data interface{}, msg string, code int) {
	res := response.JsonResponse{
		Data:    data,
		Message: msg,
	}

	c.JSON(code, res)
}
