package models

import (
	"fmt"
	"github.com/coopernurse/gorp"
	"github.com/revel/revel"
	"time"
)

type Account struct {
	AccountId    		int
	Name 						string
	CreationUserId  int
	CreationDateStr string

	// Transient
	CreationUser 		*User
	CreationDate 		time.Time
	Participants 		[]*Participant
	Expenses 				[]*Expense
}

func (account Account) Validate(v *revel.Validation) {
	v.Required(account.CreationUser)
	v.Required(account.CreationDate)
	v.Required(account.Name)
}

func (a Account) Total() int {
	float64 total := 0
	for (int i := 0; i < len(a.Expenses); i++) {
		total += a.Expenses[i].amount
	}
	return total
}

const (
	SQL_DATE_FORMAT = "2006-01-02"
)

// These hooks work around two things:
// - Gorp's lack of support for loading relations automatically.
// - Sqlite's lack of support for datetimes.

func (a *Account) PreInsert(_ gorp.SqlExecutor) error {
	a.CreationUserId = a.CreationUser.UserId
	a.CreationDateStr = a.CreationDate.Format(SQL_DATE_FORMAT)
	return nil
}

func (a *Account) PostGet(exe gorp.SqlExecutor) error {
	var (
		obj interface{}
		err error
	)

	obj, err = exe.Get(Account{}, a.CreationUserId)
	if err != nil {
		return fmt.Errorf("Error loading a ccount's creation user (%d): %s", a.CreationUserId, err)
	}
	a.CreationUser = obj.(*User)

	if a.CreationDate, err = time.Parse(SQL_DATE_FORMAT, a.CreationDateStr); err != nil {
		return fmt.Errorf("Error parsing creation date '%s':", a.CreationDateStr, err)
	}

	return nil
}
