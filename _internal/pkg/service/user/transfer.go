package user

import (
	"github.com/bastean/tgo/pkg/context/user/infrastructure/transport/mail"
	"github.com/bastean/tgo/pkg/context/user/infrastructure/transport/terminal"
)

type (
	MailConfirmation     = mail.Confirmation
	TerminalConfirmation = terminal.Confirmation
)
