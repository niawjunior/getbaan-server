package helpers

import (
	"github.com/gin-gonic/gin"
)

type ResponseData struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}

func JSON(w *gin.Context, status int, payload ...interface{}) {
	var res ResponseData

	res.Status = status
	res.Data = payload

	w.JSON(status, res)
}
