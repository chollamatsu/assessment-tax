package calTax

import "math"

type Err struct {
	Message string `json:"message"`
}

type Personal struct {
	Amount float64 `json :"amount"`
}

type Receipt struct {
	Amount float64 `json :"amount"`
}

type TotalTax struct {
	Tax      float64     `json : 'tax`
	TaxLevel []taxDetail `json: "taxLevel`
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

type taxDetail struct {
	Level string  `json:"level"`
	Tax   float64 `json:"tax"`
}

var kReceipt = float64(5 * math.Pow10(4))

var taxDetails = []taxDetail{
	{
		Level: "0-150,000",
		Tax:   0.0,
	},
	{
		Level: "150,001-500,000",
		Tax:   0.0,
	},
	{
		Level: "500,001-1,000,000",
		Tax:   0.0,
	},
	{
		Level: "1,000,001-2,000,000",
		Tax:   0.0,
	},
	{
		Level: "2,000,001 ขึ้นไป",
		Tax:   0.0,
	},
}
