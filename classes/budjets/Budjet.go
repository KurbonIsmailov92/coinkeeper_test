package budjet

import (
	"database/sql"
	"fmt"
	"time"
)

type Budjet struct {
	ID         int64
	Name       string
	Balance    int64
	IsCredit   bool
	IsTotalSum bool
}

func NewBudjet(name string, balace int64, isCredit bool, total bool) *Budjet {
	return &Budjet{
		Name:       name,
		Balance:    balace,
		IsCredit:   isCredit,
		IsTotalSum: total,
	}

}

func (b *Budjet) BudjetTopUper(amount int64) *Budjet {
	b.Balance += amount
	return &Budjet{}
}

func (b *Budjet) BudjetOutComer(amount int64) *Budjet {
	b.Balance -= amount
	return &Budjet{}
}

func (b *Budjet) BudjetCorrector(newName string, newBalace int64, newIsCredit bool, newTotal bool) *Budjet {
	if newName != b.Name {
		b.Name = newName
	}
	if newBalace != b.Balance {
		b.Balance = newBalace
	}
	if newIsCredit != b.IsCredit {
		b.IsCredit = newIsCredit
	}
	if newTotal != b.IsTotalSum {
		b.IsTotalSum = newTotal
	}
	return &Budjet{}
}

func (b *Budjet) BudjetDeleter(Budjet) {

}

func (b *Budjet) NewBudjetInserterToDb(newBudjet Budjet) *Budjet {

	Db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/coinceeper")
	if err != nil {
		panic(err)
	}
	defer Db.Close()

	insert, err := Db.Query(fmt.Sprintf("INSERT INTO budjets (name, balance, is_credit, in_total_sum, created_at) VALUES('%s','%d','%t','%t','%s')", newBudjet.Name, newBudjet.Balance, newBudjet.IsCredit, newBudjet.IsTotalSum, time.Now().IsDST()))
	if err != nil {
		panic(err)
	}
	fmt.Println("Бюджет добавлен в БД")
	defer insert.Close()

	return &Budjet{} 
}
