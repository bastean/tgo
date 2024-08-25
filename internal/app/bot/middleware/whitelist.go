package middleware

import (
	"fmt"
	"slices"

	"github.com/bastean/tgo/internal/pkg/service/logger/log"
	tele "gopkg.in/telebot.v3"
)

func Whitelist(usernames []string) tele.MiddlewareFunc {
	return func(next tele.HandlerFunc) tele.HandlerFunc {
		return func(c tele.Context) error {
			username := c.Message().Sender.Username

			if slices.Contains(usernames, username) {
				return next(c)
			}

			log.Error(fmt.Sprintf("(Whitelist) Blocked User [%s]", username))

			return nil
		}
	}
}
