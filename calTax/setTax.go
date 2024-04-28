package calTax

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
)

func SetTaxLevel(idx int, TaxVal float64) {
	taxDetails[idx].Tax = TaxVal
}

func SetPersonalDeduction(c echo.Context) error {
	personalDeduction := Personal{}
	err := c.Bind(&personalDeduction)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Err{Message: err.Error()})
	}
	// connect db
	if personalDeduction.Amount < 0 {
		return c.JSON(http.StatusBadRequest, Err{Message: lowerTax})
	}
	stmt, err := db.Prepare(updatePersonalDeduction)
	if err != nil {
		log.Fatal(faltalErrorPrepare, err)
	}

	if _, err := stmt.Exec(1, personalDeduction.Amount); err != nil {
		log.Fatal(faltalErrorExecute, err)
	}

	return c.JSON(http.StatusOK, Personal{Amount: GetPersonalDeduction()})
}

func SetKreceipt(c echo.Context) error {
	newKreceipt := Receipt{}
	err := c.Bind(&newKreceipt)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Err{Message: err.Error()})
	}
	// connect db
	stmt, err := db.Prepare(updateKreceipt)
	if err != nil {
		log.Fatal(faltalErrorPrepare, err)
	}

	if _, err := stmt.Exec(1, newKreceipt.Amount); err != nil {
		log.Fatal(faltalErrorExecute, err)
	}

	return c.JSON(http.StatusOK, GetKreceipt())
}
