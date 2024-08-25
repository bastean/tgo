package errs

import (
	"github.com/bastean/tgo/internal/pkg/service/errors"
)

func Response(who error, where string) error {
	return errors.NewInternal(&errors.Bubble{
		Where: where,
		What:  "Cannot respond to an update",
		Who:   who,
	})
}
