package expense

type Expense struct {
	ID int
    Description string
    Amount      float64
    Date        string
    Type        string
}

func NewExpense(date string, amount float64, description string, typeOfExpense string) *Expense {
    return &Expense{
        Description: description,
        Amount:      amount,
        Date:        date,
        Type:        typeOfExpense,
    }
}
