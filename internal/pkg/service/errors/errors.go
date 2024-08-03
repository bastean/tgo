package errors

import (
	"github.com/bastean/tgo/pkg/context/shared/domain/errors"
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
	BubbleUp = errors.BubbleUp
	As       = errors.As
	Is       = errors.Is
)

var (
	NewFailure  = errors.NewFailure
	NewInternal = errors.NewInternal
)
