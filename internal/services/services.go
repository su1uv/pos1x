package services

import "github.com/su1uv/pos1x/internal/db"

type Services struct {
	db *db.Queries
}

func NewServices(queries *db.Queries) *Services {
	return &Services{db: queries}
}
