package calTax

import "log"

func GetTaxLevel(TaxVal float64) TotalTax {
	total := TotalTax{
		Tax:      TaxVal,
		TaxLevel: taxDetails,
	}
	return total
}

func GetPersonalDeduction() float64 {
	stmt, err := db.Prepare("SELECT id, personalDeduction FROM constants where id=$1")
	if err != nil {
		log.Fatal("can'tprepare query one row statment", err)
	}

	rowId := 1
	row := stmt.QueryRow(rowId)
	var personalDeduction float64
	var id int

	err = row.Scan(&id, &personalDeduction)
	if err != nil {
		log.Fatal("can't Scan row into variables", err)
	}

	return personalDeduction
}
