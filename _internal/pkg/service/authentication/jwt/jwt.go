package jwt

import (
	"github.com/bastean/tgo/internal/pkg/service/env"
	"github.com/bastean/tgo/pkg/context/shared/infrastructure/authentications/jwt"
)

type Payload = jwt.Payload

var (
	JWT      = jwt.New(env.JWTSecretKey)
	Generate = JWT.Generate
	Validate = JWT.Validate
)
