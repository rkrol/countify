package models

import (
  "fmt"
  "github.com/revel/revel"
)

type Participant struct {
  ParticipantId   int
  AccountId       int
  UserId          int

  Account         Account
  User            User
}

func (participant *Participant) Validate(v *revel.Validation) {
  v.Required(participant.Account)
  v.Required(participant.User)
}
