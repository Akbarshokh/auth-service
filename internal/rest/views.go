package rest

import (
	"auth_service/internal/errs"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Status    string      `json:"status"`
	ErrorCode int         `json:"error_code"`
	ErrorNote string      `json:"error_note"`
	Data      interface{} `json:"data"`
}

//nolint:exhaustruct
func Return(c *gin.Context, data interface{}, err error) {
	switch {
	case err == nil:
		c.JSON(http.StatusOK, Response{
			Status:    StatusMsgSuccess,
			ErrorCode: ErrCodeNoError,
			ErrorNote: "",
			Data:      data,
		})

	case errors.Is(err, errs.ErrValidation):
		c.JSON(http.StatusUnprocessableEntity, Response{
			Status:    StatusMsgFailure,
			ErrorCode: ErrCodeValidation,
			ErrorNote: err.Error(),
		})

	case errors.Is(err, errs.ErrNotFound):
		c.JSON(http.StatusNotFound, Response{
			Status:    StatusMsgFailure,
			ErrorCode: ErrCodeDocumentNotFound,
			ErrorNote: err.Error(),
		})

	default:
		c.JSON(http.StatusInternalServerError, Response{
			Status:    StatusMsgFailure,
			ErrorCode: ErrCodeInternalErr,
			ErrorNote: err.Error(),
		})
	}
}
