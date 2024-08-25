package errs

import (
	"fmt"

	"github.com/bastean/tgo/internal/pkg/service/errors"
)

func HasMissingArgs(received, required int, where string) error {
	if received < required {
		return errors.NewFailure(&errors.Bubble{
			Where: where,
			What:  fmt.Sprintf("Missing arguments, %d were received and %d are required.", received, required),
		})
	}

	return nil
}
