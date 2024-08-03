package page

import (
	"github.com/bastean/tgo/internal/app/server/component/page/home"
	"github.com/bastean/tgo/internal/app/server/util/errs"
	"github.com/gin-gonic/gin"
)

func Home() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := home.Page().Render(c.Request.Context(), c.Writer); err != nil {
			errs.AbortErr(c, errs.Render(err, "Home"))
		}
	}
}
