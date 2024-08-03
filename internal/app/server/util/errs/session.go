package errs

import (
	"github.com/bastean/tgo/internal/pkg/service/errors"
)

func SessionSave(who error, where string) error {
	return errors.NewInternal(&errors.Bubble{
		Where: where,
		What:  "Failure to save session",
		Who:   who,
	})
}
