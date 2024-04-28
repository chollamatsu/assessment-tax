package calTax

import (
	"math"
	"net/http"

	"github.com/labstack/echo"
)

type Err struct {
	Message string `json:"message"`
}

type TotalTax struct {
	Tax float64 `json:"tax"`
}

type Tax struct {
	TotalIncome float64   `json:"totalIncome"`
	Wht         float64   `json:"wht"`
	Allowances  Allowance `json:"allowances"`
}

type Allowance []struct {
	AllowanceType string  `json:allowanceType`
	Amount        float64 `json:amount`
}

var personalDeduction = float64(6 * math.Pow10(4))
var kReceipt = float64(5 * math.Pow10(4))

func TaxLevel(income float64) float64 {
	if income >= 0 && income < float64(15*math.Pow10(4)) {
		return 0
	} else if income > float64(15*math.Pow10(4)) && income <= float64(5*math.Pow10(5)) {
		return 0.1 * (income - float64(15*math.Pow10(4)))
	} else if income <= float64(5*math.Pow10(5)) && income <= float64(1*math.Pow10(6)) {
		return 1.5 * (income - float64(5*math.Pow10(5)))
	} else if income <= float64(1*math.Pow10(6)) && income <= float64(2*math.Pow10(6)) {
		return 2 * (income - float64(1*math.Pow10(6)))
	} else {
		return 3.5 * (income - float64(2*math.Pow10(6)))
	}
}

func Contains(arr []string, target string) bool {
	for _, i := range arr {
		if i == target {
			return true
		}
	}
	return false
}

func ValidateAllowances(allowanceList Allowance) float64 {
	typeList := []string{"donation", "k-receipt", "tax-free-shop"}
	total := float64(0)
	for _, val := range allowanceList {
		if !Contains(typeList, val.AllowanceType) || val.Amount < 0 {
			return -1
		}
		if val.AllowanceType == "donation" {
			if val.Amount >= 1*math.Pow10(5) {
				total += 1 * math.Pow10(5)
			} else {
				total += val.Amount
			}
		}
		if val.AllowanceType == "tax-free-shop" {
			if val.Amount >= 1*math.Pow10(5) {
				total += 1 * math.Pow10(5)
			} else {
				total += val.Amount
			}
		}
		if val.AllowanceType == "k-receipt" {
			if val.Amount >= kReceipt {
				total += kReceipt
			}
		}
	}
	return float64(total)
}

func ValidateTaxProps(taxObj Tax) bool {
	totalIncome, wht, allowances := taxObj.TotalIncome, taxObj.Wht, taxObj.Allowances
	r1 := wht >= 0 || wht > totalIncome
	r2 := ValidateAllowances(allowances) != -1
	r3 := totalIncome >= 0 && wht >= 0
	if !r3 {
		return false
	} else if !r1 {
		return false
	} else if !r2 {
		return false
	}
	return true
}

func CalTaxWithWht(tax float64, wht float64) float64 {
	if wht > tax {
		return 0
	}
	return tax - wht
}

func CalTaxWithTaxLev(c echo.Context) error {
	newTax := Tax{}
	err := c.Bind(&newTax)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Err{Message: err.Error()})
	}

	isValid := ValidateTaxProps(newTax)

	if !isValid {
		return c.JSON(http.StatusBadRequest, Err{Message: "Allownance propperties does not valid"})
	}
	totalIncome := newTax.TotalIncome

	totalIncome = totalIncome - personalDeduction
	totalIncome = CalTaxWithWht(totalIncome, newTax.Wht)
	totalIncome = totalIncome - ValidateAllowances(newTax.Allowances)
	val := TaxLevel(totalIncome)
	return c.JSON(http.StatusOK, TotalTax{Tax: val})
}
