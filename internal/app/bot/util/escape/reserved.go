package escape

import (
	"regexp"
)

func ReservedCharacters(message string) string {
	reserved := regexp.MustCompile(`[.-]{1}`)

	message = reserved.ReplaceAllStringFunc(message, func(match string) string {
		return "\\" + match
	})

	return message
}
