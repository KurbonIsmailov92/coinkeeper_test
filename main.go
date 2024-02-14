package main

import (
	budjet "coinceeper/classes/budjets"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	cash := budjet.NewBudjet("Test", 3500, true, true)
	cash.NewBudjetInserterToDb(*cash)

}
