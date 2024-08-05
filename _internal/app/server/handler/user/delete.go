package user

import (
	"net/http"

	"github.com/bastean/tgo/internal/app/server/util/errs"
	"github.com/bastean/tgo/internal/app/server/util/key"
	"github.com/bastean/tgo/internal/app/server/util/reply"
	"github.com/bastean/tgo/internal/pkg/service/errors"
	"github.com/bastean/tgo/internal/pkg/service/user"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, exists := c.Get(key.UserId)

		if !exists {
			errs.AbortErr(c, errs.MissingKey(key.UserId, "Delete"))
			return
		}

		command := new(user.DeleteCommand)

		err := c.BindJSON(command)

		if err != nil {
			errs.AbortErr(c, errs.BindingJSON(err, "Delete"))
			return
		}

		command.Id = id.(string)

		err = user.Delete.Handle(command)

		if err != nil {
			errs.AbortErr(c, errors.BubbleUp(err, "Delete"))
			return
		}

		session := sessions.Default(c)

		session.Clear()

		session.Options(sessions.Options{
			Path:   "/",
			MaxAge: -1,
		})

		err = session.Save()

		if err != nil {
			errs.AbortErr(c, errs.SessionSave(err, "Delete"))
			return
		}

		c.JSON(http.StatusOK, &reply.JSON{
			Success: true,
			Message: "Account deleted",
		})
	}
}
