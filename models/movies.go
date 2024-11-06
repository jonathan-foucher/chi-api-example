package models

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Movie struct {
	Id int32 `json:"id" binding:"required"`
	Title string `json:"title" binding:"required"`
	ReleaseDate pgtype.Date `json:"release_date" binding:"required"`
}
