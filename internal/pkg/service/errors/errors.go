package errors

import (
	"github.com/bastean/tgo/pkg/context/domain/errors"
)

type (
	Bubble = errors.Bubble
)

type (
	ErrInvalidValue = errors.ErrInvalidValue
	ErrAlreadyExist = errors.ErrAlreadyExist
	ErrNotExist     = errors.ErrNotExist
	ErrFailure      = errors.ErrFailure
	ErrInternal     = errors.ErrInternal
)

var (
	Panic    = errors.Panic
	BubbleUp = errors.BubbleUp
	As       = errors.As
	Is       = errors.Is
)

var (
	NewFailure  = errors.NewFailure
	NewInternal = errors.NewInternal
)

func IsNot(err error, target error) bool {
	return err != nil && !Is(err, target)
}
