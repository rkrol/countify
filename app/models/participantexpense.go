package models

import (
	"github.com/revel/revel"
)

type ParticipantExpense struct {
	Amount						float64
	ParticipantId    	int
	ExpenseId        	int

	Participant 			Participant
	Expense 					Expense
}

func (participantExpense *ParticipantExpense) Validate(v *revel.Validation) {
	v.Required(participantExpense.Amount)
	v.Required(participantExpense.Participant)
	v.Required(participantExpense.Expense)
}
