package portfolio

import (
	"github.com/bastean/tgo/internal/app/bot/command/portfolio/review"
	"github.com/bastean/tgo/internal/app/bot/handler"
)

var Routing = handler.Command{
	review.Command: review.Run,
}
