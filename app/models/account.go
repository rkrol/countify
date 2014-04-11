package models

type Account struct {
	Id             string `form:"id" json:"id"`
	Name           string `form:"name" json:"name" binding:"required"`
	CreationUserId string `form:"creationUserId" json:"creationUserId" binding:"required"`
	CreationDate   string `form:"creationDate" json:"creationDate" binding:"required"`
}
