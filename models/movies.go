package models


type Movie struct {
	Id int32 `json:"id" binding:"required"`
	Title string `json:"title" binding:"required"`
	ReleaseDate Date `json:"release_date" binding:"required"`
}
