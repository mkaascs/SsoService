package services

import (
	"sso-service/internal/domain/dto/jwt/commands"
	"sso-service/internal/domain/dto/jwt/results"
)

type JWT interface {
	Generate(command commands.Generate) (*results.Generate, error)
	Parse(command commands.Parse) (*results.Parse, error)
}
