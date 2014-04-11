package models

type ParticipantExpense struct {
	Id            int
	Amount        float64
	ParticipantId int
	ExpenseId     int
}
