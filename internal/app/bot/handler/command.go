package handler

import (
	tele "gopkg.in/telebot.v3"
)

type Command = map[tele.Command]tele.HandlerFunc
