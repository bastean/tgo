package middleware

import (
	"net/http"

	"github.com/bastean/tgo/internal/app/server/util/reply"
	"github.com/bastean/tgo/internal/pkg/service/logger/log"
	"github.com/gin-gonic/gin"
)

func Recover(c *gin.Context, err any) {
	log.Error(err.(error).Error())
	c.AbortWithStatusJSON(http.StatusInternalServerError, &reply.JSON{Message: "Server error. Try again later."})
}
