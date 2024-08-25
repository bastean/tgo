package bot

import (
	"fmt"
	"time"

	"github.com/bastean/tgo/internal/app/bot/command"
	"github.com/bastean/tgo/internal/app/bot/middleware"
	"github.com/bastean/tgo/internal/app/bot/router"
	"github.com/bastean/tgo/internal/pkg/service/env"
	"github.com/bastean/tgo/internal/pkg/service/errors"
	"github.com/bastean/tgo/internal/pkg/service/logger/log"
	tele "gopkg.in/telebot.v3"
)

var Bot = &struct {
	Telegram string
}{
	Telegram: log.Bot("Telegram"),
}

var (
	err     error
	Session *tele.Bot
)

func Up() error {
	log.Starting(Bot.Telegram)

	settings := tele.Settings{
		Token:  env.BotTelegramToken,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	Session, err = tele.NewBot(settings)

	if err != nil {
		log.CannotBeStarted(Bot.Telegram)

		return errors.NewInternal(&errors.Bubble{
			Where: "Up",
			What:  "Failure to build a Telegram Bot",
			Who:   err,
		})
	}

	log.Started(Bot.Telegram)

	log.Info(fmt.Sprintf("%s logged in as @%s", Bot.Telegram, Session.Me.Username))

	if usernames, ok := env.HasBotTelegramWhitelistUsernames(); ok {
		Session.Use(middleware.Whitelist(usernames))
		log.Info(fmt.Sprintf("%s whitelist usernames %s", Bot.Telegram, usernames))
	}

	if err = Session.SetCommands(command.List()); err != nil {
		return errors.NewInternal(&errors.Bubble{
			Where: "Up",
			What:  "Cannot set Telegram Bot command list",
			Who:   err,
		})
	}

	router.Updates(Session)

	Session.Start()

	return nil
}

func Down() error {
	log.Stopping(Bot.Telegram)

	if _, err := Session.Close(); err != nil {
		log.CannotBeStopped(Bot.Telegram)

		return errors.NewInternal(&errors.Bubble{
			Where: "Down",
			What:  "Failure to close Telegram Bot instance",
			Who:   err,
		})
	}

	log.Stopped(Bot.Telegram)

	return nil
}
