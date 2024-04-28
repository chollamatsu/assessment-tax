package main

import (
	"log"

	"github.com/chollamatsu/assessment-tax/calTax"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	calTax.InitDB()
	// start server
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/tax/calculations", calTax.CalTaxWithTaxLev)

	g := e.Group("/admin")
	g.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if username == "adminTax" || password == "admin!" {
			return true, nil
		}
		return false, nil
	}))

	g.POST("/deductions/personal", calTax.SetPersonalDeduction)
	g.POST("/deductions/k-receipt", calTax.SetKreceipt)

	log.Fatal(e.Start(":8080"))
}
