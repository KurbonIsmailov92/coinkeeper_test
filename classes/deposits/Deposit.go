package deposit

type Deposit struct {
	Name       string
	DestAmount float64
	Balance    float64
	Percent    float64
}

func NewDeposit(name string, destination, balance float64) *Deposit {
	return &Deposit{
		Name: name,
		DestAmount: destination,
		Balance: balance,
	}
}
func (d *Deposit) DepositPercentCouner(Deposit) *Deposit{
	d.Percent = d.Balance / d.DestAmount * 100
	return &Deposit{}

}
