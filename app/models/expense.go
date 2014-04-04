package models

import (
  "fmt"
  "github.com/revel/revel"
)

type Expense struct {
  ExpenseId             int
  Name                  string
  Amount                float64
  DateStr               string
  AccountId             int
  PayerId               int

  Date                  time.Time
  Account               Account
  Payer                 Participant
  ParticipantExpenses   []*ParticipantExpense
}

func (expense *Expense) Validate(v *revel.Validation) {
  v.Required(expense.Name)
  v.Required(expense.Amount)
  v.Required(expense.Date)
  v.Required(expense.Account)
  v.Required(expense.Payer)
  v.Required(expense.ParticipantExpenses)
}
