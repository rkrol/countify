package models

type Account struct {
	Id             string `json:"id,omitempty"`
	Name           string `json:"name" binding:"required"`
	CreationUserId string `json:"creationUserId" binding:"required"`
	CreationDate   string `json:"creationDate" binding:"required"`
}
