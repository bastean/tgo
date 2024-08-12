package user

import (
	"github.com/bastean/tgo/pkg/context/application/portfolio/price"
	"github.com/bastean/tgo/pkg/context/application/user/create"
	"github.com/bastean/tgo/pkg/context/application/user/delete"
	"github.com/bastean/tgo/pkg/context/application/user/read"
	"github.com/bastean/tgo/pkg/context/application/user/update"
)

var (
	Create *create.Create
	Read   *read.Read
	Update *update.Update
	Delete *delete.Delete
)

var (
	Price *price.Price
)
