package smtp

import (
	"github.com/bastean/tgo/pkg/context/shared/infrastructure/transports/smtp"
)

type SMTP = smtp.SMTP

var (
	Open = smtp.Open
)
