package main

import (
	"log"

	"github.com/chollamatsu/assessment-tax/calTax"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	// start server
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/tax/calculations", calTax.CalTaxWithTaxLev)

	log.Fatal(e.Start(":8080"))
}
